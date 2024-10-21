// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Preinstalls } from "src/octane/Preinstalls.sol";
import { AllocPredeploys } from "script/genesis/AllocPredeploys.s.sol";
import { Test, Vm, console2 } from "forge-std/Test.sol";

interface IEIP712 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

interface IMultiCall3 {
    struct Call {
        address target;
        bytes callData;
    }

    struct Result {
        bool success;
        bytes returnData;
    }

    function aggregate(Call[] calldata calls)
        external
        payable
        returns (uint256 blockNumber, bytes[] memory returnData);
    function tryAggregate(bool requireSuccess, Call[] calldata calls)
        external
        payable
        returns (Result[] memory returnData);

    function getBlockNumber() external view returns (uint256 blockNumber);
    function getLastBlockHash() external view returns (bytes32 blockHash);
    function getChainId() external view returns (uint256 chainid);
}

interface ICreate2Deployer {
    function computeAddress(bytes32 salt, bytes32 codeHash) external view returns (address);
}

interface ISafe_v130 {
    function getChainId() external view returns (uint256);
}

interface ISafeL2_v130 {
    function domainSeparator() external view returns (bytes32);
}

interface IMultiSendCallOnly_v130 {
    function multiSend(bytes memory transactions) external payable;
}

interface IMultiSend_v130 {
    function multiSend(bytes memory transactions) external payable;
}

interface IPermit2 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

/**
 * @notice Used in commented out test_safeSingletonFactory_fallback_succeeds test
 */
contract DummyTest {
    function chainId() external view returns (uint256) {
        return block.chainid;
    }
}

/**
 * @title Preinstalls_Test
 * @notice Test suite for Preinstalls.sol. Most of Preinstalls is just static bytecode.
 *         We only tests Permit2 templating.
 */
