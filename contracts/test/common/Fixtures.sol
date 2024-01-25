// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Vm } from "forge-std/Vm.sol";
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
contract Fixtures {
    Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    uint64 constant thisChainId = 1;
    uint64 constant chainAId = 2;
    uint64 constant chainBId = 3;

    address deployer;
    address xcaller;
    address xrelayer;

    TestPortal portal;
    TestPortal chainAPortal;
    TestPortal chainBPortal;

    Counter counter;
    Counter chainACounter;
    Counter chainBCounter;

    Reverter reverter;
    Reverter chainAReverter;
    Reverter chainBReverter;

    mapping(uint64 => address) public reverters;
    mapping(uint64 => address) public counters;

    function setUp() public {
        deployer = _makeAddr("deployer");
        xcaller = _makeAddr("xcaller");
        xrelayer = _makeAddr("relayer");

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

        counters[thisChainId] = address(counter);
        counters[chainAId] = address(chainACounter);
        counters[chainBId] = address(chainBCounter);

        reverters[thisChainId] = address(reverter);
        reverters[chainAId] = address(chainAReverter);
        reverters[chainBId] = address(chainBReverter);
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
            sender: counters[sourceChainId],
            to: counters[destChainId],
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
            sender: reverters[sourceChainId],
            to: reverters[destChainId],
            data: abi.encodeWithSignature("forceRevert()"),
            gasLimit: portal.XMSG_DEFAULT_GAS_LIMIT()
        });
    }

    // copied from forge-std/src/StdCheats.sol
    // creates a labeled address and the corresponding private key
    function _makeAddrAndKey(string memory name) internal virtual returns (address addr, uint256 privateKey) {
        privateKey = uint256(keccak256(abi.encodePacked(name)));
        addr = vm.addr(privateKey);
        vm.label(addr, name);
    }

    // copied from forge-std/src/StdCheats.sol
    // creates a labeled address
    function _makeAddr(string memory name) internal virtual returns (address addr) {
        (addr,) = _makeAddrAndKey(name);
    }
}
