import { expect, it } from 'vitest'

import * as lib from './index.js'

it('exports the expected APIs', () => {
  expect(Object.keys(lib)).toEqual([
    'useOmniContracts',
    'useOrder',
    'useQuote',
    'useValidateOrder',
    'useGetOrder',
    'useGetOrderStatus',
    'useParseOpenEvent',
    'useWatchDidFill',
    'OmniProvider',
    'useOmniContext',
  ])
})
