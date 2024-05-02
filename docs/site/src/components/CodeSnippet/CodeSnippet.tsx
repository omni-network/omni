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
      const match = repoUrl.match(/github\.com\/([^/]+)\/([^/]+)\/blob\/([^/]+)\/([^#]+)(#L(\d+)-L(\d+))?/);
      if (match) {
        const [, owner, repo, branch, filePath, , startLine, endLine] = match;
        setSourceUrl(`https://github.com/${owner}/${repo}/blob/${branch}/${filePath}`);
        const rawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/${branch}/${filePath}`;
        try {
          const response = await axios.get(rawUrl);
          const allLines = response.data.split('\n');
          const lines = startLine && endLine ? allLines.slice(startLine - 1, endLine).join('\n') : allLines.join('\n');
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
    return languageMap[extension] || 'plaintext'; // Default to plaintext if no match found
  }

  function getGithubIcon() {
    // determine which svg string to get based on the theme
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
