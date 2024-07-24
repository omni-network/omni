/**
 * xtype defintions, with encoding / decoding utilities.
 */

import { decodeAbiParameters, encodeAbiParameters } from 'viem'
import { xBlockHeaderParams, xMsgParams, xSubParams, xBlockParams } from './abis'

type HexStr = `0x${string}`

export type XSub = ReturnType<typeof decodeXSub>
export type XMsg = ReturnType<typeof decodeXMsg>
export type XBlockHeader = ReturnType<typeof decodeXBlockHeader>
export type XBlock = ReturnType<typeof decodeXBlock>

export const decodeXSub = (data: HexStr) => decodeAbiParameters([xSubParams], data)[0]
export const decodeXMsg = (data: HexStr) => decodeAbiParameters([xMsgParams], data)[0]
export const decodeXBlock = (data: HexStr) => decodeAbiParameters([xBlockParams], data)[0]
export const decodeXBlockHeader = (data: HexStr) =>
  decodeAbiParameters([xBlockHeaderParams], data)[0]

export const encodeXSub = (xsub: XSub) => encodeAbiParameters([xSubParams], [xsub])
export const encodeXMsg = (xmsg: XMsg) => encodeAbiParameters([xMsgParams], [xmsg])
export const encodeXBlock = (xblock: XBlock) => encodeAbiParameters([xBlockParams], [xblock])
export const encodeXBlockHeader = (xblockHeader: XBlockHeader) =>
  encodeAbiParameters([xBlockHeaderParams], [xblockHeader])

// JSON.stringify cannot serialized bigints, so we convert them to strings

export const xMsgJson = (xmsg: XMsg) => ({
  destChainId: xmsg.destChainId.toString(),
  offset: xmsg.offset.toString(),
  sender: xmsg.sender,
  to: xmsg.to,
  data: xmsg.data,
  gasLimit: xmsg.gasLimit.toString(),
})

export const xBlockHeaderJson = (xblockHeader: XBlockHeader) => ({
  sourceChainId: xblockHeader.sourceChainId.toString(),
  consensusChainId: xblockHeader.consensusChainId.toString(),
  confLevel: xblockHeader.confLevel.toString(),
  offset: xblockHeader.offset.toString(),
  sourceBlockHash: xblockHeader.sourceBlockHash,
})

export const xSubJson = (xsub: XSub) => ({
  attestationRoot: xsub.attestationRoot,
  validatorSetId: xsub.validatorSetId.toString(),
  blockHeader: xBlockHeaderJson(xsub.blockHeader),
  msgs: xsub.msgs.map(xMsgJson),
  proof: xsub.proof,
  proofFlags: xsub.proofFlags,
  signatures: xsub.signatures,
})

export const xBlockJson = (xblock: XBlock) => ({
  blockHeader: xBlockHeaderJson(xblock.blockHeader),
  msgs: xblock.msgs.map(xMsgJson),
})
