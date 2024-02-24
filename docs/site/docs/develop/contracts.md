---
sidebar_position: 2
---

# Contracts Overview

The Omni protocol is implemented through a set of contracts that define the core functionality of the Omni network. These contracts are deployed to all integrated rollups and provide the necessary interfaces for cross-rollup communication and also available as utility contracts for dApp developers.

## Portal Contract

### Interface

#### [`IOmniPortal.sol`](https://github.com/omni-network/omni/blob/2d1e3f57c140b8824bf7d39244e0168fec73af4c/contracts/src/interfaces/IOmniPortal.sol)

- A utility contract aimed at simplifying interactions with the Omni [Portal Contact](../use/protocol.md#portal-contract).
- On-chain gateway for Omni's cross-chain messaging, enabling cross-chain contract invocations through events like `XMsg` and `XReceipt`.
- It establishes gas limits for cross-chain message execution, sets quorum thresholds for validator approvals, and maintains a record of the chain ID as well as inbound and outbound message offsets.
- Provides methods to [calculate fees](#fees) for cross-chain function calls, both with default and custom gas limits, allowing payable cross-chain calls (`xcall`) and submission of batches of cross-chain messages (`xsubmit`).

<details>
<summary>`IOmniPortal.sol` Reference Solidity Code</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title IOmniPortal
 * @notice The OmniPortal is the on-chain interface to Omni's cross-chain
 *         messaging protocol. It is used to initiate and execute cross-chain calls.
 */
interface IOmniPortal {
    /**
     * @notice Emitted when an xcall is made to a contract on another chain
     * @param destChainId Destination chain ID
     * @param streamOffset Offset this XMsg in the source -> dest XStream
     * @param sender msg.sender of the source xcall
     * @param to Address of the contract to call on the destination chain
     * @param gasLimit Gas limit for execution on destination chain
     * @param data Encoded function calldata
     */
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );

    /**
     * @notice Emitted when an XMsg is executed on its destination chain
     * @param sourceChainId Source chain ID
     * @param streamOffset Offset the XMsg in the source -> dest XStream
     * @param gasUsed Gas used in execution of the XMsg
     * @param success Whether the execution succeeded
     * @param relayer Address of the relayer who submitted the XMsg
     */
    event XReceipt(
        uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success
    );

    /**
     * @notice Emitted when a new validator set is added
     * @param setId Validator set ID
     */
    event ValidatorSetAdded(uint64 indexed setId);

    /**
     * @notice Default xmsg execution gas limit, enforced on destination chain
     * @return Gas limit
     */
    function XMSG_DEFAULT_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Maximum allowed xmsg gas limit
     * @return Maximum gas limit
     */
    function XMSG_MAX_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Minimum allowed xmsg gas limit
     * @return Minimum gas limit
     */
    function XMSG_MIN_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Numerator of the fraction of total validator power required to
     *         accept an XSubmission. Ex 2/3 -> 2
     * @return Quorum threshold numerator
     */
    function XSUB_QUORUM_NUMERATOR() external view returns (uint8);

    /**
     * @notice Denominator of the fraction of total validator power required to
     *         accept an XSubmission. Ex 2/3 -> 3
     * @return Quorum threshold denominator
     */
    function XSUB_QUORUM_DENOMINATOR() external view returns (uint8);

    /**
     * @notice Chain ID of the chain to which this portal is deployed
     * @dev Used as sourceChainId for all outbound XMsgs
     * @return Chain ID
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId
     * @param destChainId Destination chain ID
     * @return Offset
     */
    function outXStreamOffset(uint64 destChainId) external view returns (uint64);

    /**
     * @notice Offset of the last inbound XMsg that was received from sourceChainId
     * @param sourceChainId Source chain ID
     * @return Offset
     */
    function inXStreamOffset(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice Source block height of the last inbound XMsg that was received from sourceChainId
     * @param sourceChainId Source chain ID
     * @return Block height
     */
    function inXStreamBlockHeight(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice The current XMsg being executed via this portal
     * @dev If no XMsg is being executed, all fields will be zero
     * @return XMsg
     */
    function xmsg() external view returns (XTypes.Msg memory);

    /**
     * @notice Whether the current transaction is an xcall
     * @return True if current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     * @param gasLimit Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Call a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution gas limit on destination chain
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;

    /**
     * @notice Call a contract on another chain
     * @dev Uses provide gasLimit as execution gas limit on destination chain.
     *      Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit > XMSG_MAX_GAS_LIMIT
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable;

    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub An xchain submisison, including an attestation root w/ validator signatures,
     *        and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external;
}
```

</details>

### Deployed Contract

#### [`OmniPortal.sol`](https://github.com/omni-network/omni/blob/97cb666259c4490f9d945bcf68c1c45baece81b8/contracts/src/protocol/OmniPortal.sol)

- The Omni system [Portal Contact](../use/protocol.md#portal-contract) deployed to all integrated rollups.
- Exposes methods to interact with the Omni EVM and other rollups.
  - `function xcall(...)` sends a transaction to another rollup.
  - `function xsubmit(...)` receives calls from another rollup.

<details>
<summary>`OmniPortal.sol` Reference Solidity Code</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XBlockMerkleProof } from "../libraries/XBlockMerkleProof.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { Validators } from "../libraries/Validators.sol";

contract OmniPortal is IOmniPortal, IOmniPortalAdmin, OwnableUpgradeable {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @inheritdoc IOmniPortal
    uint8 public constant XSUB_QUORUM_NUMERATOR = 2;

    /// @inheritdoc IOmniPortal
    uint8 public constant XSUB_QUORUM_DENOMINATOR = 3;

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

    /// @dev Track latest seen valSetId, to avoid writing the same validator set multiple times.
    ///      Validator set ids increment monotonically
    uint64 private _latestValSetId;

    /// @dev Maps validator set id -> validator address -> power
    mapping(uint64 => mapping(address => uint64)) private _validatorSet;

    /// @dev Maps validator set id -> total power
    mapping(uint64 => uint64) private _validatorSetTotalPower;

    /// @dev The current XMsg being executed, exposed via xmsg() getter
    ///      Private state + public getter preferred over public state with default getter,
    ///      so that we can use the XMsg struct type in the interface.
    XTypes.Msg private _currentXmsg;

    constructor() {
        _disableInitializers();
        chainId = uint64(block.chainid);
    }

    function initialize(address owner_, address feeOracle_, uint64 valSetId, Validators.Validator[] memory validators)
        public
        initializer
    {
        __Ownable_init();
        _transferOwnership(owner_);
        _setFeeOracle(feeOracle_);
        _addValidators(valSetId, validators);
    }

    /**
     * XMsg functions
     */

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
        uint64 valSetId = _latestValSetId;

        // check that the attestationRoot is signed by a quorum of validators in xsub.validatorsSetId
        require(
            Validators.verifyQuorum(
                xsub.attestationRoot,
                xsub.signatures,
                _validatorSet[valSetId],
                _validatorSetTotalPower[valSetId],
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

        // update in stream block height
        inXStreamBlockHeight[xsub.blockHeader.sourceChainId] = xsub.blockHeader.blockHeight;

        // execute xmsgs
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            _exec(xsub.msgs[i]);
        }
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64 destChainId, bytes calldata data) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, gasLimit);
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

    /**
     * XMsg metadata functions
     */

    /// @inheritdoc IOmniPortal
    function xmsg() external view returns (XTypes.Msg memory) {
        return _currentXmsg;
    }

    /// @inheritdoc IOmniPortal
    function isXCall() external view returns (bool) {
        return _currentXmsg.sourceChainId != 0;
    }

    /**
     * Admin functions
     */

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

    /// @dev Set the fee oracle
    function _setFeeOracle(address feeOracle_) private {
        require(feeOracle_ != address(0), "OmniPortal: no zero feeOracle");

        address oldFeeOracle = feeOracle;
        feeOracle = feeOracle_;

        emit FeeOracleChanged(oldFeeOracle, feeOracle);
    }

    function _addValidators(uint64 valSetId, Validators.Validator[] memory validators) private {
        require(valSetId == _latestValSetId + 1, "OmniPortal: invalid valSetId");
        require(validators.length > 0, "OmniPortal: no validators");

        uint64 totalPower;
        Validators.Validator memory val;
        mapping(address => uint64) storage set = _validatorSet[valSetId];

        for (uint256 i = 0; i < validators.length; i++) {
            val = validators[i];

            require(val.addr != address(0), "OmniPortal: no zero validator");
            require(val.power > 0, "OmniPortal: no zero power");

            totalPower += val.power;
            set[val.addr] = val.power;
        }

        _validatorSetTotalPower[valSetId] = totalPower;
        _latestValSetId = valSetId;

        emit ValidatorSetAdded(valSetId);
    }
}
```

</details>

## Types Contract

#### [`XTypes.sol`](https://github.com/omni-network/omni/blob/2d1e3f57c140b8824bf7d39244e0168fec73af4c/contracts/src/libraries/XTypes.sol)

- Defines core xchain messaging types for the Omni protocol.
- Includes the `Msg` structure, which holds details of cross-chain messages such as source and destination chain IDs, a monotonically increasing stream offset, sender and target addresses, call data, and gas limit for execution.
- Contains a `BlockHeader` struct for source chain block identification, including chain ID, block height, and hash.
- Features a `Submission` struct that encapsulates a merkle root (attestationRoot) signed by validators, the validator set ID, the block header, an array of `Msg` objects to be executed, a merkle proof for these messages, proof flags, and corresponding validator signatures with public keys.
- Provides a `zeroMsg()` function returning a `Msg` struct with all zero values, used to represent an empty message.

<details>
<summary>`XTypes.sol` Reference Solidity Code</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { Validators } from "./Validators.sol";

/**
 * @title XTypes
 * @dev Defines xchain types, core to Omni's xchain messaging protocol. These
 *      types mirror those defined in omni/lib/xchain/types.go.
 */
library XTypes {
    struct Msg {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Chain ID of the destination chain
        uint64 destChainId;
        /// @dev Monotonically incremented offset of Msg in source -> dest Stream
        uint64 streamOffset;
        /// @dev msg.sender of xcall on source chain
        address sender;
        /// @dev Target address to call on destination chain
        address to;
        /// @dev Data to provide to call on destination chain
        bytes data;
        /// @dev Gas limit to use for call execution on destination chain
        uint64 gasLimit;
    }

    struct BlockHeader {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Height of the source chain block
        uint64 blockHeight;
        /// @dev Hash of the source chain block
        bytes32 blockHash;
    }

    struct Submission {
        /// @dev Merkle root of xchain block (XBlockRoot), attested to and signed by validators
        bytes32 attestationRoot;
        /// @dev Unique identifier of the validator set that attested to this root
        uint64 validatorSetId;
        /// @dev Block header, identifies xchain block
        BlockHeader blockHeader;
        /// @dev Messages to execute
        Msg[] msgs;
        /// @dev Multi proof of block header and messages, proven against attestationRoot
        bytes32[] proof;
        /// @dev Multi proof flags
        bool[] proofFlags;
        /// @dev Array of validator signatures of the attestationRoot, and their public keys
        Validators.SigTuple[] signatures;
    }

    /// @dev Zero value for Msg
    function zeroMsg() internal pure returns (Msg memory) {
        return Msg(0, 0, 0, address(0), address(0), "", 0);
    }
}
```

</details>
