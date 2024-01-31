// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "./interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "./interfaces/IOmniPortalAdmin.sol";
import { IFeeOracle } from "./interfaces/IFeeOracle.sol";
import { XBlockMerkleProof } from "./libraries/XBlockMerkleProof.sol";
import { XTypes } from "./libraries/XTypes.sol";

contract OmniPortal is IOmniPortal, IOmniPortalAdmin {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @inheritdoc IOmniPortal
    uint64 public immutable chainId;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public outXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamOffset;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public inXStreamBlockHeight;

    /// @inheritdoc IOmniPortalAdmin
    address public admin;

    /// @inheritdoc IOmniPortalAdmin
    address public feeOracle;

    /// @dev The current XMsg being executed, exposed via xmsg() getter
    ///      Private state + public getter preferred over public state with default getter,
    ///      so that we can use the XMsg struct type in the interface.
    XTypes.Msg private _currentXmsg;

    modifier onlyAdmin() {
        require(msg.sender == admin, "OmniPortal: only admin");
        _;
    }

    constructor(address admin_, address feeOracle_) {
        chainId = uint64(block.chainid);
        _setAdmin(admin_);
        _setFeeOracle(feeOracle_);
    }

    /// @inheritdoc IOmniPortalAdmin
    function setAdmin(address admin_) external onlyAdmin {
        _setAdmin(admin_);
    }

    /// @inheritdoc IOmniPortalAdmin
    function setFeeOracle(address feeOracle_) external onlyAdmin {
        _setFeeOracle(feeOracle_);
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
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256) {
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
        // TODO: verify a quorum of validators have signed off on the attestation root.

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

    /// @dev Set the admin account
    function _setAdmin(address admin_) internal {
        require(admin_ != address(0), "OmniPortal: no zero admin");

        address oldAdmin = admin;
        admin = admin_;

        emit AdminChanged(oldAdmin, admin);
    }

    /// @dev Set the fee oracle
    function _setFeeOracle(address feeOracle_) internal {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");

        address oldFeeOracle = feeOracle;
        feeOracle = feeOracle_; // TODO: supportsInterface check?

        emit FeeOracleChanged(oldFeeOracle, feeOracle);
    }
}
