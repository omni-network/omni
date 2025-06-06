/**
 * This code was AUTOGENERATED using the codama library.
 * Please DO NOT EDIT THIS FILE, instead use visitors
 * to add features, then rerun codama to update it.
 *
 * @see https://github.com/codama-idl/codama
 */

import {
  type Address,
  type Codec,
  type Decoder,
  type Encoder,
  type IAccountMeta,
  type IAccountSignerMeta,
  type IInstruction,
  type IInstructionWithAccounts,
  type IInstructionWithData,
  type ReadonlyAccount,
  type ReadonlyUint8Array,
  type TransactionSigner,
  type WritableAccount,
  type WritableSignerAccount,
  combineCodec,
  fixDecoderSize,
  fixEncoderSize,
  getAddressDecoder,
  getAddressEncoder,
  getBytesDecoder,
  getBytesEncoder,
  getProgramDerivedAddress,
  getStructDecoder,
  getStructEncoder,
  getU64Decoder,
  getU64Encoder,
  transformEncoder,
} from '@solana/kit'
import { SOLVER_INBOX_PROGRAM_ADDRESS } from '../programs/index.js'
import { type ResolvedAccount, getAccountMetaFactory } from '../shared/index.js'
import {
  type EVMCall,
  type EVMCallArgs,
  type EVMTokenExpense,
  type EVMTokenExpenseArgs,
  getEVMCallDecoder,
  getEVMCallEncoder,
  getEVMTokenExpenseDecoder,
  getEVMTokenExpenseEncoder,
} from '../types/index.js'

export const OPEN_DISCRIMINATOR = new Uint8Array([
  228, 220, 155, 71, 199, 189, 60, 45,
])

export function getOpenDiscriminatorBytes() {
  return fixEncoderSize(getBytesEncoder(), 8).encode(OPEN_DISCRIMINATOR)
}

export type OpenInstruction<
  TProgram extends string = typeof SOLVER_INBOX_PROGRAM_ADDRESS,
  TAccountOrderState extends string | IAccountMeta<string> = string,
  TAccountOwner extends string | IAccountMeta<string> = string,
  TAccountMintAccount extends string | IAccountMeta<string> = string,
  TAccountOwnerTokenAccount extends string | IAccountMeta<string> = string,
  TAccountOrderTokenAccount extends string | IAccountMeta<string> = string,
  TAccountTokenProgram extends
    | string
    | IAccountMeta<string> = 'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA',
  TAccountInboxState extends string | IAccountMeta<string> = string,
  TAccountSystemProgram extends
    | string
    | IAccountMeta<string> = '11111111111111111111111111111111',
  TRemainingAccounts extends readonly IAccountMeta<string>[] = [],
> = IInstruction<TProgram> &
  IInstructionWithData<Uint8Array> &
  IInstructionWithAccounts<
    [
      TAccountOrderState extends string
        ? WritableAccount<TAccountOrderState>
        : TAccountOrderState,
      TAccountOwner extends string
        ? WritableSignerAccount<TAccountOwner> &
            IAccountSignerMeta<TAccountOwner>
        : TAccountOwner,
      TAccountMintAccount extends string
        ? WritableAccount<TAccountMintAccount>
        : TAccountMintAccount,
      TAccountOwnerTokenAccount extends string
        ? WritableAccount<TAccountOwnerTokenAccount>
        : TAccountOwnerTokenAccount,
      TAccountOrderTokenAccount extends string
        ? WritableAccount<TAccountOrderTokenAccount>
        : TAccountOrderTokenAccount,
      TAccountTokenProgram extends string
        ? ReadonlyAccount<TAccountTokenProgram>
        : TAccountTokenProgram,
      TAccountInboxState extends string
        ? ReadonlyAccount<TAccountInboxState>
        : TAccountInboxState,
      TAccountSystemProgram extends string
        ? ReadonlyAccount<TAccountSystemProgram>
        : TAccountSystemProgram,
      ...TRemainingAccounts,
    ]
  >

export type OpenInstructionData = {
  discriminator: ReadonlyUint8Array
  orderId: Address
  nonce: bigint
  depositAmount: bigint
  destChainId: bigint
  call: EVMCall
  expense: EVMTokenExpense
}

