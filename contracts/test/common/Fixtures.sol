// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { CommonBase } from "forge-std/Base.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { FeeOracleV1 } from "src/FeeOracleV1.sol";
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

    uint256 constant baseFee = 1 gwei;

    address deployer;
    address xcaller;
    address relayer;
    address admin;

    FeeOracleV1 feeOracle;
    FeeOracleV1 chainAFeeOracle;
    FeeOracleV1 chainBFeeOracle;

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
    mapping(uint64 => address) _reverters;
    mapping(uint64 => address) _counters;

    // @dev Path to which test XBlocks are written relative to project root. Read by ts utilites
    //      to generate XSubmissions for each test XBlock (see ts/script/genxsubs/io.ts)
    string constant XBLOCKS_PATH = "test/data/xblocks.json";

    // @dev Path to which test XSubmissions are written relative to project root. XSubmissions
    //      are generated for each test XBlock, per destination chain, by ts/script/genxsubs/main.ts.
    string constant XSUBS_PATH = "test/data/xsubs.json";

    /// @dev XSubs json read from XSUBS_PATH, stored to avoid re-reading from disk
    string private _xsubsJson;

    function setUp() public {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        relayer = makeAddr("relayer");
        admin = makeAddr("admin");

        vm.startPrank(deployer);

        vm.chainId(thisChainId); // portal constructor uses block.chainid
        feeOracle = new FeeOracleV1(deployer, baseFee);
        portal = new TestPortal(admin, address(feeOracle));
        counter = new Counter(portal);
        reverter = new Reverter();

        vm.chainId(chainAId);
        chainAFeeOracle = new FeeOracleV1(deployer, baseFee);
        chainAPortal = new TestPortal(admin, address(chainAFeeOracle));
        chainACounter = new Counter(chainAPortal);
        chainAReverter = new Reverter();

        vm.chainId(chainBId);
        chainBFeeOracle = new FeeOracleV1(deployer, baseFee);
        chainBPortal = new TestPortal(admin, address(chainBFeeOracle));
        chainBCounter = new Counter(chainBPortal);
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

        TestXTypes.Block memory xblock1 = _xblock(1, 1); // sourceBlockHeight: 1, startOffset: 1
        TestXTypes.Block memory xblock2 = _xblock(2, 6); // sourceBlockHeight: 2, startOffset: 6

        // id identifies the json object we are writing to within vm state
        // see https://book.getfoundry.sh/cheatcodes/serialize-json
        string memory id = "ID";
        vm.serializeBytes(id, "xblock1", abi.encode(xblock1));
        string memory json = vm.serializeBytes(id, "xblock2", abi.encode(xblock2));

        vm.writeJson(json, fullpath);
    }

    /// @dev Read a test fixture XSubmission from XSUBS_PATH, for a given xblockName and destChainId.
    ///      XSubmissions are generated by ts/script/genxsubs/main.ts, and written to XSUBS_PATH.
    function readXSubmission(string memory xblockName, uint64 destChainId) public returns (XTypes.Submission memory) {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/", XSUBS_PATH);

        if (bytes(_xsubsJson).length == 0) _xsubsJson = vm.readFile(path);

        // matches xsub name in ts/script/genxsubs/main.ts
        string memory xsubName = string.concat(xblockName, "_xsub_destChainId", Strings.toString(destChainId));
        bytes memory parsed = vm.parseJsonBytes(_xsubsJson, string.concat(".", xsubName));

        XTypes.Submission memory xsub = abi.decode(parsed, (XTypes.Submission));

        return xsub;
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
        xmsgs[5] = _increment(chainAId, chainBId, startOffset);
        xmsgs[6] = _increment(chainAId, chainBId, startOffset + 1);
        xmsgs[7] = _increment(chainAId, chainBId, startOffset + 2);
        xmsgs[8] = _revert(chainAId, chainBId, startOffset + 3);
        xmsgs[9] = _increment(chainAId, chainBId, startOffset + 4);

        return TestXTypes.Block(XTypes.BlockHeader(chainAId, sourceBlockHeight, keccak256("blockhash")), xmsgs);
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
