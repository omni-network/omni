// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { IFeeOracleV1 } from "src/interfaces/IFeeOracleV1.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { TestXTypes } from "./TestXTypes.sol";
import { PortalHarness } from "./PortalHarness.sol";
import { Counter } from "./Counter.sol";
import { Reverter } from "./Reverter.sol";
import { GasGuzzler } from "./GasGuzzler.sol";
import { XSubmitter } from "./XSubmitter.sol";

import { CommonBase } from "forge-std/Base.sol";
import { StdCheats } from "forge-std/StdCheats.sol";

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

    uint64 constant thisChainId = 100;
    uint64 constant chainAId = 102;
    uint64 constant chainBId = 103;

    uint64 constant omniChainId = 166;
    uint64 constant omniCChainID = 1_000_166;
    uint64 constant broadcastChainId = 0; // PORTAL._BROADCAST_CHAIN_ID

    uint64 constant baseValPower = 100;
    uint64 constant genesisValSetId = 1;

    address feeOracleManager = makeAddr("feeOracleManager");
    uint256 constant feeOracleBaseGasLimit = 50_000;
    uint256 constant feeOracleProtocolFee = 1 gwei;

    address deployer;
    address xcaller;
    address relayer;
    address owner;

    uint64 xmsgMaxGasLimit = 5_000_000;
    uint64 xmsgMinGasLimit = 21_000;
    uint16 xmsgMaxDataSize = 20_000;
    uint16 xreceiptMaxErrorSize = 256;

    string constant valMnemonic = "test test test test test test test test test test test junk";

    address val1;
    address val2;
    address val3;
    address val4;
    address val5;

    mapping(uint64 => XTypes.Validator[]) validatorSet;
    mapping(address => uint256) valPrivKeys;

    uint256 val1PrivKey;
    uint256 val2PrivKey;
    uint256 val3PrivKey;
    uint256 val4PrivKey;
    uint256 val5PrivKey;

    ProxyAdmin proxyAdmin;
    ProxyAdmin chainAProxyAdmin;
    ProxyAdmin chainBProxyAdmin;

    FeeOracleV1 feeOracle;
    FeeOracleV1 feeOracleImpl;
    FeeOracleV1 chainAFeeOracle;
    FeeOracleV1 chainAfeeOracleImpl;
    FeeOracleV1 chainBFeeOracle;
    FeeOracleV1 chainBfeeOracleImpl;

    PortalHarness portal;
    PortalHarness portalImpl;
    PortalHarness chainAPortal;
    PortalHarness chainAPortalImpl;
    PortalHarness chainBPortal;
    PortalHarness chainBPortalImpl;

    Counter counter;
    Counter chainACounter;
    Counter chainBCounter;

    Reverter reverter;
    Reverter chainAReverter;
    Reverter chainBReverter;

    XSubmitter xsubmitter;
    GasGuzzler gasGuzzler;

    // helper mappings to generate XMsg.to and XMsg.sender for some sourceChainId & destChainId
    mapping(uint64 => address) _reverters;
    mapping(uint64 => address) _counters;

    // @dev Path to which test XBlocks are written relative to project root. Read by ts utilites
    //      to generate XSubmissions for each test XBlock (see ts/script/genxsubs/io.ts)
    string constant XBLOCKS_PATH = "test/xchain/data/xblocks.json";

    // @dev Path to which test XSubmissions are written relative to project root. XSubmissions
    //      are generated for each test XBlock, per destination chain, by ts/script/genxsubs/main.ts.
    string constant XSUBS_PATH = "test/xchain/data/xsubs.json";

    /// @dev XSubs json read from XSUBS_PATH, stored to avoid re-reading from disk
    string private _xsubsJson;

    function setUp() public {
        _initAddrs();
        _initValidators();
        _initContracts();
    }

    /// @dev Write test fixture XBlocks to a json file. This is allows ts utilities
    ///      to generate XSubmissions, with valid merkle roots and proofs, for each test
    ///      XBlock. These XSubmissions are then used as inputs into test cases.
    function writeXBlocks() public {
        string memory root = vm.projectRoot();
        string memory fullpath = string.concat(root, "/", XBLOCKS_PATH);

        TestXTypes.Block memory xblock1 = _xblock({ offset: 1, xmsgOffset: 1 });
        TestXTypes.Block memory xblock2 = _xblock({ offset: 2, xmsgOffset: 6 });

        TestXTypes.Block memory guzzle1 = _guzzle_xblock({ numGuzzles: 1 });
        TestXTypes.Block memory guzzle5 = _guzzle_xblock({ numGuzzles: 5 });
        TestXTypes.Block memory guzzle10 = _guzzle_xblock({ numGuzzles: 10 });
        TestXTypes.Block memory guzzle25 = _guzzle_xblock({ numGuzzles: 25 });
        TestXTypes.Block memory guzzle50 = _guzzle_xblock({ numGuzzles: 50 });
        TestXTypes.Block memory reentrancy = _reentrancy_xblock();

        TestXTypes.Block memory addValSet2 = _addValidatorSet_xblock({ valSetId: 2 });

        // id identifies the json object we are writing to within vm state
        // see https://book.getfoundry.sh/cheatcodes/serialize-json
        string memory id = "ID";

        vm.serializeBytes(id, "xblock1", abi.encode(xblock1));
        vm.serializeBytes(id, "xblock2", abi.encode(xblock2));

        vm.serializeBytes(id, "guzzle1", abi.encode(guzzle1));
        vm.serializeBytes(id, "guzzle5", abi.encode(guzzle5));
        vm.serializeBytes(id, "guzzle10", abi.encode(guzzle10));
        vm.serializeBytes(id, "guzzle25", abi.encode(guzzle25));
        vm.serializeBytes(id, "guzzle50", abi.encode(guzzle50));
        vm.serializeBytes(id, "reentrancy", abi.encode(reentrancy));

        string memory json = vm.serializeBytes(id, "addValSet2", abi.encode(addValSet2));

        vm.writeJson(json, fullpath);
    }

    /// @dev Read a test fixture XSubmission, signed with the genesis val set
    function readXSubmission(string memory name, uint64 destChainId) public returns (XTypes.Submission memory) {
        return readXSubmission(name, destChainId, genesisValSetId);
    }

    /// @dev Read a test fixture XSubmission from XSUBS_PATH, for a given xblockName and destChainId.
    ///      XSubmissions are generated by ts/script/genxsubs/main.ts, and written to XSUBS_PATH.
    ///      Sign the submission with the validator set at valSetId.
    function readXSubmission(string memory name, uint64 destChainId, uint64 valSetId)
        public
        returns (XTypes.Submission memory)
    {
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/", XSUBS_PATH);

        if (bytes(_xsubsJson).length == 0) _xsubsJson = vm.readFile(path);

        // matches xsub name in ts/script/genxsubs/main.ts
        string memory xsubName = string.concat(name, "_xsub_destChainId", Strings.toString(destChainId));
        bytes memory parsed = vm.parseJsonBytes(_xsubsJson, string.concat(".", xsubName));

        XTypes.Submission memory xsub = abi.decode(parsed, (XTypes.Submission));

        xsub.signatures = getSignatures(valSetId, xsub.attestationRoot);
        xsub.validatorSetId = valSetId;

        return xsub;
    }

    /// @dev Generate a SigTuple array for a given valSetId and digest
    function getSignatures(uint64 valSetId, bytes32 digest) internal view returns (XTypes.SigTuple[] memory sigs) {
        XTypes.Validator[] storage vals = validatorSet[valSetId];
        sigs = new XTypes.SigTuple[](vals.length);

        for (uint256 i = 0; i < vals.length; i++) {
            sigs[i] = XTypes.SigTuple(vals[i].addr, _sign(digest, valPrivKeys[vals[i].addr]));
        }

        _sortSigsByAddr(sigs);
    }

    /// @dev Sign a digest with a private key
    function _sign(bytes32 digest, uint256 pk) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    /// @dev Sort sigs by validator address. OmniPortal.xsubmit expects sigs to be sorted.
    ///      Func is not really 'pure', because it modifies the input array in place. But it does not depend on contract state.
    function _sortSigsByAddr(XTypes.SigTuple[] memory sigs) internal pure {
        uint256 n = sigs.length;
        for (uint256 i = 0; i < n - 1; i++) {
            for (uint256 j = 0; j < n - i - 1; j++) {
                if (sigs[j].validatorAddr > sigs[j + 1].validatorAddr) {
                    (sigs[j], sigs[j + 1]) = (sigs[j + 1], sigs[j]);
                }
            }
        }
    }

    /// @dev Create an xblock from chainA with xmsgs for "this" chain and chain b.
    ///      XBlocks will likely contain XMsgs for multiple chains, so we reflect that here.
    function _xblock(uint64 offset, uint64 xmsgOffset) internal view returns (TestXTypes.Block memory) {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](10);

        // intended for this chain
        xmsgs[0] = _increment(chainAId, thisChainId, xmsgOffset);
        xmsgs[1] = _increment(chainAId, thisChainId, xmsgOffset + 1);
        xmsgs[2] = _increment(chainAId, thisChainId, xmsgOffset + 2);
        xmsgs[3] = _revert(chainAId, thisChainId, xmsgOffset + 3);
        xmsgs[4] = _increment(chainAId, thisChainId, xmsgOffset + 4);

        // intended for chain b
        xmsgs[5] = _increment(chainAId, chainBId, xmsgOffset);
        xmsgs[6] = _increment(chainAId, chainBId, xmsgOffset + 1);
        xmsgs[7] = _increment(chainAId, chainBId, xmsgOffset + 2);
        xmsgs[8] = _revert(chainAId, chainBId, xmsgOffset + 3);
        xmsgs[9] = _increment(chainAId, chainBId, xmsgOffset + 4);

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: chainAId,
                consensusChainId: omniCChainID,
                confLevel: ConfLevel.Finalized,
                offset: offset,
                sourceBlockHeight: 100,
                sourceBlockHash: keccak256("blockhash")
            }),
            xmsgs
        );
    }

    function _addValidatorSet_xblock(uint64 valSetId) internal view returns (TestXTypes.Block memory) {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);
        xmsgs[0] = XTypes.Msg({
            destChainId: broadcastChainId,
            shardId: ConfLevel.toBroadcastShard(ConfLevel.Finalized),
            offset: valSetId,
            sender: address(0), // Portal._CCHAIN_SENDER
            to: address(0), // Portal._VIRTUAL_PORTAL_ADDRRESS
            data: abi.encodeWithSelector(OmniPortal.addValidatorSet.selector, valSetId, validatorSet[valSetId]),
            gasLimit: 0
        });

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: omniCChainID,
                consensusChainId: omniCChainID,
                confLevel: ConfLevel.Finalized,
                offset: valSetId,
                sourceBlockHeight: 100,
                sourceBlockHash: bytes32(0)
            }),
            xmsgs
        );
    }

    function _guzzle_xblock(uint256 numGuzzles) internal view returns (TestXTypes.Block memory) {
        uint64 offset = 1;
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](numGuzzles);

        for (uint256 i = 0; i < numGuzzles; i++) {
            // all guzzles from chain a to this chain
            xmsgs[i] = _guzzle(thisChainId, offset, 100_000);
            offset += 1;
        }

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: chainAId,
                consensusChainId: omniCChainID,
                offset: 1,
                confLevel: ConfLevel.Finalized,
                sourceBlockHeight: 100,
                sourceBlockHash: keccak256("blockhash")
            }),
            xmsgs
        );
    }

    function _reentrancy_xblock() internal view returns (TestXTypes.Block memory) {
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);
        xmsgs[0] = XTypes.Msg({
            destChainId: thisChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: 1,
            sender: address(xsubmitter),
            to: address(xsubmitter),
            data: abi.encodeWithSelector(XSubmitter.tryXSubmit.selector),
            gasLimit: 100_000
        });

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: chainAId,
                consensusChainId: omniCChainID,
                confLevel: ConfLevel.Finalized,
                offset: 1,
                sourceBlockHeight: 100,
                sourceBlockHash: keccak256("blockhash")
            }),
            xmsgs
        );
    }

    /// @dev Create a Counter.increment() XMsg from thisChainId to chainAId
    function _outbound_increment() internal view returns (XTypes.Msg memory) {
        return _increment(thisChainId, chainAId, 0);
    }

    /// @dev Create a Counter.increment() XMsg from chainAId to thisChainId
    function _inbound_increment(uint64 offset) internal view returns (XTypes.Msg memory) {
        return _increment(chainAId, thisChainId, offset);
    }

    /// @dev Create a Reverter.forceRevert() XMsg from chainAId to thisChainId
    function _inbound_revert(uint64 offset) internal view returns (XTypes.Msg memory) {
        return _revert(chainAId, thisChainId, offset);
    }

    /// @dev Create a Counter.increment() XMsg
    function _increment(uint64 sourceChainId, uint64 destChainId, uint64 offset)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            destChainId: destChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: offset,
            sender: _counters[sourceChainId],
            to: _counters[destChainId],
            data: abi.encodeWithSignature("increment()"),
            gasLimit: 100_000
        });
    }

    /// @dev Create a Reverter.forceRevert() XMsg
    function _revert(uint64 sourceChainId, uint64 destChainId, uint64 offset)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return _reverter_xmsg(sourceChainId, destChainId, offset, abi.encodeWithSignature("forceRevert()"));
    }

    /// @dev Helper to create an xmsg to the Reverter contract
    function _reverter_xmsg(uint64 sourceChainId, uint64 destChainId, uint64 offset, bytes memory data)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            destChainId: destChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: offset,
            sender: _reverters[sourceChainId],
            to: _reverters[destChainId],
            data: data,
            gasLimit: 200_000
        });
    }

    /// @dev Create a GasGuzzler.guzzle() XMsg
    function _guzzle(uint64 destChainId, uint64 offset, uint64 gasLimit) internal view returns (XTypes.Msg memory) {
        return XTypes.Msg({
            destChainId: destChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: offset,
            sender: address(gasGuzzler),
            to: address(gasGuzzler),
            data: abi.encodeWithSignature("guzzle()"),
            gasLimit: gasLimit
        });
    }

    /// @dev Initialize test addresses
    function _initAddrs() private {
        deployer = makeAddr("deployer");
        xcaller = makeAddr("xcaller");
        relayer = makeAddr("relayer");
        owner = makeAddr("owner");
        vm.deal(xcaller, 100 ether); // fund xcaller, so it can pay fees
    }

    /// @dev Initialize test validators
    function _initValidators() private {
        (val1, val1PrivKey) = deriveRememberKey(valMnemonic, 0);
        (val2, val2PrivKey) = deriveRememberKey(valMnemonic, 1);
        (val3, val3PrivKey) = deriveRememberKey(valMnemonic, 2);
        (val4, val4PrivKey) = deriveRememberKey(valMnemonic, 3);
        (val5, val5PrivKey) = deriveRememberKey(valMnemonic, 4);

        valPrivKeys[val1] = val1PrivKey;
        valPrivKeys[val2] = val2PrivKey;
        valPrivKeys[val3] = val3PrivKey;
        valPrivKeys[val4] = val4PrivKey;
        valPrivKeys[val5] = val5PrivKey;

        // only use 1-4 for val set 1
        XTypes.Validator[] storage genVals = validatorSet[genesisValSetId];
        genVals.push(XTypes.Validator(val1, baseValPower));
        genVals.push(XTypes.Validator(val2, baseValPower));
        genVals.push(XTypes.Validator(val3, baseValPower));
        genVals.push(XTypes.Validator(val4, baseValPower));

        // val set 2 adds one validator, and removes val2
        XTypes.Validator[] storage valSet2 = validatorSet[genesisValSetId + 1];
        valSet2.push(XTypes.Validator(val1, baseValPower));
        valSet2.push(XTypes.Validator(val3, baseValPower));
        valSet2.push(XTypes.Validator(val4, baseValPower));
        valSet2.push(XTypes.Validator(val5, baseValPower));
    }

    /// @dev Initialize portals and test contracts
    function _initContracts() private {
        vm.startPrank(deployer);

        IFeeOracleV1.ChainFeeParams[] memory feeParams = new IFeeOracleV1.ChainFeeParams[](3);

        feeParams[0] = IFeeOracleV1.ChainFeeParams({
            chainId: thisChainId,
            postsTo: thisChainId,
            gasPrice: 0.1 gwei, // 1 gwei
            toNativeRate: 1e6 // feeOracle.CONVERSION_RATE_DENOM , so 1:1
         });

        feeParams[1] = IFeeOracleV1.ChainFeeParams({
            chainId: chainAId,
            postsTo: chainAId,
            gasPrice: 0.1 gwei, // 1 gwei
            toNativeRate: 1e6 // feeOracle.CONVERSION_RATE_DENOM , so 1:1
         });

        feeParams[2] = IFeeOracleV1.ChainFeeParams({
            chainId: chainBId,
            postsTo: chainAId, // let's have chainB "rollup" to chainA
            gasPrice: 0.1 gwei, // 1 gwei
            toNativeRate: 1e6 // feeOracle.CONVERSION_RATE_DENOM , so 1:1
         });

        proxyAdmin = new ProxyAdmin(owner);

        feeOracleImpl = new FeeOracleV1();
        feeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(feeOracleImpl),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        FeeOracleV1.initialize.selector,
                        owner,
                        feeOracleManager,
                        feeOracleBaseGasLimit,
                        feeOracleProtocolFee,
                        feeParams
                    )
                )
            )
        );
        portalImpl = new PortalHarness();
        portal = PortalHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(portalImpl),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        OmniPortal.initialize.selector,
                        OmniPortal.InitParams(
                            owner,
                            address(feeOracle),
                            omniChainId,
                            omniCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            validatorSet[genesisValSetId]
                        )
                    )
                )
            )
        );
        counter = new Counter(portal);
        reverter = new Reverter();

        chainAProxyAdmin = new ProxyAdmin(owner);

        chainAfeeOracleImpl = new FeeOracleV1();
        chainAFeeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(chainAfeeOracleImpl),
                    address(chainAProxyAdmin),
                    abi.encodeWithSelector(
                        FeeOracleV1.initialize.selector,
                        owner,
                        feeOracleManager,
                        feeOracleBaseGasLimit,
                        feeOracleProtocolFee,
                        feeParams
                    )
                )
            )
        );
        chainAPortalImpl = new PortalHarness();
        chainAPortal = PortalHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(chainAPortalImpl),
                    address(chainAProxyAdmin),
                    abi.encodeWithSelector(
                        OmniPortal.initialize.selector,
                        OmniPortal.InitParams(
                            owner,
                            address(feeOracle),
                            omniChainId,
                            omniCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            validatorSet[genesisValSetId]
                        )
                    )
                )
            )
        );
        chainACounter = new Counter(chainAPortal);
        chainAReverter = new Reverter();

        chainBProxyAdmin = new ProxyAdmin(owner);

        chainBfeeOracleImpl = new FeeOracleV1();
        chainBFeeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(chainBfeeOracleImpl),
                    address(chainBProxyAdmin),
                    abi.encodeWithSelector(
                        FeeOracleV1.initialize.selector,
                        owner,
                        feeOracleManager,
                        feeOracleBaseGasLimit,
                        feeOracleProtocolFee,
                        feeParams
                    )
                )
            )
        );
        chainBPortalImpl = new PortalHarness();
        chainBPortal = PortalHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(chainBPortalImpl),
                    address(chainBProxyAdmin),
                    abi.encodeWithSelector(
                        OmniPortal.initialize.selector,
                        OmniPortal.InitParams(
                            owner,
                            address(feeOracle),
                            omniChainId,
                            omniCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            validatorSet[genesisValSetId]
                        )
                    )
                )
            )
        );
        chainBCounter = new Counter(chainBPortal);
        chainBReverter = new Reverter();

        gasGuzzler = new GasGuzzler();
        xsubmitter = new XSubmitter(portal);

        vm.stopPrank();

        _counters[thisChainId] = address(counter);
        _counters[chainAId] = address(chainACounter);
        _counters[chainBId] = address(chainBCounter);

        _reverters[thisChainId] = address(reverter);
        _reverters[chainAId] = address(chainAReverter);
        _reverters[chainBId] = address(chainBReverter);

        XTypes.Chain[] memory network = new XTypes.Chain[](3);
        network[0] = XTypes.Chain({ chainId: thisChainId, shards: new uint64[](2) });
        network[1] = XTypes.Chain({ chainId: chainAId, shards: new uint64[](2) });
        network[2] = XTypes.Chain({ chainId: chainBId, shards: new uint64[](2) });
        network[0].shards[0] = ConfLevel.Finalized;
        network[0].shards[1] = ConfLevel.Latest;
        network[1].shards[0] = ConfLevel.Finalized;
        network[1].shards[1] = ConfLevel.Latest;
        network[2].shards[0] = ConfLevel.Finalized;
        network[2].shards[1] = ConfLevel.Latest;

        vm.chainId(thisChainId);
        portal.setNetworkNoAuth(network);

        vm.chainId(chainAId);
        chainAPortal.setNetworkNoAuth(network);

        vm.chainId(chainBId);
        chainBPortal.setNetworkNoAuth(network);
    }
}