export type OpenInstructionDataArgs = {
  orderId: Address
  nonce: number | bigint
  depositAmount: number | bigint
  destChainId: number | bigint
  call: EVMCallArgs
  expense: EVMTokenExpenseArgs
}

export function getOpenInstructionDataEncoder(): Encoder<OpenInstructionDataArgs> {
  return transformEncoder(
    getStructEncoder([
      ['discriminator', fixEncoderSize(getBytesEncoder(), 8)],
      ['orderId', getAddressEncoder()],
      ['nonce', getU64Encoder()],
      ['depositAmount', getU64Encoder()],
      ['destChainId', getU64Encoder()],
      ['call', getEVMCallEncoder()],
      ['expense', getEVMTokenExpenseEncoder()],
    ]),
    (value) => ({ ...value, discriminator: OPEN_DISCRIMINATOR }),
  )
}

export function getOpenInstructionDataDecoder(): Decoder<OpenInstructionData> {
  return getStructDecoder([
    ['discriminator', fixDecoderSize(getBytesDecoder(), 8)],
    ['orderId', getAddressDecoder()],
    ['nonce', getU64Decoder()],
    ['depositAmount', getU64Decoder()],
    ['destChainId', getU64Decoder()],
    ['call', getEVMCallDecoder()],
    ['expense', getEVMTokenExpenseDecoder()],
  ])
}

export function getOpenInstructionDataCodec(): Codec<
  OpenInstructionDataArgs,
  OpenInstructionData
> {
  return combineCodec(
    getOpenInstructionDataEncoder(),
    getOpenInstructionDataDecoder(),
  )
}

export type OpenAsyncInput<
  TAccountOrderState extends string = string,
  TAccountOwner extends string = string,
  TAccountMintAccount extends string = string,
  TAccountOwnerTokenAccount extends string = string,
  TAccountOrderTokenAccount extends string = string,
  TAccountTokenProgram extends string = string,
  TAccountInboxState extends string = string,
  TAccountSystemProgram extends string = string,
> = {
  orderState: Address<TAccountOrderState>
  owner: TransactionSigner<TAccountOwner>
  mintAccount: Address<TAccountMintAccount>
  ownerTokenAccount: Address<TAccountOwnerTokenAccount>
  orderTokenAccount: Address<TAccountOrderTokenAccount>
  tokenProgram?: Address<TAccountTokenProgram>
  inboxState?: Address<TAccountInboxState>
  systemProgram?: Address<TAccountSystemProgram>
  orderId: OpenInstructionDataArgs['orderId']
  nonce: OpenInstructionDataArgs['nonce']
  depositAmount: OpenInstructionDataArgs['depositAmount']
  destChainId: OpenInstructionDataArgs['destChainId']
  call: OpenInstructionDataArgs['call']
  expense: OpenInstructionDataArgs['expense']
}

export async function getOpenInstructionAsync<
  TAccountOrderState extends string,
  TAccountOwner extends string,
  TAccountMintAccount extends string,
  TAccountOwnerTokenAccount extends string,
  TAccountOrderTokenAccount extends string,
  TAccountTokenProgram extends string,
  TAccountInboxState extends string,
  TAccountSystemProgram extends string,
  TProgramAddress extends Address = typeof SOLVER_INBOX_PROGRAM_ADDRESS,
>(
  input: OpenAsyncInput<
    TAccountOrderState,
    TAccountOwner,
    TAccountMintAccount,
    TAccountOwnerTokenAccount,
    TAccountOrderTokenAccount,
    TAccountTokenProgram,
    TAccountInboxState,
    TAccountSystemProgram
  >,
  config?: { programAddress?: TProgramAddress },
): Promise<
  OpenInstruction<
    TProgramAddress,
    TAccountOrderState,
    TAccountOwner,
    TAccountMintAccount,
    TAccountOwnerTokenAccount,
    TAccountOrderTokenAccount,
    TAccountTokenProgram,
    TAccountInboxState,
    TAccountSystemProgram
  >
