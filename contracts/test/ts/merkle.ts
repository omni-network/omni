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
export class XBlockMerkleTree {
  // Make an xblock tree from the provided header and messages
  static of(header: XBlockHeader, msgs: readonly XMsg[]): Tree {
    return oz.makeMerkleTree(leafHashes(header, msgs))
  }

  // Returns a multi proof provided header and messages, againts the provided tree
  static prove(tree: Tree, header: XBlockHeader, msgs: readonly XMsg[]): MultiProof {
    const leaves = leafHashes(header, msgs)
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
      return Number(a.streamOffset - b.streamOffset)
    })
  }
}

const stdLeafHash = (data: `0x${string}`) => keccak256(keccak256(hexToBytes(data)))
const msgLeafHash = (msg: XMsg) => stdLeafHash(encodeXMsg(msg))
const headerLeafHash = (header: XBlockHeader) => stdLeafHash(encodeXBlockHeader(header))

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
    if (msg.destChainId === lastSeenDestChainId && msg.streamOffset < lastSeenOffset)
      throwError('Msgs not ordered by stream offset')

    lastSeenDestChainId = msg.destChainId
    lastSeenOffset = msg.streamOffset
  }
}

function leafHashes(header: XBlockHeader, msgs: readonly XMsg[]) {
  checkMsgOrder(msgs)
  return [headerLeafHash(header), ...msgs.map(msgLeafHash)]
}

function throwError(message?: string) {
  throw new Error(message)
}
