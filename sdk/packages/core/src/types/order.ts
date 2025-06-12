import type {
  Abi,
  ContractFunctionArgs,
  ContractFunctionName,
  ContractFunctionParameters,
  Hex,
} from 'viem'
import type { AbiWriteMutability, OptionalAbi, OptionalAbis } from './abi.js'
import type { EVMAddress, SVMAddress } from './addresses.js'

export type EVMOrderId = Hex
export type SVMOrderId = SVMAddress
export type OrderId = EVMOrderId | SVMOrderId

export type OrderStatus =
  | 'not-found'
  | 'open'
  | 'rejected'
  | 'closed'
  | 'filled'
  | 'error'

// EVM and SVM deposits will be supported
type EVMDeposit = {
  readonly token?: EVMAddress
  readonly amount: bigint
}

// only EVM expenses are supported
type EVMExpense = {
  readonly spender?: EVMAddress
  readonly token?: EVMAddress
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
        readonly target: EVMAddress
      }
    : {
        readonly target: EVMAddress
      })

type NativeTransfer = {
  readonly abi?: Abi // allows auto-complete abi, type will narrow if abi is provided
  readonly target: EVMAddress
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

export type EVMOrder<abis extends OptionalAbis> = {
  readonly owner?: EVMAddress
  readonly srcChainId?: number
  readonly destChainId: number
  readonly fillDeadline?: number
  readonly calls: Calls<abis>
  readonly deposit: EVMDeposit
  readonly expense: EVMExpense
}

// isContractCall call narrows a call type based on the presence of abi
// this is okay, because when abi is defined, call type is narrowed to ContractCall
export function isContractCall<
  abi extends OptionalAbi,
  defined extends NonNullable<abi> = NonNullable<abi>,
>(call: NativeTransfer | ContractCall<defined>): call is ContractCall<defined> {
  return call.abi != null
}
