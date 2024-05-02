import React, { useEffect, useState } from 'react';
import axios from 'axios';
import CodeBlock from "@theme/CodeBlock";
import { useColorMode } from '@docusaurus/theme-common';

import './CodeSnippet.css';

const CodeSnippet = ({ repoUrl }) => {
  const [code, setCode] = useState('');
  const [language, setLanguage] = useState('plaintext');
  const [sourceUrl, setSourceUrl] = useState('');

  useEffect(() => {
    async function fetchCode() {
      // Use a single regex that can optionally capture line numbers
      const regex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)#L(\d+)-(\d+)/;
      var match = repoUrl.match(regex);
      if (!match) {
        const noLinesRegex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)/;
        match = repoUrl.match(noLinesRegex);
      }

      if (match) {
        const [wholeUrl, site, owner, repo, branch, filePath, startLine, endLine] = match;
        // Build the source URL to view on GitHub, including line numbers if available
        setSourceUrl(`https://github.com/${owner}/${repo}/blob/${branch}/${filePath}${startLine ? `#L${startLine}-L${endLine}` : ''}`);
        // Fetch the raw content from GitHub
        const rawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/${branch}/${filePath}` + (startLine && endLine ? `#L${startLine}-L${endLine}` : '');

        try {
          const response = await axios.get(rawUrl);
          const allLines = response.data.split('\n');
          // Apply line slicing only if line numbers are provided and valid
          const lines = (startLine && endLine) ? allLines.slice(parseInt(startLine) - 1, parseInt(endLine)).join('\n') : allLines.join('\n');
          setCode(lines);
          setLanguage(determineLanguage(filePath));
        } catch (error) {
          console.error('Error fetching code:', error.response ? error.response.data : error.message);
          setCode(`Error: ${error.response ? error.response.data : "Could not fetch file"}`);
        }
      } else {
        console.error("Regex match failed. Check the URL format.");
        setCode("Error: Invalid GitHub URL format.");
      }
    }

    fetchCode();
  }, [repoUrl]);

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

  function getGithubIcon() {
    const { colorMode } = useColorMode();
    return colorMode === 'dark' ? "/img/github-icon-light.svg" : "/img/github-icon-dark.svg";
  }

  return (
    <div className="code-snippet-container">
      <CodeBlock language={language} className="code-snippet-block">{code}</CodeBlock>
      <div className="code-snippet-footer">
        <a href={sourceUrl} target="_blank" rel="noopener noreferrer">
          See source on GitHub <img src={getGithubIcon()} alt="GitHub" />
        </a>
      </div>
    </div>
  );
};

export default CodeSnippet;
