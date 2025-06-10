import type { EVMOrder, OptionalAbis } from '@omni-network/core'
import { expectTypeOf, test } from 'vitest'
import { accounts } from '../../test/index.js'
import { useValidateOrder } from './useValidateOrder.js'

test('type: useInboxStatus', () => {
  const result = useValidateOrder({
    order: {
      owner: accounts[0],
      srcChainId: 1,
      destChainId: 2,
      calls: [
        {
          target: accounts[0],
          value: 0n,
        },
      ],
      deposit: {
        amount: 0n,
      },
      expense: {
        amount: 0n,
      },
    },
    enabled: true,
  })

  expectTypeOf(useValidateOrder).parameter(0).toMatchTypeOf<{
    order: EVMOrder<OptionalAbis>
    enabled: boolean
  }>()

  expectTypeOf(result.status).toEqualTypeOf<
    'pending' | 'rejected' | 'accepted' | 'error'
  >()
  expectTypeOf(result.status === 'rejected' ? result.rejectReason : undefined)
    .toBeString
  expectTypeOf(
    result.status === 'rejected' ? result.rejectDescription : undefined,
  ).toBeString
  expectTypeOf(
    result.status === 'error' ? result.error : undefined,
  ).not.toBeUndefined()
})
