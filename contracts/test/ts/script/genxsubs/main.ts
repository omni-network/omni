/**
 * Generate XSubmissions for XBlocks with valid attestation roots and merkle proofs.
 */

import { XSub } from '../../xtypes'
import { XBlockMerkleTree } from '../../merkle'
import { NamedXBlock, NamedXSub } from './types'
import { bytesToHex } from 'viem'
import { groupByDestChain } from './utils'
import { readXBlocks, writeXBlocksDecoded, writeXSubs, writeXSubsDecoded } from './io'

// Get a XSubs for each destination chain in the XBlock.
function getXSubs(b: NamedXBlock) {
  const msgs = XBlockMerkleTree.order(b.xblock.msgs)
  const tree = XBlockMerkleTree.of(b.xblock.blockHeader, msgs)
  const byDestChain = groupByDestChain(msgs)

  const xsubs: NamedXSub[] = []
  for (const [destChainId, msgs] of Object.entries(byDestChain)) {
    const proof = XBlockMerkleTree.prove(tree, b.xblock.blockHeader, msgs)

    const xsub: XSub = {
      attestationRoot: bytesToHex(XBlockMerkleTree.root(tree)),
      validatorSetId: 1n, // validatorSetId set added in contract tests
      blockHeader: b.xblock.blockHeader,
      msgs,
      proof: proof.proof.map(p => bytesToHex(p)),
      proofFlags: proof.proofFlags,
      signatures: [], // signatures over attestationRoot added in contract tests
    }

    xsubs.push({ name: xsubName(b.name, destChainId), xsub })
  }

  return xsubs
}

// matches xsub name referenced in common/Fixture.sol:readXSubmission
const xsubName = (xblockName: string, destChainId: string) =>
  [xblockName, 'xsub', 'destChainId' + destChainId].join('_')

function main() {
  const xblocks = readXBlocks()
  const xsubs = xblocks.flatMap(getXSubs)
  writeXSubs(xsubs)
  writeXBlocksDecoded(xblocks)
  writeXSubsDecoded(xsubs)
}

main()
