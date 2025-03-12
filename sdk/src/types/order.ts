import type {
  Abi,
  Address,
  ContractFunctionArgs,
  ContractFunctionName,
  ContractFunctionParameters,
} from 'viem'
import type { AbiWriteMutability, OptionalAbi, OptionalAbis } from './abi.js'

export type OrderStatus =
  | 'not-found'
  | 'open'
  | 'rejected'
  | 'closed'
  | 'filled'
  | 'error'

type Deposit = {
  readonly token?: Address
  readonly amount: bigint
}

type Expense = {
  readonly spender?: Address
  readonly token?: Address
  readonly amount: bigint
}

export type ContractCall<
  abi extends Abi,
  mutability extends AbiWriteMutability = AbiWriteMutability,
  functionName extends ContractFunctionName<
    abi,
    mutability
  > = ContractFunctionName<abi, mutability>,
  args extends ContractFunctionArgs<
    abi,
    mutability,
    functionName
  > = ContractFunctionArgs<abi, mutability, functionName>,
> = Omit<
  ContractFunctionParameters<abi, mutability, functionName, args>,
  'address' // we use `target` instead of `address
> &
  (mutability extends 'payable'
    ? {
        readonly value?: bigint
        readonly target: Address
      }
    : {
        readonly target: Address
      })

type NativeTransfer = {
  readonly abi?: Abi // allows auto-complete abi, type will narrow if abi is provided
  readonly target: Address
  readonly value: bigint
}

export type Call<abi extends OptionalAbi> = abi extends Abi
  ? ContractCall<abi>
  : NativeTransfer

export type Calls<abis extends OptionalAbis> = {
  [index in keyof abis]: abis[index] extends Abi
    ? ContractCall<abis[index]>
    : NativeTransfer
}

export type Order<abis extends OptionalAbis> = {
  readonly owner?: Address
  readonly srcChainId?: number
  readonly destChainId: number
  readonly fillDeadline?: number
  readonly calls: Calls<abis>
  readonly deposit: Deposit
  readonly expense: Expense
}

// isContractCall call narrows a call type based on the presence of abi
// this is okay, because when abi is defined, call type is narrowed to ContractCall
export function isContractCall<
  abi extends OptionalAbi,
  defined extends NonNullable<abi> = NonNullable<abi>,
>(call: NativeTransfer | ContractCall<defined>): call is ContractCall<defined> {
  return call.abi != null
}
