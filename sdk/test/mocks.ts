import { vi } from 'vitest'
import * as apiModule from '../src/internal/api.js'
import { contracts } from './shared.js'

export function mockContractsQuery() {
  vi.spyOn(apiModule, 'fetchJSON').mockImplementation((url: string) => {
    if (url.includes('/contracts')) {
      return Promise.resolve(contracts)
    }
    return apiModule.fetchJSON(url)
  })
}
