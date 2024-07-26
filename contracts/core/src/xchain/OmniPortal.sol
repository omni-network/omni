// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import { ExcessivelySafeCall } from "@nomad-xyz/excessively-safe-call/src/ExcessivelySafeCall.sol";

import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalSys } from "../interfaces/IOmniPortalSys.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Quorum } from "../libraries/Quorum.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { PausableUpgradeable } from "../utils/Pausable.sol";

import { OmniPortalConstants } from "./OmniPortalConstants.sol";
import { OmniPortalStorage } from "./OmniPortalStorage.sol";

contract OmniPortal is
    IOmniPortal,
    IOmniPortalSys,
    IOmniPortalAdmin,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    OmniPortalConstants,
    OmniPortalStorage
{
    using ExcessivelySafeCall for address;

    /**
     * @notice Modifier the requires an action is not paused. An action is paused if:
     *          - actionId is paused for all chains
     *          - actionId is paused for chainId
     *          - All actions are paused
     *         Available actions are ActionXCall and ActionXSubmit, defined in OmniPortalConstants.sol.
     */
    modifier whenNotPaused(bytes32 actionId, uint64 chainId_) {
        require(!_isPaused(actionId, _chainActionId(actionId, chainId_)), "OmniPortal: paused");
        _;
    }

    /**
     * @notice Construct the OmniPortal contract
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialization init params
     * @dev Used to reduce stack depth in initialize()
     * @custom:field owner                  The owner of the contract
     * @custom:field feeOracle              Address of the fee oracle contract
     * @custom:field omniChainId            Chain ID of Omni's EVM execution chain
     * @custom:field omniCChainId           Chain ID of Omni's consensus chain
     * @custom:field xmsgMaxGasLimit        Maximum gas limit for xmsg
     * @custom:field xmsgMinGasLimit        Minimum gas limit for xmsg
     * @custom:field xmsgMaxDataSize        Maximum size of xmsg data in bytes
     * @custom:field xreceiptMaxErrorSize   Maximum size of xreceipt error in bytes
     * @custom:field xsubValsetCutoff       Number of validator sets since the latest that validate an XSubmission
     * @custom:field cChainXMsgOffset       Offset for xmsgs from the consensus chain
     * @custom:field cChainXBlockOffset     Offset for xblocks from the consensus chain
     * @custom:field valSetId               Initial validator set id
     * @custom:field validators             Initial validator set
     */
    struct InitParams {
        address owner;
        address feeOracle;
        uint64 omniChainId;
        uint64 omniCChainId;
        uint64 xmsgMaxGasLimit;
        uint64 xmsgMinGasLimit;
        uint16 xmsgMaxDataSize;
        uint16 xreceiptMaxErrorSize;
        uint8 xsubValsetCutoff;
        uint64 cChainXMsgOffset;
        uint64 cChainXBlockOffset;
        uint64 valSetId;
        XTypes.Validator[] validators;
    }

    /**
     * @notice Initialize the OmniPortal contract
     */
    function initialize(InitParams calldata p) public initializer {
        __Ownable_init(p.owner);

        _setFeeOracle(p.feeOracle);
        _setXMsgMaxGasLimit(p.xmsgMaxGasLimit);
        _setXMsgMinGasLimit(p.xmsgMinGasLimit);
        _setXMsgMaxDataSize(p.xmsgMaxDataSize);
        _setXReceiptMaxErrorSize(p.xreceiptMaxErrorSize);
        _setXSubValsetCutoff(p.xsubValsetCutoff);
        _addValidatorSet(p.valSetId, p.validators);

        omniChainId = p.omniChainId;
        omniCChainId = p.omniCChainId;

        // omni consensus chain uses Finalised+Broadcast shard
        uint64 omniCShard = ConfLevel.toBroadcastShard(ConfLevel.Finalized);
        inXMsgOffset[p.omniCChainId][omniCShard] = p.cChainXMsgOffset;
        inXBlockOffset[p.omniCChainId][omniCShard] = p.cChainXBlockOffset;
    }

    function chainId() public view returns (uint64) {
        return uint64(block.chainid);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Outbound xcall functions                            //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Call a contract on another chain.
     * @param destChainId   Destination chain ID
     * @param conf          Confirmation level
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, uint8 conf, address to, bytes calldata data, uint64 gasLimit)
        external
        payable
        whenNotPaused(ActionXCall, destChainId)
    {
        require(isSupportedDest[destChainId], "OmniPortal: unsupported dest");
        require(to != VirtualPortalAddress, "OmniPortal: no portal xcall");
        require(gasLimit <= xmsgMaxGasLimit, "OmniPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "OmniPortal: gasLimit too low");
        require(data.length <= xmsgMaxDataSize, "OmniPortal: data too large");

        // conf level will always be first byte of shardId. for now, shardId is just conf level
        uint64 shardId = uint64(conf);
        require(isSupportedShard[shardId], "OmniPortal: unsupported shard");

        uint256 fee = feeFor(destChainId, data, gasLimit);
        require(msg.value >= fee, "OmniPortal: insufficient fee");

        outXMsgOffset[destChainId][shardId] += 1;

        emit XMsg(destChainId, shardId, outXMsgOffset[destChainId][shardId], msg.sender, to, data, gasLimit, fee);
    }

    /**
     * @notice Calculate the fee for calling a contract on another chain
     *         Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, gasLimit);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Inbound xcall functions                             //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub  An xchain submission, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub)
        external
        whenNotPaused(ActionXSubmit, xsub.blockHeader.sourceChainId)
        nonReentrant
    {
        XTypes.Msg[] calldata xmsgs = xsub.msgs;
        XTypes.BlockHeader calldata xheader = xsub.blockHeader;
        uint64 valSetId = xsub.validatorSetId;

        require(xheader.consensusChainId == omniCChainId, "OmniPortal: wrong cchain ID");
        require(xmsgs.length > 0, "OmniPortal: no xmsgs");
        require(valSetTotalPower[valSetId] > 0, "OmniPortal: unknown val set");
        require(valSetId >= _minValSet(), "OmniPortal: old val set");

        // check that the attestationRoot is signed by a quorum of validators in xsub.validatorsSetId
        require(
            Quorum.verify(
                xsub.attestationRoot,
                xsub.signatures,
                valSet[valSetId],
                valSetTotalPower[valSetId],
                XSubQuorumNumerator,
                XSubQuorumDenominator
            ),
            "OmniPortal: no quorum"
        );

        // check that blockHeader and xmsgs are included in attestationRoot
        require(
            XBlockMerkleProof.verify(xsub.attestationRoot, xheader, xmsgs, xsub.proof, xsub.proofFlags),
            "OmniPortal: invalid proof"
        );

        // execute xmsgs
        for (uint256 i = 0; i < xmsgs.length; i++) {
            _exec(xheader, xmsgs[i]);
        }
    }

    /**
     * @notice Returns the current XMsg being executed via this portal.
     *          - xmsg().sourceChainId  Chain ID of the source xcall
     *          - xmsg().sender         msg.sender of the source xcall
     *         If no XMsg is being executed, all fields will be zero.
     *          - xmsg().sourceChainId  == 0
     *          - xmsg().sender         == address(0)
     */
    function xmsg() external view returns (XTypes.MsgContext memory) {
        return _xmsg;
    }

    /**
     * @notice Returns true the current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool) {
        return _xmsg.sourceChainId != 0;
    }

    /**
     * @notice Execute an xmsg.
     * @dev Verify if an XMsg is next in its XStream, execute it, increment inXMsgOffset, emit an XReceipt event
     */
    function _exec(XTypes.BlockHeader memory xheader, XTypes.Msg calldata xmsg_) internal {
        uint64 sourceChainId = xheader.sourceChainId;
        uint64 destChainId = xmsg_.destChainId;
        uint64 shardId = xmsg_.shardId;
        uint64 offset = xmsg_.offset;

        require(destChainId == chainId() || destChainId == BroadcastChainId, "OmniPortal: wrong dest chain");
        require(offset == inXMsgOffset[sourceChainId][shardId] + 1, "OmniPortal: wrong offset");

        // verify xmsg conf level matches xheader conf level
        // allow finalized blocks to for any xmsg, so that finalized blocks may correct "fuzzy" xmsgs
        require(
            ConfLevel.Finalized == xheader.confLevel || xheader.confLevel == uint8(shardId),
            "OmniPortal: wrong conf level"
        );

        if (inXBlockOffset[sourceChainId][shardId] < xheader.offset) {
            inXBlockOffset[sourceChainId][shardId] = xheader.offset;
        }

        inXMsgOffset[sourceChainId][shardId] += 1;

        // do not allow user xcalls to the portal
        // only sys xcalls (to _VIRTUAL_PORTAL_ADDRESS) are allowed to be executed on the portal
        if (xmsg_.to == address(this)) {
            emit XReceipt(
                sourceChainId,
                shardId,
                offset,
                0,
                msg.sender,
                false,
                abi.encodeWithSignature("Error(string)", "OmniPortal: no xcall to portal")
            );

            return;
        }

        // set _xmsg to the one we're executing, allowing external contracts to query the current xmsg via xmsg()
        _xmsg = XTypes.MsgContext(sourceChainId, xmsg_.sender);

        (bool success, bytes memory result, uint256 gasUsed) = xmsg_.to == VirtualPortalAddress // calls to VirtualPortalAddress are syscalls
            ? _syscall(xmsg_.data)
            : _call(xmsg_.to, xmsg_.gasLimit, xmsg_.data);

        // reset xmsg to zero
        delete _xmsg;

        bytes memory errorMsg = success ? bytes("") : result;

        emit XReceipt(sourceChainId, shardId, offset, gasUsed, msg.sender, success, errorMsg);
    }

    /**
     * @notice Call an external contract.
     * @dev Returns the result of the call, the gas used, and whether the call was successful.
     * @param to                The address of the contract to call.
     * @param gasLimit          Gas limit of the call.
     * @param data              Calldata to send to the contract.
     */
    function _call(address to, uint256 gasLimit, bytes calldata data) internal returns (bool, bytes memory, uint256) {
        uint256 gasLeftBefore = gasleft();

        // use excessivelySafeCall for external calls to prevent large return bytes mem copy
        (bool success, bytes memory result) =
            to.excessivelySafeCall({ _gas: gasLimit, _value: 0, _maxCopy: xreceiptMaxErrorSize, _calldata: data });

        uint256 gasLeftAfter = gasleft();

        // Ensure relayer sent enough gas for the call
        // See https://github.com/OpenZeppelin/openzeppelin-contracts/blob/bd325d56b4c62c9c5c1aff048c37c6bb18ac0290/contracts/metatx/MinimalForwarder.sol#L58-L68
        if (gasLeftAfter <= gasLimit / 63) {
            // We use invalid opcode to consume all gas and bubble-up the effects, to emulate an "OutOfGas" exception
            assembly {
                invalid()
            }
        }

        return (success, result, gasLeftBefore - gasLeftAfter);
    }

    /**
     * @notice Call a function on the current contract.
     * @dev Reverts on failure. We match _call() return signature for symmetry.
     * @param data      Calldata to execute on the current contract.
     */
    function _syscall(bytes calldata data) internal returns (bool, bytes memory, uint256) {
        uint256 gasUsed = gasleft();
        (bool success, bytes memory result) = address(this).call(data);
        gasUsed = gasUsed - gasleft();

        // if not success, revert with same reason
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }

        return (success, result, gasUsed);
    }

    /**
     * @notice Returns the minimum validator set id that can be used for xsubmissions
     */
    function _minValSet() internal view returns (uint64) {
        return latestValSetId > xsubValsetCutoff
            // plus 1, so the number of accepted valsets == XSubValsetCutoff
            ? (latestValSetId - xsubValsetCutoff + 1)
            : 1;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Syscall functions                               //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Add a new validator set.
     * @dev Only callable via xcall from Omni's consensus chain
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function addValidatorSet(uint64 valSetId, XTypes.Validator[] calldata validators) external {
        require(msg.sender == address(this), "OmniPortal: only self");
        require(_xmsg.sourceChainId == omniCChainId, "OmniPortal: only cchain");
        require(_xmsg.sender == CChainSender, "OmniPortal: only cchain sender");
        _addValidatorSet(valSetId, validators);
    }

    /**
     * @notice Add a new validator set
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function _addValidatorSet(uint64 valSetId, XTypes.Validator[] calldata validators) internal {
        uint256 numVals = validators.length;
        require(numVals > 0, "OmniPortal: no validators");
        require(valSetTotalPower[valSetId] == 0, "OmniPortal: duplicate val set");

        uint64 totalPower;
        XTypes.Validator memory val;
        mapping(address => uint64) storage valSet = valSet[valSetId];

        for (uint256 i = 0; i < numVals; i++) {
            val = validators[i];

            require(val.addr != address(0), "OmniPortal: no zero validator");
            require(val.power > 0, "OmniPortal: no zero power");
            require(valSet[val.addr] == 0, "OmniPortal: duplicate validator");

            totalPower += val.power;
            valSet[val.addr] = val.power;
        }

        valSetTotalPower[valSetId] = totalPower;

        if (valSetId > latestValSetId) latestValSetId = valSetId;

        emit ValidatorSetAdded(valSetId);
    }

    /**
     * @notice Set the network of supported chains & shards
     * @dev Only callable via xcall from Omni's consensus chain
     * @param network_  The new network
     */
    function setNetwork(XTypes.Chain[] calldata network_) external {
        require(msg.sender == address(this), "OmniPortal: only self");
        require(_xmsg.sourceChainId == omniCChainId, "OmniPortal: only cchain");
        require(_xmsg.sender == CChainSender, "OmniPortal: only cchain sender");
        _setNetwork(network_);
    }

    /**
     * @notice Set the network of supported chains & shards
     * @param network_  The new network
     */
    function _setNetwork(XTypes.Chain[] calldata network_) internal {
        _clearNetwork();

        XTypes.Chain calldata c;
        for (uint256 i = 0; i < network_.length; i++) {
            c = network_[i];
            network.push(c);

            // if not this chain, mark as supported dest
            if (c.chainId != chainId()) {
                isSupportedDest[c.chainId] = true;
                continue;
            }

            // if this chain, mark shards as supported
            for (uint256 j = 0; j < c.shards.length; j++) {
                isSupportedShard[c.shards[j]] = true;
            }
        }
    }

    /**
     * @notice Clear the network of supported chains & shards
     */
    function _clearNetwork() private {
        XTypes.Chain storage c;
        for (uint256 i = 0; i < network.length; i++) {
            c = network[i];

            // if not this chain, mark as unsupported dest
            if (c.chainId != chainId()) {
                isSupportedDest[c.chainId] = false;
                continue;
            }

            // if this chain, mark shards as unsupported
            for (uint256 j = 0; j < c.shards.length; j++) {
                isSupportedShard[c.shards[j]] = false;
            }
        }
        delete network;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Admin functions                                 //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Set the inbound xmsg offset for a chain and shard
     * @param sourceChainId    Source chain ID
     * @param shardId          Shard ID
     * @param offset           New xmsg offset
     */
    function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) external onlyOwner {
        inXMsgOffset[sourceChainId][shardId] = offset;
        emit InXMsgOffsetSet(sourceChainId, shardId, offset);
    }

    /**
     * @notice Set the inbound xblock offset for a chain and shard
     * @param sourceChainId    Source chain ID
     * @param shardId          Shard ID
     * @param offset           New xblock offset
     */
    function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) external onlyOwner {
        inXBlockOffset[sourceChainId][shardId] = offset;
        emit InXBlockOffsetSet(sourceChainId, shardId, offset);
    }

    /**
     * @notice Set the fee oracle
     */
    function setFeeOracle(address feeOracle_) external onlyOwner {
        _setFeeOracle(feeOracle_);
    }

    /**
     * @notice Transfer all collected fees to the give address
     * @param to    The address to transfer the fees to
     */
    function collectFees(address to) external onlyOwner {
        uint256 amount = address(this).balance;

        // .transfer() is fine, owner should provide an EOA address that will not
        // consume more than 2300 gas on transfer, and we are okay .transfer() reverts
        payable(to).transfer(amount);

        emit FeesCollected(to, amount);
    }

    /**
     * @notice Set the minimum gas limit for xmsg
     */
    function setXMsgMinGasLimit(uint64 gasLimit) external onlyOwner {
        _setXMsgMinGasLimit(gasLimit);
    }

    /**
     * @notice Set the maximum gas limit for xmsg
     */
    function setXMsgMaxGasLimit(uint64 gasLimit) external onlyOwner {
        _setXMsgMaxGasLimit(gasLimit);
    }

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function setXMsgMaxDataSize(uint16 numBytes) external onlyOwner {
        _setXMsgMaxDataSize(numBytes);
    }

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function setXReceiptMaxErrorSize(uint16 numBytes) external onlyOwner {
        _setXReceiptMaxErrorSize(numBytes);
    }

    /**
     * @notice Set the number of validator sets since the latest that can validate an XSubmission
     */
    function setXSubValsetCutoff(uint8 xsubValsetCutoff_) external onlyOwner {
        _setXSubValsetCutoff(xsubValsetCutoff_);
    }

    /**
     * @notice Pause xcalls and xsubissions from all chains
     */
    function pause() external onlyOwner {
        _pauseAll();
        emit Paused();
    }

    /**
     * @notice Unpause xcalls and xsubissions from all chains
     */
    function unpause() external onlyOwner {
        _unpauseAll();
        emit Unpaused();
    }

    /**
     * @notice Pause xcalls to all chains
     */
    function pauseXCall() external onlyOwner {
        _pause(ActionXCall);
        emit XCallPaused();
    }

    /**
     * @notice Unpause xcalls to all chains
     */
    function unpauseXCall() external onlyOwner {
        _unpause(ActionXCall);
        emit XCallUnpaused();
    }

    /**
     * @notice Pause xcalls to a specific chain
     * @param chainId_   Destination chain ID
     */
    function pauseXCallTo(uint64 chainId_) external onlyOwner {
        _pause(_chainActionId(ActionXCall, chainId_));
        emit XCallToPaused(chainId_);
    }

    /**
     * @notice Unpause xcalls to a specific chain
     * @param chainId_   Destination chain ID
     */
    function unpauseXCallTo(uint64 chainId_) external onlyOwner {
        _unpause(_chainActionId(ActionXCall, chainId_));
        emit XCallToUnpaused(chainId_);
    }

    /**
     * @notice Pause xsubmissions from all chains
     */
    function pauseXSubmit() external onlyOwner {
        _pause(ActionXSubmit);
        emit XSubmitPaused();
    }

    /**
     * @notice Unpause xsubmissions from all chains
     */
    function unpauseXSubmit() external onlyOwner {
        _unpause(ActionXSubmit);
        emit XSubmitUnpaused();
    }

    /**
     * @notice Pause xsubmissions from a specific chain
     * @param chainId_    Source chain ID
     */
    function pauseXSubmitFrom(uint64 chainId_) external onlyOwner {
        _pause(_chainActionId(ActionXSubmit, chainId_));
        emit XSubmitFromPaused(chainId_);
    }

    /**
     * @notice Unpause xsubmissions from a specific chain
     * @param chainId_    Source chain ID
     */
    function unpauseXSubmitFrom(uint64 chainId_) external onlyOwner {
        _unpause(_chainActionId(ActionXSubmit, chainId_));
        emit XSubmitFromUnpaused(chainId_);
    }

    /**
     * @notice Return true if actionId for is paused for the given chain
     */
    function isPaused(bytes32 actionId, uint64 chainId_) external view returns (bool) {
        return _isPaused(actionId, _chainActionId(actionId, chainId_));
    }

    /**
     * @notice Return true if actionId is paused for all chains
     */
    function isPaused(bytes32 actionId) external view returns (bool) {
        return _isPaused(actionId);
    }

    /*
    * @notice Return true if all actions are paused
     */
    function isPaused() external view returns (bool) {
        return _isAllPaused();
    }

    /**
     * @notice An action id with a qualifiying chain id, used as pause keys.
     */
    function _chainActionId(bytes32 actionId, uint64 chainId_) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(actionId, chainId_));
    }

    /**
     * @notice Set the minimum gas limit for xmsg
     */
    function _setXMsgMinGasLimit(uint64 gasLimit) internal {
        require(gasLimit > 0, "OmniPortal: no zero min gas");
        xmsgMinGasLimit = gasLimit;
        emit XMsgMinGasLimitSet(gasLimit);
    }

    /**
     * @notice Set the maximum gas limit for xmsg
     */
    function _setXMsgMaxGasLimit(uint64 gasLimit) internal {
        require(gasLimit > 0, "OmniPortal: no zero max gas");
        xmsgMaxGasLimit = gasLimit;
        emit XMsgMaxGasLimitSet(gasLimit);
    }

    /**
     * @notice Set the maximum data bytes for xmsg
     */
    function _setXMsgMaxDataSize(uint16 numBytes) internal {
        require(numBytes > 0, "OmniPortal: no zero max size");
        xmsgMaxDataSize = numBytes;
        emit XMsgMaxDataSizeSet(numBytes);
    }

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function _setXReceiptMaxErrorSize(uint16 numBytes) internal {
        require(numBytes > 0, "OmniPortal: no zero max size");
        xreceiptMaxErrorSize = numBytes;
        emit XReceiptMaxErrorSizeSet(numBytes);
    }

    /**
     * @notice Set the number of validator sets since the latest that can validate an XSubmission
     */
    function _setXSubValsetCutoff(uint8 xsubValsetCutoff_) internal {
        require(xsubValsetCutoff_ > 0, "OmniPortal: no zero cutoff");
        xsubValsetCutoff = xsubValsetCutoff_;
        emit XSubValsetCutoffSet(xsubValsetCutoff_);
    }

    /**
     * @notice Set the fee oracle
     */
    function _setFeeOracle(address feeOracle_) internal {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");
        feeOracle = feeOracle_;
        emit FeeOracleSet(feeOracle_);
    }
}
