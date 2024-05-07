import { useColorMode } from '@docusaurus/theme-common';
import CodeBlock from "@theme/CodeBlock";
import { useEffect, useState } from 'react';
import useGitHubCodeBlock from './useGitHubCodeBlock';

import './GitHubCodeBlock.css';

const GitHubCodeBlock = ({ repoUrl }) => {
  const [data, setData] = useState(null);
  const colorMode = useColorMode();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const result = await useGitHubCodeBlock(repoUrl);
        setData(result);
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, [repoUrl]);

  if (!data) {
    return (
      <div className="code-snippet-container">
        <CodeBlock language="plaintext" className="code-snippet-block">Fetching code from GitHub...</CodeBlock>
      </div>
    );
  }

  const { code, isCodeMatch, mainUrl, sourceUrl, language, isError } = data;

  return (
    <div className="code-snippet-container">
      <CodeBlock language={language} className="code-snippet-block">{code}</CodeBlock>

      {!isCodeMatch &&
        <div className="code-snippet-warning">
          <a href={mainUrl}>
            <strong>Warning: </strong>Code shown does not match the main branch. Please visit the repository URL for actual code.
          </a>
        </div>
      }

      {!isError &&
        <div className="code-snippet-footer">
          <a href={sourceUrl}>
            See source on GitHub <img src={colorMode.colorMode === 'dark' ? "/img/github-icon-light.svg" : "/img/github-icon-dark.svg"} alt="GitHub" />
          </a>
        </div>
      }
    </div>
  );
}

export default GitHubCodeBlock;
