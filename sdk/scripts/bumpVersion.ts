// updates package version.ts files to use version in code w/out importing package.json

import path from 'node:path'

console.log('bumping version files...')

// get package.json files
const packagePaths = await Array.fromAsync(
  new Bun.Glob('**/package.json').scan(),
)

let count = 0
for (const packagePath of packagePaths) {
  type Package = {
    name?: string | undefined
    private?: boolean | undefined
    version?: string | undefined
  }
  const file = Bun.file(packagePath)
  const packageJson = (await file.json()) as Package

  // skip private
  if (packageJson.private) continue

  count += 1
  console.log(`${packageJson.name} — ${packageJson.version}`)

  const versionFilePath = path.resolve(
    path.dirname(packagePath),
    'src',
    'version.ts',
  )
  await Bun.write(
    versionFilePath,
    `export const version = '${packageJson.version}'\n`,
  )
}

console.log('done bumping version files...')
