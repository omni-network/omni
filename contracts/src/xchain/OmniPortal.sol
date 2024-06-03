// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-upgrades/contracts/security/PausableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import { ExcessivelySafeCall } from "@nomad-xyz/excessively-safe-call/src/ExcessivelySafeCall.sol";

import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Quorum } from "../libraries/Quorum.sol";
import { XRegistryNames } from "../libraries/XRegistryNames.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";

import { OmniPortalConstants } from "./OmniPortalConstants.sol";
import { OmniPortalStorage } from "./OmniPortalStorage.sol";

contract OmniPortal is
    IOmniPortal,
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
     * @notice Initialize the OmniPortal contract
     * @param owner_                    The owner of the contract
     * @param feeOracle_                Address of the fee oracle contract
     * @param xregistry_                Address of the xregistry replica contract
     * @param omniChainId_              Chain ID of Omni's EVM execution chain
     * @param omniCChainID_             Virtual chain ID used in xmsgs from Omni's consensus chain
     * @param xmsgDefaultGasLimit_      Default gas limit for xmsg
     * @param xmsgMaxGasLimit_          Maximum gas limit for xmsg
     * @param xmsgMinGasLimit_          Minimum gas limit for xmsg
     * @param xreceiptMaxErrorBytes_    Maximum error bytes for xreceipt)
     * @param valSetId                  Initial validator set id
     * @param validators                Initial validator set
     */
    function initialize(
        address owner_,
        address feeOracle_,
        address xregistry_,
        uint64 omniChainId_,
        uint64 omniCChainID_,
        uint64 xmsgDefaultGasLimit_,
        uint64 xmsgMaxGasLimit_,
        uint64 xmsgMinGasLimit_,
        uint16 xreceiptMaxErrorBytes_,
        uint64 valSetId,
        XTypes.Validator[] memory validators
    ) public initializer {
        _transferOwnership(owner_);
        _setFeeOracle(feeOracle_);
        _setXRegistry(xregistry_);
        _setXMsgDefaultGasLimit(xmsgDefaultGasLimit_);
        _setXMsgMaxGasLimit(xmsgMaxGasLimit_);
        _setXMsgMinGasLimit(xmsgMinGasLimit_);
        _setXReceiptMaxErrorBytes(xreceiptMaxErrorBytes_);
        _addValidatorSet(valSetId, validators);

        omniChainId = omniChainId_;
        omniCChainID = omniCChainID_;

        // cchain xmsg & xblock offsets are equal to valSetId
        // omni cchain is Finalized only
        inXMsgOffset[omniCChainID_][ConfLevel.Finalized] = valSetId;
        inXBlockOffset[omniCChainID_][ConfLevel.Finalized] = valSetId;

        // initialize omniChainId valSetId - xmsgs from omni are required to initSourceChain
        // omni chain is Finalized only
        inXStreamValidatorSetId[omniChainId_][ConfLevel.Finalized] = valSetId;

        // initialize omniCChainID valSetId - it is not initialized via initSourceChain
        inXStreamValidatorSetId[omniCChainID_][ConfLevel.Finalized] = valSetId;
    }

    function chainId() public view returns (uint64) {
        return uint64(block.chainid);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Outbound xcall functions                            //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Call a contract on another chain.
     *          (Default xmsg gas limit, ConfLevel.Finalized)
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable whenNotPaused {
        _xcall(destChainId, msg.sender, to, data, xmsgDefaultGasLimit, ConfLevel.Finalized);
    }

    /**
     * @notice Call a contract on another chain,
     *          (Default xmsg gas limit, explicit ConfLevel)
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param conf          Confirmation level
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint8 conf) external payable whenNotPaused {
        _xcall(destChainId, msg.sender, to, data, xmsgDefaultGasLimit, conf);
    }

    /**
     * @notice Call a contract on another.
     *           (Explcit gas limit , ConfLevel.Finalized)
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit)
        external
        payable
        whenNotPaused
    {
        _xcall(destChainId, msg.sender, to, data, gasLimit, ConfLevel.Finalized);
    }

    /**
     * @notice Call a contract on another chain.
     *          (Explicit gas limit, explicit ConfLevel)
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     * @param conf          Confirmation level
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit, uint8 conf)
        external
        payable
        whenNotPaused
    {
        _xcall(destChainId, msg.sender, to, data, gasLimit, conf);
    }

    /**
     * @notice Calculate the fee for calling a contract on another chain. Uses xmsgDefaultGasLimit.
     *         Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, xmsgDefaultGasLimit);
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

    /**
     * @notice Initiate an xcall.
     * @dev Validate the xcall, emit an XMsg, increment dest chain outXStreamOffset
     */
    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit, uint8 conf)
        private
    {
        // Only support ConfLevel.Finalized and ConfLevel.Fast for now
        require(conf == ConfLevel.Finalized || conf == ConfLevel.Fast, "OmniPortal: invalid conf level");
        require(destChainId != chainId(), "OmniPortal: no same-chain xcall");
        require(destChainId != _BROADCAST_CHAIN_ID, "OmniPortal: no broadcast xcall");
        require(isSupportedChain(destChainId), "OmniPortal: unsupported chain");
        require(to != _VIRTUAL_PORTAL_ADDRESS, "OmniPortal: no portal xcall");
        require(gasLimit <= xmsgMaxGasLimit, "OmniPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "OmniPortal: gasLimit too low");
        require(msg.value >= feeFor(destChainId, data, gasLimit), "OmniPortal: insufficient fee");

        // conf level will always be first byte of shardId. for now, shardId is just conf level
        uint64 shardId = uint64(conf);

        outXMsgOffset[destChainId][shardId] += 1;

        emit XMsg(destChainId, shardId, outXMsgOffset[destChainId][shardId], sender, to, data, gasLimit, msg.value);
    }

    /**
     * @notice Returns true if `destChainId` is supported destination chain.
     */
    function isSupportedChain(uint64 destChainId) public view returns (bool) {
        return destChainId != chainId()
            && XRegistryBase(xregistry).has(destChainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry);
    }

    /**
     * @notice Initialize a source chain's in stream validator set
     */
    function initSourceChain(uint64 srcChainId) external {
        require(msg.sender == xregistry, "OmniPortal: only xregistry");

        // We initialize for finalized and latset conf level shards.
        // TODO: update XRegistry & XRegistryReplica apps to include supported shards in the chain's metadata

        uint64 latestOmniValSetId = inXStreamValidatorSetId[omniChainId][ConfLevel.Finalized];
        inXStreamValidatorSetId[srcChainId][ConfLevel.Finalized] = latestOmniValSetId;
        inXStreamValidatorSetId[srcChainId][ConfLevel.Latest] = latestOmniValSetId;
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

        require(xmsgs.length > 0, "OmniPortal: no xmsgs");

        // validator set id for this submission
        uint64 valSetId = xsub.validatorSetId;

        // check that the validator set is known and has non-zero power
        require(validatorSetTotalPower[valSetId] > 0, "OmniPortal: unknown val set");

        // last seen validator set id for this source chain
        uint64 lastValSetId = inXStreamValidatorSetId[xheader.sourceChainId][xheader.confLevel];

        // require the validator set id is initialized (initSourceChain has beed called)
        require(lastValSetId > 0, "OmniPortal: no val set");

        // check that the submission's validator set is the same as the last, or the next one
        require(valSetId >= lastValSetId, "OmniPortal: old val set");

        // check that the attestationRoot is signed by a quorum of validators in xsub.validatorsSetId
        require(
            Quorum.verify(
                xsub.attestationRoot,
                xsub.signatures,
                validatorSet[valSetId],
                validatorSetTotalPower[valSetId],
                XSUB_QUORUM_NUMERATOR,
                XSUB_QUORUM_DENOMINATOR
            ),
            "OmniPortal: no quorum"
        );

        // check that blockHeader and xmsgs are included in attestationRoot
        require(
            XBlockMerkleProof.verify(xsub.attestationRoot, xsub.blockHeader, xsub.msgs, xsub.proof, xsub.proofFlags),
            "OmniPortal: invalid proof"
        );

        // source chain block height of this submission
        uint64 xblockOffset = xsub.blockHeader.offset;

        // last seen xblock offset for this source chain
        uint64 lastXBlockOffset = inXBlockOffset[xheader.sourceChainId][xheader.confLevel];

        // update in stream block height, if it's new
        if (xblockOffset > lastXBlockOffset) inXBlockOffset[xheader.sourceChainId][xheader.confLevel] = xblockOffset;

        // update in stream validator set id, if it's new
        if (valSetId > lastValSetId) inXStreamValidatorSetId[xheader.sourceChainId][xheader.confLevel] = valSetId;

        XTypes.Msg calldata xmsg_;

        // execute xmsgs
        for (uint256 i = 0; i < xmsgs.length; i++) {
            xmsg_ = xsub.msgs[i];

            // TODO: we can remove xmsg sourceChainId, and instead set _xmsg.sourceChainId to xsub.blockHeader.sourceChainId
            require(xmsg_.sourceChainId == xsub.blockHeader.sourceChainId, "OmniPortal: wrong sourceChainId");

            // require xmsg conf level matches block header conf level
            require(uint8(xmsg_.shardId) == xsub.blockHeader.confLevel, "OmniPortal: wrong confLevel");

            _exec(xmsg_);
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
    function _exec(XTypes.Msg calldata xmsg_) internal {
        require(
            xmsg_.destChainId == chainId() || xmsg_.destChainId == _BROADCAST_CHAIN_ID, "OmniPortal: wrong destChainId"
        );
        require(xmsg_.offset == inXMsgOffset[xmsg_.sourceChainId][xmsg_.shardId] + 1, "OmniPortal: wrong offset");

        inXMsgOffset[xmsg_.sourceChainId][xmsg_.shardId] += 1;

        // do not allow user xcalls to the portal
        // only sys xcalls (to _VIRTUAL_PORTAL_ADDRESS) are allowed to be executed on the portal
        if (xmsg_.to == address(this)) {
            emit XReceipt(
                xmsg_.sourceChainId,
                xmsg_.shardId,
                xmsg_.offset,
                0,
                msg.sender,
                false,
                abi.encodeWithSignature("Error(string)", "OmniPortal: no xcall to portal")
            );
            return;
        }

        // set _xmsg to the one we're executing
        _xmsg = XTypes.MsgShort(xmsg_.sourceChainId, xmsg_.sender);

        // xcalls to _VIRTUAL_PORTAL_ADDRESS are system calls
        bool isSysCall = xmsg_.to == _VIRTUAL_PORTAL_ADDRESS;

        (bool success, bytes memory result, uint256 gasUsed) =
            isSysCall ? _execSys(xmsg_.data) : _exec(xmsg_.to, xmsg_.gasLimit, xmsg_.data);

        // reset xmsg to zero
        delete _xmsg;

        // empty error if success is true
        bytes memory errorMsg = success ? bytes("") : result;

        emit XReceipt(xmsg_.sourceChainId, xmsg_.shardId, xmsg_.offset, gasUsed, msg.sender, success, errorMsg);
    }

    /**
     * @notice Execute a call at `to` with `data`, enfocring `gasLimit`. Returns success (true/false) and gasUsed.
     *         Requires that enough gas is left to execute the call.
     */
    function _exec(address to, uint64 gasLimit, bytes calldata data) internal returns (bool, bytes memory, uint256) {
        uint256 gasLeftBefore = gasleft();

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) =
            to.excessivelySafeCall({ _gas: gasLimit, _value: 0, _maxCopy: xreceiptMaxErrorBytes, _calldata: data });

        uint256 gasLeftAfter = gasleft();

        // Esnure relayer sent enough gas for the call
        // See https://github.com/OpenZeppelin/openzeppelin-contracts/blob/bd325d56b4c62c9c5c1aff048c37c6bb18ac0290/contracts/metatx/MinimalForwarder.sol#L58-L68
        if (gasLeftAfter <= gasLimit / 63) {
            // We could use invalid opcode to consume all gas and bubble-up the effects, since
            // and emulate an "OutOfGas" exception
            assembly {
                invalid()
            }
        }

        return (success, result, gasLeftBefore - gasLeftAfter);
    }

    /**
     * @notice Execute a system call with `data` at this contract, returning success and gasUsed.
     *         System calls must succeed.
     */
    function _execSys(bytes calldata data) internal returns (bool, bytes memory, uint256) {
        uint256 gasUsed = gasleft();

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) = address(this).call(data);

        gasUsed = gasUsed - gasleft();

        // if not success, revert with same reason
        if (!success) {
            // solhint-disable-next-line no-inline-assembly
            assembly {
                revert(add(result, 32), mload(result))
            }
        }

        return (success, result, gasUsed);
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
     * @notice Set the XRegistry replica contract
     */
    function setXRegistry(address xregistry_) external onlyOwner {
        _setXRegistry(xregistry_);
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
     * @notice Set the default gas limit for xmsg
     */
    function setXMsgDefaultGasLimit(uint64 gasLimit) external onlyOwner {
        _setXMsgDefaultGasLimit(gasLimit);
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
    function setXReceiptMaxErrorBytes(uint16 maxErrorBytes) external onlyOwner {
        _setXReceiptMaxErrorBytes(maxErrorBytes);
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
     * @notice Set the default gas limit for xmsg
     */
    function _setXMsgDefaultGasLimit(uint64 gasLimit) internal {
        require(gasLimit > 0, "OmniPortal: no zero default gas");

        uint64 oldDefault = xmsgDefaultGasLimit;
        xmsgDefaultGasLimit = gasLimit;

        emit XMsgDefaultGasLimitChanged(oldDefault, gasLimit);
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
     * @notice Set the maximum error bytes for xreceipt
     */
    function _setXReceiptMaxErrorBytes(uint16 maxErrorBytes) internal {
        require(maxErrorBytes > 0, "OmniPortal: no zero max bytes");

        uint16 oldMax = xreceiptMaxErrorBytes;
        xreceiptMaxErrorBytes = maxErrorBytes;

        emit XReceiptMaxErrorBytesChanged(oldMax, maxErrorBytes);
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

    /**
     * @notice Set the xregistry replica contract address.
     */
    function _setXRegistry(address xregistry_) private {
        require(xregistry_ != address(0), "OmniPortal: no zero xregistry");
        xregistry = xregistry_;
    }

    /**
     * @notice Add a new validator set.
     * @dev Only callable via xcall from Omni's consensus chain
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function addValidatorSet(uint64 valSetId, XTypes.Validator[] memory validators) external {
        require(msg.sender == address(this), "OmniPortal: only self");
        require(_xmsg.sourceChainId == omniCChainID, "OmniPortal: only cchain");
        require(_xmsg.sender == _CCHAIN_SENDER, "OmniPortal: only cchain sender");
        _addValidatorSet(valSetId, validators);
    }

    /**
     * @notice Add a new validator set
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function _addValidatorSet(uint64 valSetId, XTypes.Validator[] memory validators) private {
        uint256 numVals = validators.length;
        require(numVals > 0, "OmniPortal: no validators");

        uint64 totalPower;
        XTypes.Validator memory val;
        mapping(address => uint64) storage valSet = validatorSet[valSetId];

        for (uint256 i = 0; i < numVals; i++) {
            val = validators[i];

            require(val.addr != address(0), "OmniPortal: no zero validator");
            require(val.power > 0, "OmniPortal: no zero power");
            require(valSet[val.addr] == 0, "OmniPortal: duplicate validator");

            totalPower += val.power;
            valSet[val.addr] = val.power;
        }

        validatorSetTotalPower[valSetId] = totalPower;

        emit ValidatorSetAdded(valSetId);
    }
}
