import { BaseError } from '@omni-network/core'

export class NoClientError extends BaseError {
  override name = 'NoClientError' as const
}
