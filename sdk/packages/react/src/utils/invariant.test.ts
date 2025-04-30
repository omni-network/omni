import { expect, it } from 'vitest'
import { invariant } from './invariant.js'

it('default: should not throw when condition is true', () => {
  expect(() => invariant(true, 'This should not throw')).not.toThrow()
})

it('default: should throw with the default msg when condition is false', () => {
  expect(() => invariant(false)).toThrowError(
    'Invariant violation: condition not met',
  )
})

it('behaviour: should throw with provided msg when condition is false', () => {
  const msg = 'Custom error message for testing'
  expect(() => invariant(false, msg)).toThrowError(msg)
})
