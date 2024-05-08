import { useColorMode } from '@docusaurus/theme-common'
import CodeBlock from '@theme/CodeBlock'
import { getNumberOfLines, useCodeBlock } from './useCodeBlock'

import './GitHubCodeBlock.css'

const GitHubCodeBlock = ({ url }: { url: string }) => {
  const { data, error, isLoading, isError } = useCodeBlock({ url })

  if (isLoading) return <CodeBlockLoading numLines={getNumberOfLines(url)}/>
  if (isError) return <CodeBlockError url={url} error={error} />

  const { code, matchesMain, mainURL, sourceURL, language } = data

  return (
    <div className="code-snippet-container">
      <CodeBlock language={language} className="code-snippet-block">
        {code}
      </CodeBlock>

      {!matchesMain && (
        <div className="code-snippet-warning">
          <a href={mainURL}>
            <strong>Warning: </strong>Code shown does not match the main branch.
            Please visit the repository URL for actual code.
          </a>
        </div>
      )}

      <div className="code-snippet-footer">
        <a href={sourceURL}>See source on GitHub </a>
        <GithubIcon />
      </div>
    </div>
  )
}

const GithubIcon = () => {
  const { colorMode } = useColorMode()
  return (
    <img
      src={
        colorMode === 'dark'
          ? '/img/github-icon-light.svg'
          : '/img/github-icon-dark.svg'
      }
      alt="GitHub"
    />
  )
}

const CodeBlockLoading = ({ numLines }: { numLines: number }) => {
  const blankLines = Array.from({ length: numLines }, (_, index) => (
    <div key={index} className="loading-line" />
  ));

  return (
    <div className="code-snippet-container">
      <CodeBlock language="plaintext" className="code-snippet-block">
        Fetching code from GitHub...
        {blankLines}
      </CodeBlock>
    </div>
  );
};

const CodeBlockError = ({ url, error }: { url: string; error: Error }) => {
  console.error('CodeBlockError:', error)
  return (
    <div className="code-snippet-container">
      <CodeBlock language="plaintext" className="code-snippet-block">
        Oops :( Something went wrong.
      </CodeBlock>
      <div className="code-snippet-footer">
        <a href={url}>Check source on GitHub </a>
        <GithubIcon />
      </div>
    </div>
  )
}

export default GitHubCodeBlock
