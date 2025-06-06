/**
 * This code was AUTOGENERATED using the codama library.
 * Please DO NOT EDIT THIS FILE, instead use visitors
 * to add features, then rerun codama to update it.
 *
 * @see https://github.com/codama-idl/codama
 */

import {
  type Address,
  type ReadonlyUint8Array,
  containsBytes,
  fixEncoderSize,
  getBytesEncoder,
} from '@solana/kit'
import type {
  ParsedClaimInstruction,
  ParsedCloseInstruction,
  ParsedInitInstruction,
  ParsedMarkFilledInstruction,
  ParsedOpenInstruction,
  ParsedRejectInstruction,
} from '../instructions/index.js'

export const SOLVER_INBOX_PROGRAM_ADDRESS =
  'AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv' as Address<'AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv'>

export enum SolverInboxAccount {
  InboxState = 0,
  OrderState = 1,
}

export function identifySolverInboxAccount(
  account: { data: ReadonlyUint8Array } | ReadonlyUint8Array,
): SolverInboxAccount {
  const data = 'data' in account ? account.data : account
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([161, 5, 9, 33, 125, 185, 63, 116]),
      ),
      0,
    )
  ) {
    return SolverInboxAccount.InboxState
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([60, 123, 67, 162, 96, 43, 173, 225]),
      ),
      0,
    )
  ) {
    return SolverInboxAccount.OrderState
  }
  throw new Error(
    'The provided account could not be identified as a solverInbox account.',
  )
}

export enum SolverInboxInstruction {
  Claim = 0,
  Close = 1,
  Init = 2,
  MarkFilled = 3,
  Open = 4,
  Reject = 5,
}

export function identifySolverInboxInstruction(
  instruction: { data: ReadonlyUint8Array } | ReadonlyUint8Array,
): SolverInboxInstruction {
  const data = 'data' in instruction ? instruction.data : instruction
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([62, 198, 214, 193, 213, 159, 108, 210]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.Claim
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([98, 165, 201, 177, 108, 65, 206, 96]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.Close
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([220, 59, 207, 236, 108, 250, 47, 100]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.Init
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([192, 137, 170, 0, 70, 5, 127, 160]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.MarkFilled
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([228, 220, 155, 71, 199, 189, 60, 45]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.Open
  }
  if (
    containsBytes(
      data,
      fixEncoderSize(getBytesEncoder(), 8).encode(
        new Uint8Array([135, 7, 63, 85, 131, 114, 111, 224]),
      ),
      0,
    )
  ) {
    return SolverInboxInstruction.Reject
  }
  throw new Error(
    'The provided instruction could not be identified as a solverInbox instruction.',
  )
}

export type ParsedSolverInboxInstruction<
  TProgram extends string = 'AwminMpVyPSX86m3w9KWcxjovtnjwxiKZUNTDgDqrctv',
> =
  | ({
      instructionType: SolverInboxInstruction.Claim
    } & ParsedClaimInstruction<TProgram>)
  | ({
      instructionType: SolverInboxInstruction.Close
    } & ParsedCloseInstruction<TProgram>)
  | ({
      instructionType: SolverInboxInstruction.Init
    } & ParsedInitInstruction<TProgram>)
  | ({
      instructionType: SolverInboxInstruction.MarkFilled
    } & ParsedMarkFilledInstruction<TProgram>)
  | ({
      instructionType: SolverInboxInstruction.Open
    } & ParsedOpenInstruction<TProgram>)
  | ({
      instructionType: SolverInboxInstruction.Reject
    } & ParsedRejectInstruction<TProgram>)
