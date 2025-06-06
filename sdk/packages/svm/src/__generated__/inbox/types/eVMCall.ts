/**
 * This code was AUTOGENERATED using the codama library.
 * Please DO NOT EDIT THIS FILE, instead use visitors
 * to add features, then rerun codama to update it.
 *
 * @see https://github.com/codama-idl/codama
 */

import {
  type Codec,
  type Decoder,
  type Encoder,
  type ReadonlyUint8Array,
  addDecoderSizePrefix,
  addEncoderSizePrefix,
  combineCodec,
  fixDecoderSize,
  fixEncoderSize,
  getBytesDecoder,
  getBytesEncoder,
  getStructDecoder,
  getStructEncoder,
  getU32Decoder,
  getU32Encoder,
  getU128Decoder,
  getU128Encoder,
} from '@solana/kit'

/**
 * EVM call to execute on destination chain
 * If the call is a native transfer, `target` is the recipient address, and `selector` / `params` are empty.
 */

export type EVMCall = {
  target: ReadonlyUint8Array
  selector: ReadonlyUint8Array
  value: bigint
  params: ReadonlyUint8Array
}

export type EVMCallArgs = {
  target: ReadonlyUint8Array
  selector: ReadonlyUint8Array
  value: number | bigint
  params: ReadonlyUint8Array
}

export function getEVMCallEncoder(): Encoder<EVMCallArgs> {
  return getStructEncoder([
    ['target', fixEncoderSize(getBytesEncoder(), 20)],
    ['selector', fixEncoderSize(getBytesEncoder(), 4)],
    ['value', getU128Encoder()],
    ['params', addEncoderSizePrefix(getBytesEncoder(), getU32Encoder())],
  ])
}

export function getEVMCallDecoder(): Decoder<EVMCall> {
  return getStructDecoder([
    ['target', fixDecoderSize(getBytesDecoder(), 20)],
    ['selector', fixDecoderSize(getBytesDecoder(), 4)],
    ['value', getU128Decoder()],
    ['params', addDecoderSizePrefix(getBytesDecoder(), getU32Decoder())],
  ])
}

export function getEVMCallCodec(): Codec<EVMCallArgs, EVMCall> {
  return combineCodec(getEVMCallEncoder(), getEVMCallDecoder())
}
