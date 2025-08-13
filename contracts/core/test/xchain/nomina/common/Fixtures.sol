// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { IFeeOracleV1 } from "src/interfaces/IFeeOracleV1.sol";
import { NominaPortal } from "src/xchain/nomina/NominaPortal.sol";
import { TestXTypes } from "test/xchain/common/TestXTypes.sol";
import { PortalHarness } from "./PortalHarness.sol";
import { Counter } from "./Counter.sol";
import { Reverter } from "test/xchain/common/Reverter.sol";
import { GasGuzzler } from "test/xchain/common/GasGuzzler.sol";
import { XSubmitter } from "./XSubmitter.sol";

import { CommonBase } from "forge-std/Base.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { XSubGen } from "test/utils/nomina/XSubGen.sol";

/**
 * @title Fixtures
 * @dev Defines test dependencies. Deploys portals and test contracts, and provides
 *      utilites for generating XMsgs and XSubmissions between them.
 */
contract Fixtures is CommonBase, StdCheats, XSubGen {
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

    uint64 constant nominaChainId = 166;
    uint64 constant nominaCChainID = 1_000_166;
    uint64 constant broadcastChainId = 0; // PORTAL._BROADCAST_CHAIN_ID

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

    FeeOracleV1 feeOracle;
    FeeOracleV1 feeOracleImpl;
    FeeOracleV1 chainAFeeOracle;
    FeeOracleV1 chainAfeeOracleImpl;
    FeeOracleV1 chainBFeeOracle;
    FeeOracleV1 chainBfeeOracleImpl;

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

    function setUp() public virtual {
        _initAddrs();
        _initContracts();
    }

    /// @dev Generate a SigTuple array for a given valSetId and digest
    function getSignatures(uint64 valSetId, bytes32 digest) internal view returns (XTypes.SigTuple[] memory sigs) {
        return _getSignatures(valSetId, digest);
    }

    /// @dev Generate a new valset of an arbitrary size
    function newValSet(uint32 valSetSize) public returns (uint64 valSetId) {
        require(valSetSize >= 3, "Validator set size is too small");

        valSetId = ++valsetCount;
        Validator[] storage newValset = valset[valSetId];

        for (uint32 i; i < valSetSize; ++i) {
            Validator memory validator;
            (address val, uint256 valPk) = deriveRememberKey(mnemonic, i);
            validator = Validator(val, baseValPower, valPk);
            newValset.push(validator);
        }

        return valSetId;
    }

    /// @dev Sort sigs by validator address. NominaPortal.xsubmit expects sigs to be sorted.
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

    /// @dev Manually create an xblock from sourceChainId with crafted xmsgs
    function _xblock(uint64 sourceChainId, uint8 confLevel, uint64 offset, XTypes.Msg[] memory xmsgs)
        internal
        pure
        returns (TestXTypes.Block memory)
    {
        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: sourceChainId,
                consensusChainId: nominaCChainID,
                confLevel: confLevel,
                offset: offset,
                sourceBlockHeight: 100,
                sourceBlockHash: keccak256("blockhash")
            }),
            xmsgs
        );
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
                consensusChainId: nominaCChainID,
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
            data: abi.encodeWithSelector(NominaPortal.addValidatorSet.selector, valSetId, getVals(valSetId)),
            gasLimit: 0
        });

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: nominaCChainID,
                consensusChainId: nominaCChainID,
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
                consensusChainId: nominaCChainID,
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
                consensusChainId: nominaCChainID,
                confLevel: ConfLevel.Finalized,
                offset: 1,
                sourceBlockHeight: 100,
                sourceBlockHash: keccak256("blockhash")
            }),
            xmsgs
        );
    }

    function _deadCall_xblock(uint256 deadCalls, bytes memory data) internal view returns (TestXTypes.Block memory) {
        uint64 offset = 1;
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](deadCalls);

        for (uint256 i = 0; i < deadCalls; i++) {
            // all dead calls from chain a to this chain
            xmsgs[i] = _deadCall(thisChainId, offset, data);
            offset += 1;
        }

        return TestXTypes.Block(
            XTypes.BlockHeader({
                sourceChainId: chainAId,
                consensusChainId: nominaCChainID,
                offset: 1,
                confLevel: ConfLevel.Finalized,
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

    function _deadCall(uint64 destChainId, uint64 offset, bytes memory data)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            destChainId: destChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: offset,
            sender: address(this),
            to: address(0xdead),
            data: data,
            gasLimit: uint64(21_000 + (16 * data.length))
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

        feeOracleImpl = new FeeOracleV1();
        feeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(feeOracleImpl),
                    owner,
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
                    owner,
                    abi.encodeWithSelector(
                        NominaPortal.initialize.selector,
                        NominaPortal.InitParams(
                            owner,
                            address(feeOracle),
                            nominaChainId,
                            nominaCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            getVals(genesisValSetId)
                        )
                    )
                )
            )
        );
        counter = new Counter(portal);
        reverter = new Reverter();

        chainAfeeOracleImpl = new FeeOracleV1();
        chainAFeeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(chainAfeeOracleImpl),
                    owner,
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
                    owner,
                    abi.encodeWithSelector(
                        NominaPortal.initialize.selector,
                        NominaPortal.InitParams(
                            owner,
                            address(feeOracle),
                            nominaChainId,
                            nominaCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            getVals(genesisValSetId)
                        )
                    )
                )
            )
        );
        chainACounter = new Counter(chainAPortal);
        chainAReverter = new Reverter();

        chainBfeeOracleImpl = new FeeOracleV1();
        chainBFeeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    address(chainBfeeOracleImpl),
                    owner,
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
                    owner,
                    abi.encodeWithSelector(
                        NominaPortal.initialize.selector,
                        NominaPortal.InitParams(
                            owner,
                            address(feeOracle),
                            nominaChainId,
                            nominaCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            10, // xsub valset cutoff
                            1, // cchain xmsg offset
                            1, // cchain xblock offset
                            genesisValSetId,
                            getVals(genesisValSetId)
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
