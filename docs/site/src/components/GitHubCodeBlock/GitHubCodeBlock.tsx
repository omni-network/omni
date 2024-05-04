import React, { useEffect, useState } from 'react';
import axios from 'axios';
import CodeBlock from "@theme/CodeBlock";
import { useColorMode } from '@docusaurus/theme-common';

import './GitHubCodeBlock.css';

const GitHubCodeBlock = ({ repoUrl }) => {
  const [code, setCode] = useState('');
  const [language, setLanguage] = useState('plaintext');
  const [sourceUrl, setSourceUrl] = useState('');
  const [mainUrl, setMainUrl] = useState('');
  const [isCodeMatch, setCodeMatch] = useState(true);

  useEffect(() => {
    fetchCode();
  }, [repoUrl]);

  async function fetchCode() {
    if (!repoUrl) {
      console.error("No URL provided.");
      setCode("Error: No URL provided.");
      return;
    }

    var match = matchUrl(repoUrl);
    if (!match) {
      return;
    }

    const [wholeUrl, site, owner, repo, branch, filePath, startLine, endLine] = match;

    const rawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/${branch}/${filePath}` + (startLine && endLine ? `#L${startLine}-L${endLine}` : ''); // Fetch the raw content from GitHub
    const mainRawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/main/${filePath}`; // main branch raw content URL
    setSourceUrl(`https://github.com/${owner}/${repo}/blob/${branch}/${filePath}` + (startLine && endLine ? `#L${startLine}-L${endLine}` : '')); // Build the source URL to view on GitHub, including line numbers if available
    setMainUrl(`https://github.com/${owner}/${repo}/blob/main/${filePath}`); // main branch source URL

    try {
      const response = await axios.get(rawUrl); // Fetch the raw content from the specified commit hash
      const lines = extractLines(response, startLine, endLine);

      const mainResponse = await axios.get(mainRawUrl); // Fetch the raw content from the main branch
      const mainLines = extractLines(mainResponse, startLine, endLine);

      if (lines !== mainLines) { // Check if the code matches the main branch
        setCodeMatch(false);
      }

      setCode(lines);
      setLanguage(determineLanguage(filePath));
    } catch (error) {
      console.error('Error fetching code:', error.response ? error.response.data : error.message);
      setCode(`Error: ${error.response ? error.response.data : "Could not fetch file"}`);
    }
  }

  function extractLines(mainResponse, startLine: string, endLine: string) {
    const allLines = mainResponse.data.split('\n');
    const lines = (startLine && endLine) ? allLines.slice(parseInt(startLine) - 1, parseInt(endLine)).join('\n') : allLines.join('\n');
    return lines;
  }

  function matchUrl(url: string) {
    const permalinkLinesRegex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)#L(\d+)-L(\d+)/;
    const repoFileRegex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)/;
    let match = url.match(permalinkLinesRegex);
    if (match) {
      if (match[4].length !== 40) {
        console.error("Invalid URL: Permalink with lines URL should include a commit hash.");
        setCode("Error: Invalid GitHub Permalink Lines URL format.");
        return null;
      }
      return match;
    }
    match = url.match(repoFileRegex);
    if (!match) {
      console.error("Invalid URL: Check the URL format.");
      setCode("Error: Invalid GitHub URL format.");
    }
    return match;
  }

  function determineLanguage(filePath) {
    const extension = filePath.slice(filePath.lastIndexOf('.'));
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
    };
    return languageMap[extension] || 'plaintext';
  }

  function getGitHubIcon() {
    const { colorMode } = useColorMode();
    return colorMode === 'dark' ? "/img/github-icon-light.svg" : "/img/github-icon-dark.svg";
  }

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
      <div className="code-snippet-footer">
        <a href={sourceUrl}>
          See source on GitHub <img src={getGitHubIcon()} alt="GitHub" />
        </a>
      </div>
    </div>
  );
}

export default GitHubCodeBlock;
