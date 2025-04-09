import { useQuery } from '@tanstack/react-query'
import { client } from '../../client-modules/react-query-client'

export const queryOpts = (url: string) => ({
  queryKey: ['code-block', url],
  queryFn: () => getCodeBlock(url),
})

export function useCodeBlock({ url }: { url: string }) {
  return useQuery(queryOpts(url), client)
}

async function getCodeBlock(url: string) {
  const file = parseURL(url)
  const code = await getCode(file)
  const codeOnMain = await getCode(onMain(file))
  const matchesMain = code === codeOnMain

  return {
    code,
    matchesMain,
    mainURL: sourceURL(onMain(file)),
    sourceURL: sourceURL(file),
    language: determineLanguage(file.path),
  }
}

async function getCode(f: GithubFile) {
  const response = await fetch(rawURL(f))

  if (!response.ok) {
    throw new Error(
      `Failed to fetch code from GitHub: status=${response.statusText}`,
    )
  }

  const data = await response.text()
  const lines = data.split('\n')

  if (f.isSnippet) return lines.slice(f.startLine - 1, f.endLine).join('\n')
  return lines.join('\n')
}

const rawURL = (f: GithubFile) =>
  `https://raw.githubusercontent.com/${f.owner}/${f.repo}/${f.ref}/${f.path}`

const sourceURL = (f: GithubFile) =>
  f.isSnippet
    ? `https://github.com/${f.owner}/${f.repo}/blob/${f.ref}/${f.path}#L${f.startLine}-L${f.endLine}`
    : `https://github.com/${f.owner}/${f.repo}/blob/${f.ref}/${f.path}`

const onMain = (f: GithubFile) => ({ ...f, ref: 'main' })

type GithubFile =
  | {
      isSnippet: false
      owner: string
      repo: string
      ref: string
      path: string
    }
  | {
      isSnippet: true
      owner: string
      repo: string
      ref: string
      path: string
      startLine: number
      endLine: number
    }

function parseURL(url: string): GithubFile {
  if (isSnippetURL(url)) return parseSnippetURL(url)
  if (isFileURL(url)) return parseFileURL(url)
  throw new Error('Invalid URL: not a GitHub URL')
}

// matches link to file snippet w/ line numbers
// ex. https://github.com/omni-network/omni/blob/0593036/contracts/src/protocol/OmniPortal.sol#L135-L151
const ghSnippetRegex =
  /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)#L(\d+)-L(\d+)/

// matches link to a whole file
// ex. https://github.com/omni-network/omni/blob/0593036/contracts/src/protocol/OmniPortal.sol
const ghFileRegex =
  /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)/

const isSnippetURL = (url: string) => ghSnippetRegex.test(url)
const isFileURL = (url: string) => ghFileRegex.test(url)

function parseSnippetURL(url: string): GithubFile {
  const match = url.match(ghSnippetRegex)

  const [, , owner, repo, ref, path, startLineS, endLineS] = match

  const startLine = parseInt(startLineS)
  const endLine = parseInt(endLineS)

  if (isNaN(startLine) || isNaN(endLine)) {
    throw new Error(
      `Invalid Permalink: non-numeric line numbers ${startLineS} ${endLineS}`,
    )
  }

  return { isSnippet: true, owner, repo, ref, path, startLine, endLine }
}

function parseFileURL(url: string): GithubFile {
  const match = url.match(ghFileRegex)
  const [, , owner, repo, ref, path] = match
  return { isSnippet: false, owner, repo, ref, path }
}

function determineLanguage(filePath: string) {
  const extension = filePath.slice(filePath.lastIndexOf('.'))
  const languageMap = {
    '.js': 'javascript',
    '.ts': 'typescript',
    '.py': 'python',
    '.cpp': 'cpp',
    '.c': 'c',
    '.java': 'java',
    '.rs': 'rust',
    '.html': 'html',
    '.css': 'css',
    '.md': 'markdown',
    '.sh': 'bash',
    '.sol': 'solidity',
    '.go': 'go',
    '.json': 'json',
    '.yml': 'yaml',
    '.yaml': 'yaml',
    '.toml': 'toml',
    '.xml': 'xml',
    '.sql': 'sql',
  }
  return languageMap[extension] || 'plaintext'
}

export function getNumberOfLines(url: string) : number {
  const file = parseURL(url)
  return file.isSnippet ? file.endLine - file.startLine : 0
}
