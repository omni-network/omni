import { keccak256, toBytes } from 'viem'

const orderDataType =
  'OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)'

export const typeHash = keccak256(toBytes(orderDataType))
