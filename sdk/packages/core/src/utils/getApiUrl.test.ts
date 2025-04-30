import { describe, expect, test } from 'vitest'
import { apiUrls } from '../constants/apis.js'
import { getApiUrl } from './getApiUrl.js'

describe('getApiUrl', () => {
  test('default: returns the correct API URL for a valid environment', () => {
    expect(getApiUrl('devnet')).toBe(apiUrls.devnet)
  })

  test('behaviour: returns the mainnet API URL by default', () => {
    expect(getApiUrl()).toBe(apiUrls.mainnet)
  })

  test('behaviour: returns provided API URL by default', () => {
    expect(getApiUrl('http://localhost:3000')).toBe('http://localhost:3000')
  })
})
