// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";
// solhint-disable-next-line no-unused-import
import { IERC7683 } from "src/erc7683/IERC7683.sol";
// solhint-disable-next-line no-unused-import
import { SolverNet } from "src/lib/SolverNet.sol";

contract SolverNetOmegaFixtures is Script {
    IERC20 internal constant nomina = IERC20(0xb7dA35e8f69aA04422cBEF30850Da77614343087);
    IOmniPortal internal constant portal = IOmniPortal(0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3);
    ISolverNetInbox internal constant inbox = ISolverNetInbox(0x7EA0CeB70D5Df75a730E6cB3EADeC12EfdFe80a1);

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );
}
