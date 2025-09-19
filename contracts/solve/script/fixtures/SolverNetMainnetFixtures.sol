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

contract SolverNetMainnetFixtures is Script {
    IERC20 internal constant omni = IERC20(0x36E66fbBce51e4cD5bd3C62B637Eb411b18949D4);
    IERC20 internal constant nomina = IERC20(0x6e6F6d696e61decd6605bD4a57836c5DB6923340);
    IOmniPortal internal constant portal = IOmniPortal(0x5e9A8Aa213C912Bf54C86bf64aDB8ed6A79C04d1);
    ISolverNetInbox internal constant inbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );
}
