export class BaseError extends Error {
  override name = 'BaseError'
  reason?: string

  constructor(message: string, reason?: string) {
    super(message)
    this.reason = reason
  }
}

// TODO narrow types and add explicit reason values
export class OpenError extends BaseError {
  override name = 'OpenError' as const
}

export class TxReceiptError extends BaseError {
  override name = 'TxReceiptError' as const
}

export class GetOrderError extends BaseError {
  override name = 'GetOrderError' as const
}

export class DidFillError extends BaseError {
  override name = 'DidFillError' as const
}

export class ValidateOrderError extends BaseError {
  override name = 'ValidateOrderError' as const
}

export class ParseOpenEventError extends BaseError {
  override name = 'ParseOpenEventError' as const
}

export class LoadContractsError extends BaseError {
  override name = 'LoadContractsError' as const
}