contract Preinstalls_Test is Test, AllocPredeploys {
    /**
     * @notice Supplied as native bridge balance
     */
    uint256 omniTotalSupply = 100e6 * 1e18;

    /**
     * @notice Preinstall configuration borrowed from AllocPredeploys test
     */
    function setUp() public virtual {
        this.runNoStateDump(
            AllocPredeploys.Config({
                manager: makeAddr("manager"),
                upgrader: makeAddr("upgrader"),
                chainId: 165,
                nativeBridgeBalance: omniTotalSupply,
                enableStakingAllowlist: false,
                output: ""
            })
        );
    }

    /**
     * @notice Test getPermit2Code templating. This templating inserts immutable variables into the bytecode.
     */
    function test_getPermit2Code() public {
        bytes32 typeHash =
            keccak256(abi.encodePacked("EIP712Domain(string name,uint256 chainId,address verifyingContract)"));
        bytes32 nameHash = keccak256(abi.encodePacked("Permit2"));
        uint256 chainId = 165;
        bytes32 domainSeparator = keccak256(abi.encode(typeHash, nameHash, chainId, Preinstalls.Permit2));

        vm.etch(Preinstalls.Permit2, Preinstalls.getPermit2Code(chainId));

        vm.chainId(chainId);
        assertEq(IEIP712(Preinstalls.Permit2).DOMAIN_SEPARATOR(), domainSeparator);
    }

    function test_multiCall3_aggregate_succeeds() public {
        IMultiCall3.Call[] memory calls = new IMultiCall3.Call[](3);
        calls[0] = IMultiCall3.Call({
            target: Preinstalls.MultiCall3,
            callData: abi.encodeWithSelector(IMultiCall3.getBlockNumber.selector)
        });
        calls[1] = IMultiCall3.Call({
            target: Preinstalls.MultiCall3,
            callData: abi.encodeWithSelector(IMultiCall3.getLastBlockHash.selector)
        });
        calls[2] = IMultiCall3.Call({
            target: Preinstalls.MultiCall3,
            callData: abi.encodeWithSelector(IMultiCall3.getChainId.selector)
        });
        (, bytes[] memory returnData) = IMultiCall3(Preinstalls.MultiCall3).aggregate(calls);
        assertEq(uint256(bytes32(returnData[0])), block.number, "MultiCall3 getBlockNumber result mismatch");
        assertEq(bytes32(returnData[1]), blockhash(block.number - 1), "MultiCall3 getLastBlockHash result mismatch");
        assertEq(uint256(bytes32(returnData[2])), block.chainid, "MultiCall3 getChainId result mismatch");
    }

    function test_multiCall3_tryAggregate_succeeds() public {
        IMultiCall3.Call[] memory calls = new IMultiCall3.Call[](3);
        calls[0] = IMultiCall3.Call({
            target: Preinstalls.Create2Deployer,
            callData: abi.encodeWithSelector(
                ICreate2Deployer.computeAddress.selector,
                keccak256(abi.encode("SALT")),
                keccak256(abi.encodePacked(Preinstalls.Create2DeployerCode))
            )
        });
        calls[1] = IMultiCall3.Call({
            target: Preinstalls.Safe_v130,
            callData: abi.encodeWithSelector(ISafe_v130.getChainId.selector)
        });
        calls[2] = IMultiCall3.Call({
            target: Preinstalls.SafeL2_v130,
            callData: abi.encodeWithSelector(ISafeL2_v130.domainSeparator.selector)
        });

        IMultiCall3.Result[] memory returnData = IMultiCall3(Preinstalls.MultiCall3).tryAggregate(true, calls);
        assertEq(
            address(uint160(uint256(bytes32(returnData[0].returnData)))),
            0x87C9fcd2CcAc5fDe445970AbC4fF1e6e74f325d7,
            "Create2Deployer computeAddress result mismatch"
        ); // Precomputed address
        assertEq(uint256(bytes32(returnData[1].returnData)), 165, "Safe_v130 getChainId result mismatch");
        assertEq(
            bytes32(returnData[2].returnData),
            keccak256(
                abi.encode(
                    bytes32(hex"47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a79469218"),
                    uint256(165),
                    0xfb1bffC9d739B8D520DaF37dF666da4C687191EA
                )
            ),
            "SafeL2_v130 domainSeparator result mismatch"
        );
    }

    function test_multiSendCallOnly_v130_multiSend_succeeds() public {
        bytes memory callData = abi.encodeWithSelector(ISafe_v130.getChainId.selector);
        bytes memory transaction =
            abi.encodePacked(uint8(0), Preinstalls.Safe_v130, uint256(0), uint256(callData.length), callData);

        IMultiSendCallOnly_v130(Preinstalls.MultiSendCallOnly_v130).multiSend(transaction);
    }

    /**
     * @notice Contract address is unverified everywhere, assuming it is this: https://github.com/safe-global/safe-singleton-factory/blob/main/contracts-zk/SafeSingeltonFactory.sol
     * @notice Test keeps failing due to a StackUnderflow error. It seems the deployment initiates, but the contract deployment does not complete. DummyTest is very simple.
     * @dev Here's the stack trace:
     *    [8937393460516729750] Preinstalls_Test::test_safeSingletonFactory_fallback_succeeds()
     *      ├─ [8937393460516725182] 0x914d7Fec6aaC8cd542e72Bca78B30650d45643d7::53414c54(00000000000000000000000000000000000000000000000000000000f293e1c100000000000000000000000000000000000000000000000000000000000000006080604052348015600f57600080fd5b50607680601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80639a8a059214602d575b600080fd5b4660405190815260200160405180910390f3fea26469706673582212203929177a6e5a24875e37fc9f5c0e124385964bf07ce9f637239545d60dde3f2964736f6c63430008180033)
     *      │   ├─ [0] → new <unknown>@ 0x4Db51C70c85DB889BB746568b636A51F33Fad21c
     *      │   │   └─ ← [StackUnderflow] EvmError: StackUnderflow
     *      │   └─ ← [Revert] EvmError: Revert
     *      └─ ← [Revert] revert: Deployment using SafeSingletonFactory failed
     * @dev DeterministicDeploymentProxy has the same bytecode as SafeSingletonFactory, so it will likely encounter the same problem
     */
    function test_safeSingletonFactory_fallback_succeeds() public {
        bytes memory callData = bytes.concat(
            hex"0000000000000000000000000000000000000000000000000000000000000000", type(DummyTest).creationCode
        );
        (bool success,) = payable(Preinstalls.SafeSingletonFactory).call{ value: 0 }(callData);
        require(success, "Deployment using SafeSingletonFactory failed");

        address deployment = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            bytes1(hex"ff"),
                            Preinstalls.SafeSingletonFactory,
                            bytes32(0),
                            keccak256(type(DummyTest).creationCode)
                        )
                    )
                )
            )
        );
        assertTrue(deployment.code.length != 0, "Contract deployed with SafeSingletonFactory unusable");
    }

    function test_multiSend_v130_multiSend_succeeds() public {
        IPermit2(Preinstalls.Permit2).DOMAIN_SEPARATOR();
        bytes memory callData = abi.encodeWithSelector(IPermit2.DOMAIN_SEPARATOR.selector);
        bytes memory transaction =
            abi.encodePacked(uint8(0), Preinstalls.Permit2, uint256(0), uint256(callData.length), callData);

        (bool success, bytes memory data) =
            Preinstalls.MultiSend_v130.delegatecall(abi.encodeWithSignature("multiSend(bytes memory)", transaction));
        console2.log(success);
        console2.logBytes(data);
    }
}