> {
  // Program address.
  const programAddress = config?.programAddress ?? SOLVER_INBOX_PROGRAM_ADDRESS

  // Original accounts.
  const originalAccounts = {
    orderState: { value: input.orderState ?? null, isWritable: true },
    owner: { value: input.owner ?? null, isWritable: true },
    mintAccount: { value: input.mintAccount ?? null, isWritable: true },
    ownerTokenAccount: {
      value: input.ownerTokenAccount ?? null,
      isWritable: true,
    },
    orderTokenAccount: {
      value: input.orderTokenAccount ?? null,
      isWritable: true,
    },
    tokenProgram: { value: input.tokenProgram ?? null, isWritable: false },
    inboxState: { value: input.inboxState ?? null, isWritable: false },
    systemProgram: { value: input.systemProgram ?? null, isWritable: false },
  }
  const accounts = originalAccounts as Record<
    keyof typeof originalAccounts,
    ResolvedAccount
  >

  // Original args.
  const args = { ...input }

  // Resolve default values.
  if (!accounts.tokenProgram.value) {
    accounts.tokenProgram.value =
      'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA' as Address<'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'>
  }
  if (!accounts.inboxState.value) {
    accounts.inboxState.value = await getProgramDerivedAddress({
      programAddress,
      seeds: [
        getBytesEncoder().encode(
          new Uint8Array([105, 110, 98, 111, 120, 95, 115, 116, 97, 116, 101]),
        ),
      ],
    })
  }
  if (!accounts.systemProgram.value) {
    accounts.systemProgram.value =
      '11111111111111111111111111111111' as Address<'11111111111111111111111111111111'>
  }

  const getAccountMeta = getAccountMetaFactory(programAddress, 'programId')
  const instruction = {
    accounts: [
      getAccountMeta(accounts.orderState),
      getAccountMeta(accounts.owner),
      getAccountMeta(accounts.mintAccount),
      getAccountMeta(accounts.ownerTokenAccount),
      getAccountMeta(accounts.orderTokenAccount),
      getAccountMeta(accounts.tokenProgram),
      getAccountMeta(accounts.inboxState),
      getAccountMeta(accounts.systemProgram),
    ],
    programAddress,
    data: getOpenInstructionDataEncoder().encode(
      args as OpenInstructionDataArgs,
    ),
  } as OpenInstruction<
    TProgramAddress,
    TAccountOrderState,
    TAccountOwner,
    TAccountMintAccount,
    TAccountOwnerTokenAccount,
    TAccountOrderTokenAccount,
    TAccountTokenProgram,
    TAccountInboxState,
    TAccountSystemProgram
  >

  return instruction
}

export type OpenInput<
  TAccountOrderState extends string = string,
  TAccountOwner extends string = string,
  TAccountMintAccount extends string = string,
  TAccountOwnerTokenAccount extends string = string,
  TAccountOrderTokenAccount extends string = string,
  TAccountTokenProgram extends string = string,
  TAccountInboxState extends string = string,
  TAccountSystemProgram extends string = string,
> = {
  orderState: Address<TAccountOrderState>
  owner: TransactionSigner<TAccountOwner>
  mintAccount: Address<TAccountMintAccount>
  ownerTokenAccount: Address<TAccountOwnerTokenAccount>
  orderTokenAccount: Address<TAccountOrderTokenAccount>
  tokenProgram?: Address<TAccountTokenProgram>
  inboxState: Address<TAccountInboxState>
  systemProgram?: Address<TAccountSystemProgram>
  orderId: OpenInstructionDataArgs['orderId']
  nonce: OpenInstructionDataArgs['nonce']
  depositAmount: OpenInstructionDataArgs['depositAmount']
  destChainId: OpenInstructionDataArgs['destChainId']
  call: OpenInstructionDataArgs['call']
  expense: OpenInstructionDataArgs['expense']
}

export function getOpenInstruction<
  TAccountOrderState extends string,
  TAccountOwner extends string,
  TAccountMintAccount extends string,
  TAccountOwnerTokenAccount extends string,
  TAccountOrderTokenAccount extends string,
  TAccountTokenProgram extends string,
  TAccountInboxState extends string,
  TAccountSystemProgram extends string,
  TProgramAddress extends Address = typeof SOLVER_INBOX_PROGRAM_ADDRESS,
>(
  input: OpenInput<
    TAccountOrderState,
    TAccountOwner,
    TAccountMintAccount,
    TAccountOwnerTokenAccount,
    TAccountOrderTokenAccount,
    TAccountTokenProgram,
    TAccountInboxState,
    TAccountSystemProgram
  >,
  config?: { programAddress?: TProgramAddress },
): OpenInstruction<
  TProgramAddress,
  TAccountOrderState,
  TAccountOwner,
  TAccountMintAccount,
  TAccountOwnerTokenAccount,
  TAccountOrderTokenAccount,
  TAccountTokenProgram,
  TAccountInboxState,
  TAccountSystemProgram
