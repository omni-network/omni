// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonBase } from "forge-std/Base.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { TestXTypes } from "./TestXTypes.sol";
import { TestPortal } from "./TestPortal.sol";
import { Counter } from "./Counter.sol";
import { Reverter } from "./Reverter.sol";

/**
 * @title Fixtures
 * @dev Defines test dependencies. Deploys portals and test contracts, and provides
 *      utilites for generating XMsgs and XSubmissions between them.
 */
contract Fixtures is CommonBase, StdCheats {
    // We introduce three "chains", this chain, chain a, and chain b. We deploy
    // portals and contracts for each chain. Obviously, they all live in the
    // test EVM state. But it's useful to semantically group contracts by
    // chain, so that we may describe and test XMsgs between them.
    //
    // All contracts on "this" chain, do not have a prefix (portal, counter, reverter).
    // That way tests can refer to them without introducing the idea of a "chain". Most
    // test cases need only think about "this" chain.

    uint64 constant thisChainId = 1;
    uint64 constant chainAId = 2;
    uint64 constant chainBId = 3;

    address deployer;
    address xcaller;
    address relayer;

    TestPortal portal;
    TestPortal chainAPortal;
    TestPortal chainBPortal;

    Counter counter;
    Counter chainACounter;
    Counter chainBCounter;

    Reverter reverter;
    Reverter chainAReverter;
    Reverter chainBReverter;

    // helper mappings to generate XMsg.to and XMsg.sender for some sourceChainId & destChainId
    mapping(uint64 => address) private _reverters;
    mapping(uint64 => address) private _counters;

    // @dev path to which test XBlocks are written (see WriteXBlocks.sol), relative to project root
    string constant XBLOCKS_PATH = "test/data/xblocks.json";

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        relayer = makeAddr("relayer");

        vm.startPrank(deployer);

        vm.chainId(thisChainId); // portal constructor uses block.chainid
        portal = new TestPortal();
        counter = new Counter();
        reverter = new Reverter();

        vm.chainId(chainAId);
        chainAPortal = new TestPortal();
        chainACounter = new Counter();
        chainAReverter = new Reverter();

        vm.chainId(chainBId);
        chainBPortal = new TestPortal();
        chainBCounter = new Counter();
        chainBReverter = new Reverter();

        vm.stopPrank();

        _counters[thisChainId] = address(counter);
        _counters[chainAId] = address(chainACounter);
        _counters[chainBId] = address(chainBCounter);

        _reverters[thisChainId] = address(reverter);
        _reverters[chainAId] = address(chainAReverter);
        _reverters[chainBId] = address(chainBReverter);
    }

    /// @dev Write test fixture XBlocks to a json file. This is allows ts utilities
    ///      to generate XSubmissions, with valid merkle roots and proofs, for each test
    ///      XBlock. These XSubmissions are then used as inputs into test cases.
    function writeXBlocks() public {
        string memory root = vm.projectRoot();
        string memory fullpath = string.concat(root, "/", XBLOCKS_PATH);

        TestXTypes.Block memory xblock1 = _xblock(1, 0); //  sourceBlockHeight: 1, startOffset: 0
        TestXTypes.Block memory xblock2 = _xblock(2, 10); // sourceBlockHeight: 2, startOffset: 10

        string memory id = "ID";
        vm.serializeBytes(id, "xblock1", abi.encode(xblock1));
        string memory json = vm.serializeBytes(id, "xblock2", abi.encode(xblock2));

        vm.writeJson(json, fullpath);
    }

    /// @dev Create an xblock from chainA with xmsgs for "this" chain and chain b.
    ///      XBlocks will likely contain XMsgs for multiple chains, so we reflect that here.
    function _xblock(uint64 sourceBlockHeight, uint64 startOffset) internal view returns (TestXTypes.Block memory) {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](10);

        // intended for this chain
        xmsgs[0] = _increment(chainAId, thisChainId, startOffset);
        xmsgs[1] = _increment(chainAId, thisChainId, startOffset + 1);
        xmsgs[2] = _increment(chainAId, thisChainId, startOffset + 2);
        xmsgs[3] = _revert(chainAId, thisChainId, startOffset + 3);
        xmsgs[4] = _increment(chainAId, thisChainId, startOffset + 4);

        // intended for chain b
        xmsgs[5] = _increment(chainAId, chainBId, startOffset + 5);
        xmsgs[6] = _increment(chainAId, chainBId, startOffset + 6);
        xmsgs[7] = _increment(chainAId, chainBId, startOffset + 7);
        xmsgs[9] = _revert(chainAId, chainBId, startOffset + 8);
        xmsgs[9] = _increment(chainAId, chainBId, startOffset + 9);

        return TestXTypes.Block(XTypes.BlockHeader(chainAId, sourceBlockHeight, keccak256("blockhash")), xmsgs);
    }

    /// @dev Create an test XSubmission
    function _xsub(XTypes.Msg[] memory xmsgs) internal pure returns (XTypes.Submission memory) {
        return XTypes.Submission({
            attestationRoot: bytes32(0), // TODO: still unchecked
            blockHeader: XTypes.BlockHeader(0, 0, 0), // TODO: still unchecked
            msgs: xmsgs,
            proof: new bytes32[](0), // TODO: still unchecked
            proofFlags: new bool[](0), // TODO: still unchecked
            signatures: new XTypes.SigTuple[](0) // TODO: still unchecked
         });
    }

    /// @dev Create a Counter.increment() XMsg from thisChainId to chainAId
    function _outbound_increment() internal view returns (XTypes.Msg memory) {
        return _increment(thisChainId, chainAId, 0);
    }

    /// @dev Create a Counter.increment() XMsg from chainAId to thisChainId
    function _inbound_increment(uint64 streamOffset) internal view returns (XTypes.Msg memory) {
        return _increment(chainAId, thisChainId, streamOffset);
    }

    /// @dev Create a Reverter.forceRevert() XMsg from chainAId to thisChainId
    function _inbound_revert(uint64 streamOffset) internal view returns (XTypes.Msg memory) {
        return _revert(chainAId, thisChainId, streamOffset);
    }

    /// @dev Create a Counter.increment() XMsg
    function _increment(uint64 sourceChainId, uint64 destChainId, uint64 streamOffset)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            sourceChainId: sourceChainId,
            destChainId: destChainId,
            streamOffset: streamOffset,
            sender: _counters[sourceChainId],
            to: _counters[destChainId],
            data: abi.encodeWithSignature("increment()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    /// @dev Create a Reverter.forceRevert() XMsg
    function _revert(uint64 sourceChainId, uint64 destChainId, uint64 offset)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            sourceChainId: sourceChainId,
            destChainId: destChainId,
            streamOffset: offset,
            sender: _reverters[sourceChainId],
            to: _reverters[destChainId],
            data: abi.encodeWithSignature("forceRevert()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }
}
