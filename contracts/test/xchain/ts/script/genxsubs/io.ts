import fs from 'fs'
import { decodeXBlock, encodeXSub, xBlockJson, xSubJson } from '../../xtypes'
import { NamedXSub, NamedXBlock } from './types'

// paths match those defined in test/common/Fixture.sol
const XBLOCKS_PATH = 'test/xchain/data/xblocks.json'
const XSUBS_PATH = 'test/xchain/data/xsubs.json'

// to write decoded xblocks and xsubs, for debugging
const XBLOCKS_DECODED_PATH = 'test/xchain/data/xblocks_decoded.json'
const XSUBS_DECODED_PATH = 'test/xchain/data/xsubs_decoded.json'

export function readXBlocks(): NamedXBlock[] {
  const input = JSON.parse(fs.readFileSync(XBLOCKS_PATH, 'utf8'))
  return Object.keys(input).map(key => ({ name: key, xblock: decodeXBlock(input[key]) }))
}

export function writeXBlocksDecoded(xblocks: NamedXBlock[]) {
  const output = Object.fromEntries(xblocks.map(b => [b.name, xBlockJson(b.xblock)]))
  fs.writeFileSync(XBLOCKS_DECODED_PATH, JSON.stringify(output, null, 2))
}

export function writeXSubsDecoded(xsubs: NamedXSub[]) {
  const output = Object.fromEntries(xsubs.map(s => [s.name, xSubJson(s.xsub)]))
  fs.writeFileSync(XSUBS_DECODED_PATH, JSON.stringify(output, null, 2))
}

export function writeXSubs(xsubs: NamedXSub[]) {
  const output = Object.fromEntries(xsubs.map(s => [s.name, encodeXSub(s.xsub)]))
  fs.writeFileSync(XSUBS_PATH, JSON.stringify(output, null, 2))
}
