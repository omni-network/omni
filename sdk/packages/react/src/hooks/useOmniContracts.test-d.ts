import type { OmniContracts } from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import { expectTypeOf, test } from 'vitest'
import {
  type UseOmniContractsResult,
  useOmniContracts,
} from './useOmniContracts.js'

test('type: useOmniContracts', () => {
  const result = useOmniContracts()
  expectTypeOf(result).toEqualTypeOf<UseOmniContractsResult>()
  expectTypeOf(result).toEqualTypeOf<UseQueryResult<OmniContracts>>()

  expectTypeOf(result.data).toEqualTypeOf<OmniContracts | undefined>()

  if (result.data) {
    expectTypeOf(result.data.inbox).toBeString()
    expectTypeOf(result.data.outbox).toBeString()
  }
})
