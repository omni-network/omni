import { expect, it } from 'vitest'
import * as lib from './index.js'

it('exports the expected methods', () => {
  expect(Object.keys(lib)).toEqual(['OmniProvider', 'useOmniContext'])
})
