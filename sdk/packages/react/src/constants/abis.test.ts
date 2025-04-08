import { readFile } from 'node:fs/promises'
import { dirname, join } from 'node:path'
import { fileURLToPath } from 'node:url'
import { expect, test } from 'vitest'

import { inboxABI, middlemanABI, outboxABI } from './abis.js'

const ASSETS_PATH = fileURLToPath(
  new URL('../test/assets', dirname(import.meta.url)),
)

async function readContractFile(
  name: string,
): Promise<Array<Record<string, unknown>>> {
  const contents = await readFile(join(ASSETS_PATH, `${name}.json`), 'utf8')
  return JSON.parse(contents)
}

test.each([
  ['SolverNetInbox', inboxABI],
  ['SolverNetMiddleman', middlemanABI],
  ['SolverNetOutbox', outboxABI],
])(
  '%s contract contains all the ABIs expected by the SDK',
  async (name, expectedABI) => {
    const contract = await readContractFile(name)
    for (const expected of expectedABI) {
      const match = contract.find((abi) => abi.name === expected.name)
      expect(match).toEqual(expected)
    }
  },
)
