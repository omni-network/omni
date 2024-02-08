// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Validators } from "../libraries/Validators.sol";

contract OmniPortal is IOmniPortal, IOmniPortalAdmin, Ownable {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @inheritdoc IOmniPortal
    uint8 public constant XSUB_QUORUM_THRESHOLD_NUMERATOR = 66;

    /// @inheritdoc IOmniPortal
    uint8 public constant XSUB_QUORUM_THRESHOLD_DENOMINATOR = 100;

    /// @inheritdoc IOmniPortal
    uint64 public immutable chainId;

    /// @inheritdoc IOmniPortalAdmin
    address public feeOracle;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public outXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamBlockHeight;

    /// @dev Track latest seen valSetId, to avoid writing the same validator set multiple times
    ///      validator set ideas increment monotonically
    uint64 private _latestValSetId;

    /// @dev Maps validator set id -> validator address -> power
    mapping(uint64 => mapping(address => uint64)) private _validatorSet;

    /// @dev Maps validator set id -> total power
    mapping(uint64 => uint64) private _validatorSetTotalPower;

    /// @dev The current XMsg being executed, exposed via xmsg() getter
    ///      Private state + public getter preferred over public state with default getter,
    ///      so that we can use the XMsg struct type in the interface.
    XTypes.Msg private _currentXmsg;

    constructor(address owner_, address feeOracle_, uint64 valSetId, Validators.Validator[] memory validators_)
        Ownable(owner_)
    {
        chainId = uint64(block.chainid);
        _setFeeOracle(feeOracle_);
        _addValidators(valSetId, validators_);
    }

    /// @inheritdoc IOmniPortalAdmin
    function setFeeOracle(address feeOracle_) external onlyOwner {
        _setFeeOracle(feeOracle_);
    }

    /// @inheritdoc IOmniPortalAdmin
    function collectFees(address to) external onlyOwner {
        uint256 amount = address(this).balance;

        // .transfer() is fine, owner should provide an EOA address that will not
        // consume more than 2300 gas on transfer, and we are okay .transfer() reverts
        payable(to).transfer(amount);

        emit FeesCollected(to, amount);
    }

    /// @inheritdoc IOmniPortal
    function xmsg() external view returns (XTypes.Msg memory) {
        return _currentXmsg;
    }

    /// @inheritdoc IOmniPortal
    function isXCall() external view returns (bool) {
        return _currentXmsg.sourceChainId != 0;
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64 destChainId, bytes calldata data) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, gasLimit);
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable {
        _xcall(destChainId, msg.sender, to, data, gasLimit);
    }

    /// @inheritdoc IOmniPortal
    function xsubmit(XTypes.Submission calldata xsub) external {
        // This only works because validator sets cannot yet change
        // TODO: replace with uint64 valSetId = xsub.valSetId;
        uint64 valSetId = _latestValSetId;

        require(
            Validators.verifyQuorum(
                xsub.attestationRoot,
                xsub.signatures,
                _validatorSet[valSetId],
                _validatorSetTotalPower[valSetId],
                XSUB_QUORUM_THRESHOLD_NUMERATOR,
                XSUB_QUORUM_THRESHOLD_DENOMINATOR
            ),
            "OmniPortal: no quorum"
        );

        require(
            XBlockMerkleProof.verify(xsub.attestationRoot, xsub.blockHeader, xsub.msgs, xsub.proof, xsub.proofFlags),
            "OmniPortal: invalid proof"
        );

        inXStreamBlockHeight[xsub.blockHeader.sourceChainId] = xsub.blockHeader.blockHeight;

        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            _exec(xsub.msgs[i]);
        }
    }

    /// @dev Emit an XMsg event, increment dest chain outXStreamOffset
    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit) private {
        require(msg.value >= feeFor(destChainId, data, gasLimit), "OmniPortal: insufficient fee");
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");
        require(destChainId != chainId, "OmniPortal: no same-chain xcall");

        outXStreamOffset[destChainId] += 1;

        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
    }

    /// @dev Verify an XMsg is next in its XStream, execute it, increment inXStreamOffset, emit an XReceipt
    function _exec(XTypes.Msg calldata xmsg_) internal {
        require(xmsg_.destChainId == chainId, "OmniPortal: wrong destChainId");
        require(xmsg_.streamOffset == inXStreamOffset[xmsg_.sourceChainId] + 1, "OmniPortal: wrong streamOffset");

        // set xmsg to the one we're executing
        _currentXmsg = xmsg_;

        // increment offset before executing xcall, to avoid reentrancy loop
        inXStreamOffset[xmsg_.sourceChainId] += 1;

        // we enforce a maximum on xcall, but we trim to max here just in case
        uint256 gasLimit = xmsg_.gasLimit > XMSG_MAX_GAS_LIMIT ? XMSG_MAX_GAS_LIMIT : xmsg_.gasLimit;

        // execute xmsg, tracking gas used
        uint256 gasUsed = gasleft();
        (bool success,) = xmsg_.to.call{ gas: gasLimit }(xmsg_.data);
        gasUsed = gasUsed - gasleft();

        // reset xmsg to zero
        _currentXmsg = XTypes.zeroMsg();

        emit XReceipt(xmsg_.sourceChainId, xmsg_.streamOffset, gasUsed, msg.sender, success);
    }

    /// @dev Set the fee oracle
    function _setFeeOracle(address feeOracle_) internal {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");

        address oldFeeOracle = feeOracle;
        feeOracle = feeOracle_;

        emit FeeOracleChanged(oldFeeOracle, feeOracle);
    }

    function _addValidators(uint64 valSetId, Validators.Validator[] memory validators_) internal {
        require(valSetId == _latestValSetId + 1, "OmniPortal: invalid valSetId");
        require(validators_.length > 0, "OmniPortal: no validators");

        uint64 totalPower;
        Validators.Validator memory val;
        mapping(address => uint64) storage set = _validatorSet[valSetId];

        for (uint256 i = 0; i < validators_.length; i++) {
            val = validators_[i];

            require(val.addr != address(0), "OmniPortal: no zero validator");
            require(val.power > 0, "OmniPortal: no zero power");

            totalPower += val.power;
            set[val.addr] = val.power;
        }

        _validatorSetTotalPower[valSetId] = totalPower;
        _latestValSetId = valSetId;
    }
}
