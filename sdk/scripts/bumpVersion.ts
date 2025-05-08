// updates package version.ts files to use version in code w/out importing package.json

import fs from 'node:fs/promises'
import path from 'node:path'

console.log('bumping version files...')

const packagePaths = [
  'packages/react/package.json',
  'packages/core/package.json',
]

let count = 0
for (const packagePath of packagePaths) {
  type Package = {
    name?: string | undefined
    private?: boolean | undefined
    version?: string | undefined
  }
  const fileContents = await fs.readFile(packagePath, 'utf-8')
  const packageJson = JSON.parse(fileContents) as Package

  // skip private
  if (packageJson.private) continue

  count += 1
  console.log(`${packageJson.name} — ${packageJson.version}`)

  const versionFilePath = path.resolve(
    path.dirname(packagePath),
    'src',
    'version.ts',
  )
  await fs.writeFile(
    versionFilePath,
    `export const version = '${packageJson.version}'\n`,
  )
}

console.log('done bumping version files...')
