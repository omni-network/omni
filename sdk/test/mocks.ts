import { vi } from 'vitest'
import * as apiModule from '../src/internal/api.js'
import { mockContracts } from './shared.js'

export function mockContractsQuery() {
  vi.spyOn(apiModule, 'fetchJSON').mockImplementation((url: string) => {
    if (url.includes('/contracts')) {
      return Promise.resolve(mockContracts)
    }
    return apiModule.fetchJSON(url)
  })
}
