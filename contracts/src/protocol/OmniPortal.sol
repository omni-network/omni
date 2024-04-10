// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Quorum } from "../libraries/Quorum.sol";

import { OmniPortalConstants } from "./OmniPortalConstants.sol";
import { OmniPortalStorage } from "./OmniPortalStorage.sol";

contract OmniPortal is IOmniPortal, IOmniPortalAdmin, OwnableUpgradeable, OmniPortalConstants, OmniPortalStorage {
    /**
     * @notice Chain ID of the chain to which this portal is deployed
     */
    uint64 public immutable chainId;

    /**
     * @notice Construct the OmniPortal contract
     */
    constructor() {
        _disableInitializers();
        chainId = uint64(block.chainid);
    }

    /**
     * @notice Initialize the OmniPortal contract
     * @param owner_        The owner of the contract
     * @param feeOracle_    Address of the fee oracle contract
     * @param omniEChainId_ Chain ID of Omni's EVM execution chain
     * @param omniCChainID_ Virtual chain ID used in xmsgs from Omni's consensus chain
     * @param valSetId      Initial validator set id
     * @param validators    Initial validator set
     */
    function initialize(
        address owner_,
        address feeOracle_,
        uint64 omniEChainId_,
        uint64 omniCChainID_,
        uint64 valSetId,
        XTypes.Validator[] memory validators
    ) public initializer {
        __Ownable_init();
        _transferOwnership(owner_);
        _setFeeOracle(feeOracle_);
        _addValidatorSet(valSetId, validators);

        omniEChainId = omniEChainId_;
        omniCChainID = omniCChainID_;

        // cchain stream offset & block heights are equal to valSetId
        inXStreamOffset[omniCChainID_] = valSetId;
        inXStreamBlockHeight[omniCChainID_] = valSetId;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Outbound xcall functions                            //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Call a contract on another chain Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution
     *         gas limit on destination chain
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /**
     * @notice Call a contract on another chain Uses provide gasLimit as execution gas limit on
     *          destination chain. Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit >
     *          XMSG_MAX_GAS_LIMIT
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable {
        _xcall(destChainId, msg.sender, to, data, gasLimit);
    }

    /**
     * @notice Calculate the fee for calling a contract on another chain. Uses
     *         OmniPortal.XMSG_DEFAULT_GAS_LIMIT. Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, XMSG_DEFAULT_GAS_LIMIT);
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
    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit) private {
        require(msg.value >= feeFor(destChainId, data, gasLimit), "OmniPortal: insufficient fee");
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");
        require(destChainId != chainId, "OmniPortal: no same-chain xcall");
        require(destChainId != _BROADCAST_CHAIN_ID, "OmniPortal: no broadcast xcall");
        require(to != _VIRTUAL_PORTAL_ADDRESS, "OmniPortal: no portal xcall");

        outXStreamOffset[destChainId] += 1;

        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Inbound xcall functions                             //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Calculate the gas required to submit a batch of XMsgs to this chain.
     *         gasForSubmission(xsub) must be >= the gas required to execute the submission.
     *         Implementaion matches xsubmit, but replaces storage writes and xmsg exeuction with
     *         gas additions.
     */
    function estimateGas(XTypes.Submission calldata xsub) external view returns (uint256) {
        // track gas used during estimate
        uint256 gasStart = gasleft();

        // additional gas, to account for operations that are not included in the gas estimate
        uint256 gasAdditions = 0;

        // verify the submission
        _verifiy(xsub);

        // validator set id for this submission
        uint64 valSetId = xsub.validatorSetId;

        // last seen validator set id for this source chain
        uint64 lastValSetId = inXStreamValidatorSetId[xsub.blockHeader.sourceChainId];

        // source chain block height of this submission
        uint64 blockHeight = xsub.blockHeader.blockHeight;

        // last seen block height for this source chain
        uint64 lastBlockHeight = inXStreamBlockHeight[xsub.blockHeader.sourceChainId];

        // update in stream block height, if it's new
        if (blockHeight > lastBlockHeight) {
            // 20,000 gas for new SSSTORe, 2,900 gas for overwtie SSTORE
            gasAdditions += lastBlockHeight == 0 ? 20_000 : 2900;
        }

        // update in stream validator set id, if it's new
        if (valSetId > lastValSetId) {
            // 20,000 gas for new SSSTORE, 2,900 gas for overwtie SSTORE
            gasAdditions += lastValSetId == 0 ? 20_000 : 2900;
        }

        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            XTypes.Msg calldata xmsg_ = xsub.msgs[i];

            // add each xmsg's gasLimit
            // sys xcalls do not have gas limit, so we estimate their gas
            gasAdditions += _isSysCall(xmsg_) ? _estimateGas(xmsg_.data) : xmsg_.gasLimit;

            // per-xmsg overhead, to account _exec gas usage (SLOAD / STORE offset, checks, etc.)
            gasAdditions += 20_000;
        }

        // general overhead, and gas missed in estimate
        // we'd rather overestimate than underestimate
        gasAdditions += 20_000;

        return (gasStart - gasleft()) + gasAdditions;
    }

    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub  An xchain submisison, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external {
        // verify the submission
        _verifiy(xsub);

        // validator set id for this submission
        uint64 valSetId = xsub.validatorSetId;

        // last seen validator set id for this source chain
        uint64 lastValSetId = inXStreamValidatorSetId[xsub.blockHeader.sourceChainId];

        // source chain block height of this submission
        uint64 blockHeight = xsub.blockHeader.blockHeight;

        // last seen block height for this source chain
        uint64 lastBlockHeight = inXStreamBlockHeight[xsub.blockHeader.sourceChainId];

        // update in stream block height, if it's new
        if (blockHeight > lastBlockHeight) inXStreamBlockHeight[xsub.blockHeader.sourceChainId] = blockHeight;

        // update in stream validator set id, if it's new
        if (valSetId > lastValSetId) inXStreamValidatorSetId[xsub.blockHeader.sourceChainId] = valSetId;

        // execute xmsgs
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            _exec(xsub.msgs[i]);
        }
    }

    /**
     * @notice Verify an xsbmission. Reverts if the submission is invalid.
     */
    function _verifiy(XTypes.Submission calldata xsub) internal view {
        require(xsub.msgs.length > 0, "OmniPortal: no xmsgs");

        // validator set id for this submission
        uint64 valSetId = xsub.validatorSetId;

        // last seen validator set id for this source chain
        uint64 lastValSetId = inXStreamValidatorSetId[xsub.blockHeader.sourceChainId];

        // check that the validator set is known and has non-zero power
        require(validatorSetTotalPower[valSetId] > 0, "OmniPortal: unknown val set");

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
            xmsg_.destChainId == chainId || xmsg_.destChainId == _BROADCAST_CHAIN_ID, "OmniPortal: wrong destChainId"
        );
        require(xmsg_.streamOffset == inXStreamOffset[xmsg_.sourceChainId] + 1, "OmniPortal: wrong streamOffset");

        // set _xmsg to the one we're executing
        _xmsg = XTypes.MsgShort(xmsg_.sourceChainId, xmsg_.sender);

        // increment offset before executing xcall, to avoid reentrancy loop
        inXStreamOffset[xmsg_.sourceChainId] += 1;

        (bool success, uint256 gasUsed) =
            _isSysCall(xmsg_) ? _execSys(xmsg_.data) : _exec(xmsg_.to, xmsg_.gasLimit, xmsg_.data);

        // reset xmsg to zero
        delete _xmsg;

        emit XReceipt(xmsg_.sourceChainId, xmsg_.streamOffset, gasUsed, msg.sender, success);
    }

    /**
     * @notice Execute a call at `to` with `data`, enfocring `gasLimit`. Returns success (true/false) and gasUsed.
     *         Requires that enough gas is left to execute the call.
     */
    function _exec(address to, uint64 gasLimit, bytes calldata data) internal returns (bool, uint256) {
        // trim gasLimit to max. this requirement is checked in xcall(...), but we trim here to be safe
        if (gasLimit > XMSG_MAX_GAS_LIMIT) gasLimit = XMSG_MAX_GAS_LIMIT;

        // require gasLeft is enough to execute the call. this protects against malicious relayers
        // purposefully setting gasLimit just low enough such that the last xmsg in a submission
        // fails, despite it's sufficient gasLimit
        //
        // We add a small buffer to account for the gas usage from here up until the call.
        // TODO: is buffer of 100 correct? Better more than less
        require(gasLimit + 100 < gasleft(), "OmniPortal: gasLimit too low");

        uint256 gasUsed = gasleft();

        // solhint-disable-next-line avoid-low-level-calls
        (bool success,) = to.call{ gas: gasLimit }(data);

        gasUsed = gasUsed - gasleft();

        return (success, gasUsed);
    }

    /**
     * @notice Execute a system call with `data` at this contract, returning success and gasUsed.
     *         System calls must succeed.
     */
    function _execSys(bytes calldata data) internal returns (bool, uint256) {
        uint256 gasUsed = gasleft();

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) = address(this).call(data);

        gasUsed = gasUsed - gasleft();

        // if not success, revert with the error message
        if (!success) revert(_revertReason(result));

        return (success, gasUsed);
    }

    /**
     * @notice Returns the revert reason from an address.call result.
     * @dev Only works for address.call() that were unsuccessful, and reverted with a reason.
     * @custom:attriubtion  https://github.com/Uniswap/v3-periphery/blob/v1.0.0/contracts/base/Multicall.sol#L17
     * @custom:attriubtion  https://ethereum.stackexchange.com/a/83577
     * @param result    The result of an address.call
     */
    function _revertReason(bytes memory result) internal pure returns (string memory) {
        if (result.length < 68) return "no revert reason";

        // solhint-disable-next-line no-inline-assembly
        assembly {
            result := add(result, 0x04)
        }

        return abi.decode(result, (string));
    }

    /**
     * @notice Returns true if the xmsg is a system call, false otherwise
     */
    function _isSysCall(XTypes.Msg calldata xmsg_) internal pure returns (bool) {
        return xmsg_.to == _VIRTUAL_PORTAL_ADDRESS;
    }

    /**
     * @notice Estimate the gas used by a system call with `data`
     */
    function _estimateGas(bytes calldata data) internal view returns (uint256) {
        bytes4 sig = bytes4(data[:4]);
        bytes memory args = data[4:];

        if (sig == this.addValidatorSet.selector) {
            (uint64 valSetId, XTypes.Validator[] memory validators) = abi.decode(args, (uint64, XTypes.Validator[]));
            return _estimateGas_addValidatorSet(valSetId, validators);
        }

        revert("OmniPortal: invalid sys call");
    }

    /**
     * @notice Estimate the gas used by a system call with `data`. Matches addValidatorSet
     *         implementation, but replaces storage writes with gas additions.
     */
    // solhint-disable-next-line func-name-mixedcase
    function _estimateGas_addValidatorSet(uint64 valSetId, XTypes.Validator[] memory validators)
        internal
        view
        returns (uint256)
    {
        uint256 gasStart = gasleft();
        uint256 gasAdditions;

        require(validators.length > 0, "OmniPortal: no validators");

        uint64 totalPower;
        XTypes.Validator memory val;
        mapping(address => uint64) storage valSet = validatorSet[valSetId];

        for (uint256 i = 0; i < validators.length; i++) {
            val = validators[i];

            require(val.addr != address(0), "OmniPortal: no zero validator");
            require(val.power > 0, "OmniPortal: no zero power");
            require(valSet[val.addr] == 0, "OmniPortal: duplicate validator");

            totalPower += val.power;

            // fresh SSTORE for validator
            gasAdditions += 25_000;
        }

        // fresh SSTORE for totalPower
        gasAdditions += 25_000;

        return (gasStart - gasleft()) + gasAdditions;
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
     * @notice Set the fee oracle
     */
    function _setFeeOracle(address feeOracle_) private {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");

        address oldFeeOracle = feeOracle;
        feeOracle = feeOracle_;

        emit FeeOracleChanged(oldFeeOracle, feeOracle);
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
        require(validators.length > 0, "OmniPortal: no validators");

        uint64 totalPower;
        XTypes.Validator memory val;
        mapping(address => uint64) storage valSet = validatorSet[valSetId];

        for (uint256 i = 0; i < validators.length; i++) {
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