> {
  // Program address.
  const programAddress = config?.programAddress ?? SOLVER_INBOX_PROGRAM_ADDRESS

  // Original accounts.
  const originalAccounts = {
    orderState: { value: input.orderState ?? null, isWritable: true },
    owner: { value: input.owner ?? null, isWritable: true },
    mintAccount: { value: input.mintAccount ?? null, isWritable: true },
    ownerTokenAccount: {
      value: input.ownerTokenAccount ?? null,
      isWritable: true,
    },
    orderTokenAccount: {
      value: input.orderTokenAccount ?? null,
      isWritable: true,
    },
    tokenProgram: { value: input.tokenProgram ?? null, isWritable: false },
    inboxState: { value: input.inboxState ?? null, isWritable: false },
    systemProgram: { value: input.systemProgram ?? null, isWritable: false },
  }
  const accounts = originalAccounts as Record<
    keyof typeof originalAccounts,
    ResolvedAccount
  >

  // Original args.
  const args = { ...input }

  // Resolve default values.
  if (!accounts.tokenProgram.value) {
    accounts.tokenProgram.value =
      'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA' as Address<'TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA'>
  }
  if (!accounts.systemProgram.value) {
    accounts.systemProgram.value =
      '11111111111111111111111111111111' as Address<'11111111111111111111111111111111'>
  }

  const getAccountMeta = getAccountMetaFactory(programAddress, 'programId')
  const instruction = {
    accounts: [
      getAccountMeta(accounts.orderState),
      getAccountMeta(accounts.owner),
      getAccountMeta(accounts.mintAccount),
      getAccountMeta(accounts.ownerTokenAccount),
      getAccountMeta(accounts.orderTokenAccount),
      getAccountMeta(accounts.tokenProgram),
      getAccountMeta(accounts.inboxState),
      getAccountMeta(accounts.systemProgram),
    ],
    programAddress,
    data: getOpenInstructionDataEncoder().encode(
      args as OpenInstructionDataArgs,
    ),
  } as OpenInstruction<
    TProgramAddress,
    TAccountOrderState,
    TAccountOwner,
    TAccountMintAccount,
    TAccountOwnerTokenAccount,
    TAccountOrderTokenAccount,
    TAccountTokenProgram,
    TAccountInboxState,
    TAccountSystemProgram
  >

  return instruction
}

export type ParsedOpenInstruction<
  TProgram extends string = typeof SOLVER_INBOX_PROGRAM_ADDRESS,
  TAccountMetas extends readonly IAccountMeta[] = readonly IAccountMeta[],
> = {
  programAddress: Address<TProgram>
  accounts: {
    orderState: TAccountMetas[0]
    owner: TAccountMetas[1]
    mintAccount: TAccountMetas[2]
    ownerTokenAccount: TAccountMetas[3]
    orderTokenAccount: TAccountMetas[4]
    tokenProgram: TAccountMetas[5]
    inboxState: TAccountMetas[6]
    systemProgram: TAccountMetas[7]
  }
  data: OpenInstructionData
}

export function parseOpenInstruction<
  TProgram extends string,
  TAccountMetas extends readonly IAccountMeta[],
>(
  instruction: IInstruction<TProgram> &
    IInstructionWithAccounts<TAccountMetas> &
    IInstructionWithData<Uint8Array>,
): ParsedOpenInstruction<TProgram, TAccountMetas> {
  if (instruction.accounts.length < 8) {
    // TODO: Coded error.
    throw new Error('Not enough accounts')
  }
  let accountIndex = 0
  const getNextAccount = () => {
    const accountMeta = instruction.accounts![accountIndex]!
    accountIndex += 1
    return accountMeta
  }
  return {
    programAddress: instruction.programAddress,
    accounts: {
      orderState: getNextAccount(),
      owner: getNextAccount(),
      mintAccount: getNextAccount(),
      ownerTokenAccount: getNextAccount(),
      orderTokenAccount: getNextAccount(),
      tokenProgram: getNextAccount(),
      inboxState: getNextAccount(),
      systemProgram: getNextAccount(),
    },
    data: getOpenInstructionDataDecoder().decode(instruction.data),
  }
}
