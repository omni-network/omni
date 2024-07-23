/**
 * ts implementation of Omni's XBlock merkle tree. Used for test purposes only,
 * to generate test XSubmissions.
 *
 * Mirrors go implementation in lib/xchain/merkle.go.
 * Uses @openzeppelin/merkle-tree for core merkle tree implementation.
 */

import * as oz from '@openzeppelin/merkle-tree/dist/core'
import { keccak256 } from 'ethereum-cryptography/keccak'
import { equalsBytes, hexToBytes } from 'ethereum-cryptography/utils'
import { Bytes } from '@openzeppelin/merkle-tree/dist/bytes'
import { encodeXMsg, encodeXBlockHeader, XMsg, XBlockHeader } from './xtypes'

type Tree = readonly Bytes[]

type MultiProof = {
  leaves: Bytes[]
  proof: Bytes[]
  proofFlags: boolean[]
}

// XBlockMerkleTree, static methods only
export class XMsgMerkleTree {
  // Make an xblock tree from the provided header and messages
  static of(msgs: readonly XMsg[]): Tree {
    return oz.makeMerkleTree(leafHashes(msgs))
  }

  // Returns a multi proof provided messages, against the provided tree
  static prove(tree: Tree, msgs: readonly XMsg[]): MultiProof {
    const leaves = leafHashes(msgs)
    const indices = leaves.map(l => leafIndex(tree, l))
    return oz.getMultiProof([...tree], indices)
  }

  // Verify a multi proof
  static verify(root: Bytes, p: MultiProof): boolean {
    const impliedRoot = oz.processMultiProof(p)
    return equalsBytes(impliedRoot, root)
  }

  // Get the root of a tree
  static root(tree: Tree): Bytes {
    return tree[0]
  }

  // Order msgs by destination chain id, then offset
  static order(msgs: readonly XMsg[]): readonly XMsg[] {
    // Number(bigint) cast okay, because test chain ids and offsets are small
    return [...msgs].sort((a, b) => {
      if (a.destChainId !== b.destChainId) return Number(a.destChainId - b.destChainId)
      return Number(a.offset - b.offset)
    })
  }
}

const DST_XBLOCK_HEADER = 1
const DST_XMSG = 2

export const attestationRoot = (header: XBlockHeader, msgRoot: Bytes) =>
  oz.makeMerkleTree([msgRoot, headerLeafHash(header)])[0]
const stdLeafHash = (dst: number, data: `0x${string}`) => keccak256(keccak256(new Uint8Array([dst, ...hexToBytes(data)])));

const msgLeafHash = (msg: XMsg) => stdLeafHash(DST_XMSG, encodeXMsg(msg))
const headerLeafHash = (header: XBlockHeader) => stdLeafHash(DST_XBLOCK_HEADER, encodeXBlockHeader(header))

function leafIndex(tree: readonly Bytes[], leaf: Bytes) {
  const index = tree.findIndex(node => equalsBytes(node, leaf))
  if (index === -1) throwError('Leaf is not in tree')
  return index
}

function checkMsgOrder(msgs: readonly XMsg[]) {
  let lastSeenDestChainId = 0n
  let lastSeenOffset = 0n

  for (const msg of msgs) {
    if (msg.destChainId < lastSeenDestChainId) throwError('Msgs not ordered by dest chain id')
    if (msg.destChainId === lastSeenDestChainId && msg.offset < lastSeenOffset)
      throwError('Msgs not ordered by stream offset')

    lastSeenDestChainId = msg.destChainId
    lastSeenOffset = msg.offset
  }
}

function leafHashes(msgs: readonly XMsg[]) {
  checkMsgOrder(msgs)
  return [...msgs.map(msgLeafHash)]
}

function throwError(message?: string) {
  throw new Error(message)
}
