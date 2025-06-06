import { readFile, readdir, writeFile } from 'node:fs/promises'
import { join } from 'node:path'
import { fileURLToPath } from 'node:url'
import { rootNodeFromAnchor } from '@codama/nodes-from-anchor'
import { renderJavaScriptVisitor } from '@codama/renderers'
import { createFromRoot } from 'codama'
import * as recast from 'recast'
import * as typescriptParser from 'recast/parsers/typescript.js'

const inboxIDLPath = fileURLToPath(
  new URL('../../../../anchor/localnet/solver_inbox.json', import.meta.url),
)
const inboxClientPath = fileURLToPath(
  new URL('../src/__generated__/inbox', import.meta.url),
)

async function rewriteFileExtensions(
  rootPath: string,
  filePath: string,
): Promise<void> {
  const rootFile = join(rootPath, 'index.ts')
  const code = await readFile(filePath, 'utf8')
  const ast = recast.parse(code, { parser: typescriptParser })
  recast.types.visit(ast, {
    visitExportAllDeclaration(pathNode) {
      const source = pathNode.node.source
      if (typeof source?.value === 'string') {
        if (source.value.startsWith('./')) {
          // In root file, exports are directories, so we need to add /index.js
          source.value += filePath === rootFile ? '/index.js' : '.js'
        }
      }
      this.traverse(pathNode)
    },
    visitImportDeclaration(pathNode) {
      const source = pathNode.node.source
      if (typeof source?.value === 'string') {
        if (source.value.startsWith('.')) {
          source.value += '/index.js'
        }
      }
      this.traverse(pathNode)
    },
  })
  const newCode = recast.print(ast).code
  await writeFile(filePath, newCode, 'utf8')
}

async function rewriteExtensions(
  rootPath: string,
  dirPath = rootPath,
): Promise<void> {
  const entries = await readdir(dirPath, { withFileTypes: true })
  await Promise.all(
    entries.map(async (entry) => {
      if (entry.isDirectory()) {
        await rewriteExtensions(rootPath, join(dirPath, entry.name))
      } else if (entry.isFile()) {
        await rewriteFileExtensions(rootPath, join(dirPath, entry.name))
      }
    }),
  )
}

async function generateClient(
  idlPath: string,
  clientPath: string,
): Promise<void> {
  const idl = JSON.parse(await readFile(idlPath, 'utf8'))
  const codama = createFromRoot(rootNodeFromAnchor(idl))
  await codama.accept(renderJavaScriptVisitor(clientPath))
  await rewriteExtensions(clientPath)
}

await generateClient(inboxIDLPath, inboxClientPath)
