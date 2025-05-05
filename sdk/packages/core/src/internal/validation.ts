import { type ZodMiniType, type core, regex, string } from '@zod/mini'
import { Result } from 'typescript-result'
import type { Hex } from 'viem'

export type ValidationIssue = core.$ZodIssue
export type ValidationSchema<T = unknown> = ZodMiniType<T, T>

export type ValidationErrorParams<T = unknown> = {
  schema: ValidationSchema<T>
  input: unknown
  issues: ValidationIssue[]
  message?: string
}

export class ValidationError<T = unknown> extends Error {
  readonly type = 'SchemaParseError'
  readonly schema: ValidationSchema<T>
  readonly input: unknown
  readonly issues: ValidationIssue[]

  constructor(params: ValidationErrorParams<T>) {
    super(params.message ?? 'Schema validation failed')
    this.schema = params.schema
    this.input = params.input
    this.issues = params.issues
  }
}

export function safeValidate<T>(
  schema: ValidationSchema<T>,
  input: unknown,
): Result<T, ValidationError<T>> {
  const parsed = schema.safeParse(input)
  return parsed.success
    ? Result.ok(parsed.data)
    : Result.error(
        new ValidationError({ schema, input, issues: parsed.error.issues }),
      )
}

export const hexStringSchema = string().check(
  regex(/0x[0-9A-Fa-f]{2,}/),
) as ValidationSchema<Hex>
