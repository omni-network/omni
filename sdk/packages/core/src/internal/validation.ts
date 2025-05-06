import type { ZodMiniType, core } from '@zod/mini'
import { type AsyncResult, Result } from 'typescript-result'

export type ValidationIssue = core.$ZodIssue
export type ValidationSchema<O = unknown, I = O> = ZodMiniType<O, I>

export type ValidationErrorParams<O = unknown, I = O> = {
  schema: ValidationSchema<O, I>
  input: unknown
  issues: ValidationIssue[]
  message?: string
}

export class ValidationError<O = unknown, I = O> extends Error {
  readonly type = 'SchemaParseError'
  readonly schema: ValidationSchema<O, I>
  readonly input: unknown
  readonly issues: ValidationIssue[]

  constructor(params: ValidationErrorParams<O, I>) {
    super(params.message ?? 'Schema validation failed')
    this.schema = params.schema
    this.input = params.input
    this.issues = params.issues
  }
}

export function safeValidate<O, I = O>(
  schema: ValidationSchema<O, I>,
  input: unknown,
): Result<O, ValidationError<O, I>> {
  const parsed = schema.safeParse(input)
  return parsed.success
    ? Result.ok(parsed.data)
    : Result.error(
        new ValidationError({ schema, input, issues: parsed.error.issues }),
      )
}

export function safeValidateAsync<O, I = O>(
  schema: ValidationSchema<O, I>,
  input: unknown,
): AsyncResult<O, ValidationError<O, I>> {
  return Result.fromAsync(Promise.resolve(input))
    .map(schema.safeParseAsync)
    .map((parsed) => {
      return parsed.success
        ? Result.ok(parsed.data)
        : Result.error(
            new ValidationError({ schema, input, issues: parsed.error.issues }),
          )
    })
}
