// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SolverNet } from "../lib/SolverNet.sol";

/**
 * @title ISolverNetBindings
 * @notice This interface is only used to generate bindings for ISolverNet types.
 *         This is necessary because inbox / outbox abis operate on encoded ISolverNet bytes,
 *         so types are not represented in their ABIs.
 * @dev We use a single function per type so that we can use go-ethereum's
 *      abi[method].Inputs to encode / decode one type at a time.
 */
interface ISolverNetBindings {
    function orderData(SolverNet.OrderData calldata) external view;
    function order(SolverNet.Order calldata) external view;
    function header(SolverNet.Header calldata) external view;
    function deposit(SolverNet.Deposit calldata) external view;
    function call(SolverNet.Call calldata) external view;
    function expense(SolverNet.Expense calldata) external view;
    function fillOriginData(SolverNet.FillOriginData calldata) external view;
}
