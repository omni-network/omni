// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { MockToken } from "test/utils/MockToken.sol";
import { SolveInbox } from "src/solve/SolveInbox.sol";
import { Solve } from "src/solve/Solve.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title InboxBase
 * @dev Shared test solve utils / fixtures.
 */
contract InboxBase is Test {
    SolveInbox inbox;

    address user = makeAddr("user");
    address solver = makeAddr("solver");

    MockToken token1;
    MockToken token2;

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public {
        inbox = deploySolveInbox();
        token1 = new MockToken();
        token2 = new MockToken();
    }

    function randCall() internal returns (Solve.Call memory) {
        uint256 rand = vm.randomUint(1, 1000);
        return Solve.Call({
            destChainId: uint64(rand),
            value: rand * 1 ether,
            target: address(uint160(rand)),
            data: abi.encode("data", rand)
        });
    }

    function mintAndApprove(Solve.TokenDeposit[] memory deposits) internal {
        for (uint256 i = 0; i < deposits.length; i++) {
            MockToken(deposits[i].token).approve(address(inbox), deposits[i].amount);
            MockToken(deposits[i].token).mint(user, deposits[i].amount);
        }
    }

    function deploySolveInbox() internal returns (SolveInbox) {
        address impl = address(new SolveInbox());
        return SolveInbox(
            address(
                new TransparentUpgradeableProxy(
                    impl,
                    makeAddr("proxy-admin-owner"),
                    abi.encodeCall(
                        SolveInbox.initialize, (address(this), solver, makeAddr("portal"), makeAddr("outbox"))
                    )
                )
            )
        );
    }
}
