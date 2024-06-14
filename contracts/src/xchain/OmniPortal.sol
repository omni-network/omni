// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-upgrades/contracts/security/PausableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import { ExcessivelySafeCall } from "@nomad-xyz/excessively-safe-call/src/ExcessivelySafeCall.sol";

import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalSys } from "../interfaces/IOmniPortalSys.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Quorum } from "../libraries/Quorum.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";

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
     * @custom:field omniCChainId           Virtual chain ID used in xmsgs from Omni's consensus chain
     * @custom:field xmsgMaxGasLimit        Maximum gas limit for xmsg
     * @custom:field xmsgMinGasLimit        Minimum gas limit for xmsg
     * @custom:field xmsgMaxDataSize        Maximum size of xmsg data in bytes
     * @custom:field xreceiptMaxErrorSize   Maximum size of xreceipt error in bytes
     * @custom:field cChainXMsgOffset       Offset for xmsgs from the consensus chain
     * @custom:field cChainXBlockOffset    Offset for xblocks from the consensus chain
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
        uint64 cChainXMsgOffset;
        uint64 cChainXBlockOffset;
        uint64 valSetId;
        XTypes.Validator[] validators;
    }

    /**
     * @notice Initialize the OmniPortal contract
     */
    function initialize(InitParams calldata p) public initializer {
        _transferOwnership(p.owner);
        _setFeeOracle(p.feeOracle);
        _setXMsgMaxGasLimit(p.xmsgMaxGasLimit);
        _setXMsgMinGasLimit(p.xmsgMinGasLimit);
        _setXMsgMaxDataSize(p.xmsgMaxDataSize);
        _setXReceiptMaxErrorSize(p.xreceiptMaxErrorSize);
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
        whenNotPaused
    {
        // TODO: uncomment when cchain setNetwork xmsgs are enabled
        // require(isSupportedDest[destChainId], "OmniPortal: unsupported dest");

        require(to != _VIRTUAL_PORTAL_ADDRESS, "OmniPortal: no portal xcall");
        require(gasLimit <= xmsgMaxGasLimit, "OmniPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "OmniPortal: gasLimit too low");
        require(data.length <= xmsgMaxDataSize, "OmniPortal: data too large");

        // conf level will always be first byte of shardId. for now, shardId is just conf level
        uint64 shardId = uint64(conf);

        // TODO: uncomment when cchain setNetwork xmsgs are enabled
        // require(isSupportedShard[shardId], "OmniPortal: unsupported shard");

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
     * @param xsub  An xchain submisison, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external whenNotPaused nonReentrant {
        XTypes.Msg[] calldata xmsgs = xsub.msgs;
        XTypes.BlockHeader calldata xheader = xsub.blockHeader;
        uint64 valSetId = xsub.validatorSetId;

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
                XSUB_QUORUM_NUMERATOR,
                XSUB_QUORUM_DENOMINATOR
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
    function xmsg() external view returns (XTypes.MsgShort memory) {
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
     * @dev Verify an XMsg is next in its XStream, execute it, increment inXStreamOffset, emit an XReceipt
     */
    function _exec(XTypes.BlockHeader memory xheader, XTypes.Msg calldata xmsg_) internal {
        uint64 sourceChainId = xmsg_.sourceChainId;
        uint64 destChainId = xmsg_.destChainId;
        uint64 shardId = xmsg_.shardId;
        uint64 offset = xmsg_.offset;

        require(sourceChainId == xheader.sourceChainId, "OmniPortal: wrong source chain"); // TODO: we can remove xmsg sourceChainId, and instead just used xheader.sourceChainId
        require(destChainId == chainId() || destChainId == _BROADCAST_CHAIN_ID, "OmniPortal: wrong dest chain");
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
        _xmsg = XTypes.MsgShort(sourceChainId, xmsg_.sender);

        (bool success, bytes memory result, uint256 gasUsed) = xmsg_.to == _VIRTUAL_PORTAL_ADDRESS // calls to _VIRTUAL_PORTAL_ADDRESS are syscalls
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
     * @param gasLimit          Gas limit of the call
     * @param data              Calldata to send to the contract.
     */
    function _call(address to, uint256 gasLimit, bytes calldata data) internal returns (bool, bytes memory, uint256) {
        uint256 gasLeftBefore = gasleft();

        // use excessivelySafeCall for external calls to prevent large return bytes mem copy
        (bool success, bytes memory result) =
            to.excessivelySafeCall({ _gas: gasLimit, _value: 0, _maxCopy: xreceiptMaxErrorSize, _calldata: data });

        uint256 gasLeftAfter = gasleft();

        // Esnure relayer sent enough gas for the call
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
        return latestValSetId > XSUB_VALSET_CUTOFF
            // plus 1, so the number of accepted valsets == XSUB_VALSET_CUTOFF
            ? (latestValSetId - XSUB_VALSET_CUTOFF + 1)
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
        require(_xmsg.sender == _CCHAIN_SENDER, "OmniPortal: only cchain sender");
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
        require(_xmsg.sender == _CCHAIN_SENDER, "OmniPortal: only cchain sender");
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
     * @notice Pause xcalls
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause xcalls
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    /**
     * @notice Set the minimum gas limit for xmsg
     */
    function _setXMsgMinGasLimit(uint64 gasLimit) internal {
        require(gasLimit > 0, "OmniPortal: no zero min gas");

        uint64 oldMin = xmsgMinGasLimit;
        xmsgMinGasLimit = gasLimit;

        emit XMsgMinGasLimitChanged(oldMin, gasLimit);
    }

    /**
     * @notice Set the maximum gas limit for xmsg
     */
    function _setXMsgMaxGasLimit(uint64 gasLimit) internal {
        require(gasLimit > 0, "OmniPortal: no zero max gas");

        uint64 oldMax = xmsgMaxGasLimit;
        xmsgMaxGasLimit = gasLimit;

        emit XMsgMaxGasLimitChanged(oldMax, gasLimit);
    }

    /**
     * @notice Set the maximum data bytes for xmsg
     */
    function _setXMsgMaxDataSize(uint16 numBytes) internal {
        require(numBytes > 0, "OmniPortal: no zero max size");

        uint16 oldMax = xmsgMaxDataSize;
        xmsgMaxDataSize = numBytes;

        emit XMsgMaxDataSizeChanged(oldMax, numBytes);
    }

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function _setXReceiptMaxErrorSize(uint16 numBytes) internal {
        require(numBytes > 0, "OmniPortal: no zero max size");

        uint16 oldMax = xreceiptMaxErrorSize;
        xreceiptMaxErrorSize = numBytes;

        emit XReceiptMaxErrorSizeChanged(oldMax, numBytes);
    }

    /**
     * @notice Set the fee oracle
     */
    function _setFeeOracle(address feeOracle_) private {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");

        address oldFeeOracle = feeOracle;
        feeOracle = feeOracle_;

        emit FeeOracleChanged(oldFeeOracle, feeOracle);
    }
}
