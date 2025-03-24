import {
  type InferInput,
  array,
  bigint,
  hexadecimal,
  literal,
  number,
  object,
  optional,
  pipe,
  startsWith,
  string,
  union,
} from 'valibot'

export const hexSchema = pipe(
  string(),
  startsWith('0x', 'The hexadecimal must start with "0x".'),
  hexadecimal('The hexadecimal is badly formatted.'),
)
export type Hex = `0x${string}`

export const orderCallSchema = object({
  target: hexSchema,
  value: optional(bigint()),
})

export const orderDepositSchema = object({
  token: hexSchema,
  amount: bigint(),
})

export const orderExpenseSchema = object({
  spender: hexSchema,
  token: hexSchema,
  amount: bigint(),
})

export const orderSchema = object({
  owner: hexSchema,
  srcChainId: optional(number()),
  destChainId: number(),
  fillDeadline: number(),
  calls: array(orderCallSchema),
  deposit: orderDepositSchema,
  expense: orderExpenseSchema,
})
export type Order = InferInput<typeof orderSchema>

export const quoteEntrySchema = object({
  token: hexSchema,
  amount: hexSchema,
})

export const quoteSchema = object({
  deposit: quoteEntrySchema,
  expense: quoteEntrySchema,
})
export type Quote = InferInput<typeof quoteSchema>

export const requestErrorSchema = object({
  code: number(),
  message: string(),
})

export const contractAddressesSchema = object({
  inbox: hexSchema,
  outbox: hexSchema,
  middleman: hexSchema,
})
export type ContractAddresses = InferInput<typeof contractAddressesSchema>

export const validationAcceptedSchema = object({
  status: literal('accepted'),
})

export const validationRejectedSchema = object({
  status: literal('rejected'),
  rejectReason: string(),
  rejectDescription: string(),
})

export const validationErrorSchema = object({
  error: requestErrorSchema,
})

export const validationResultSchema = union([
  validationAcceptedSchema,
  validationRejectedSchema,
  validationErrorSchema,
])
export type ValidationResult = InferInput<typeof validationResultSchema>
