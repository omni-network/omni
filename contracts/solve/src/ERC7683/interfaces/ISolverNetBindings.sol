// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ISolverNet } from "./ISolverNet.sol";

/**
 * @title ISolverNetBindings
 * @notice This interface is only used to generate bindings for ISolverNet types.
 *         This is necessary because inbox / oubox abis operate on encoded ISolverNet bytes,
 *         so types are not represented in their ABIs.
 * @dev We use a single function per type so that we can use go-ethereum's
 *      abi[method].Inputs to encode / decode one type at a time.
 */
interface ISolverNetBindings is ISolverNet {
    function deposit(Deposit calldata) external view;
    function tokenExpense(TokenExpense calldata) external view;
    function call(Call calldata) external view;
    function orderData(OrderData calldata) external view;
    function fillOriginData(FillOriginData calldata) external view;
}
