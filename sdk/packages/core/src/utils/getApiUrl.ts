import { apiUrls } from '../constants/apis.js'
import type { Environment } from '../types/config.js'

export function getApiUrl(
  envOrApiBaseUrl: Environment | string = 'mainnet',
): string {
  return apiUrls[envOrApiBaseUrl as Environment] ?? envOrApiBaseUrl
}
