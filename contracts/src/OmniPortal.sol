// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "./interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "./interfaces/IOmniPortalAdmin.sol";
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
    uint256 public fee;

    /// @dev The current XMsg being executed, exposed via xmsg() getter
    ///      Private state + public getter preferred over public state with default getter,
    ///      so that we can use the XMsg struct type in the interface.
    XTypes.Msg private _currentXmsg;

    modifier onlyAdmin() {
        require(msg.sender == admin, "OmniPortal: not admin");
        _;
    }

    constructor(address admin_, uint256 fee_) {
        chainId = uint64(block.chainid);

        _setAdmin(admin_);
        _setFee(fee_);
    }

    /// @inheritdoc IOmniPortalAdmin
    function setAdmin(address admin_) external onlyAdmin {
        _setAdmin(admin_);
    }

    function _setAdmin(address admin_) internal {
        require(admin_ != address(0), "OmniPortal: no zero admin");
        admin = admin_;
    }

    /// @inheritdoc IOmniPortalAdmin
    function setFee(uint256 fee_) external onlyAdmin {
        _setFee(fee_);
    }

    function _setFee(uint256 fee_) internal {
        fee = fee_;
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
    function feeFor(uint64, bytes calldata) external view returns (uint256) {
        return fee;
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64, bytes calldata, uint64) external view returns (uint256) {
        return fee;
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
}
