import { expect, it } from 'vitest'

import * as lib from './index.js'

it('exports the expected APIs', () => {
  expect(Object.keys(lib)).toEqual([
    'withExecAndTransfer',
    'useOmniContracts',
    'useOrder',
    'useQuote',
    'useValidateOrder',
    'useGetOrderStatus',
    'useParseOpenEvent',
    'OmniProvider',
    'useOmniContext',
  ])
})
