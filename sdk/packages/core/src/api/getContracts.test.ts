import { testContracts } from '@omni-network/test-utils'
import { Result } from 'typescript-result'
import { beforeEach, expect, test, vi } from 'vitest'
import { ValidationError, safeValidate } from '../internal/validation.js'
import { omniContractsSchema } from '../types/contracts.js'

const { createSafeFetchType } = vi.hoisted(() => ({
  createSafeFetchType: vi.fn(),
}))
vi.mock('../internal/api.js', () => {
  return { createSafeFetchType }
})

beforeEach(() => {
  // ensures import("./getContracts.js") gets re-evaluated with the wanted mock
  vi.resetModules()
})

test('default: returns contracts addresses', async () => {
  createSafeFetchType.mockImplementationOnce(() => {
    return (url: string) => {
      expect(url).toBe('http://localhost/contracts')
      return Result.ok(testContracts)
    }
  })
  const { getContracts } = await import('./getContracts.js')
  await expect(getContracts('http://localhost')).resolves.toEqual(testContracts)
})

test('behaviour: handles invalid response format', async () => {
  createSafeFetchType.mockImplementationOnce(() => {
    return (url: string) => {
      expect(url).toBe('http://localhost/contracts')
      return safeValidate(omniContractsSchema, { invalidField: 'value' })
    }
  })
  const { getContracts } = await import('./getContracts.js')
  const expectRejection = expect(async () => {
    await getContracts('http://localhost')
  }).rejects
  await expectRejection.toBeInstanceOf(ValidationError)
  await expectRejection.toHaveProperty('message', 'Schema validation failed')
})
