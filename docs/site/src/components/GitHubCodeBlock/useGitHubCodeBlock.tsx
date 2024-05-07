export default async function useGitHubCodeBlock(repoUrl: string) {
    let match = matchUrl(repoUrl);
    if (match[0] === "Error") {
        return {
            code: match[1],
            isCodeMatch: true,
            mainUrl: "",
            sourceUrl: "",
            language: "plaintext",
            isError: true,
        };
    }

    const [, , owner, repo, branch, filePath, startLine, endLine] = match;
    const rawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/${branch}/${filePath}` + (startLine && endLine ? `#L${startLine}-L${endLine}` : '');
    const mainRawUrl = `https://raw.githubusercontent.com/${owner}/${repo}/main/${filePath}`;

    const linesRes = await fetchCodeLinesFromUrl(rawUrl, startLine, endLine);
    const mainLinesRes = await fetchCodeLinesFromUrl(mainRawUrl, startLine, endLine);

    const isCodeMatch = linesRes === mainLinesRes;
    const sourceUrl = `https://github.com/${owner}/${repo}/blob/${branch}/${filePath}` + (startLine && endLine ? `#L${startLine}-L${endLine}` : '');
    const mainUrl = `https://github.com/${owner}/${repo}/blob/main/${filePath}`;

    return {
        code: linesRes,
        isCodeMatch,
        mainUrl,
        sourceUrl,
        language: determineLanguage(filePath),
        isError: false,
    };
}

async function fetchCodeLinesFromUrl(url: string, startLine: string, endLine: string) {
    const response = await fetch(url);
    const data = await response.text();
    const lines = extractLines(data, startLine, endLine);
    return lines;
}

function extractLines(mainResponse: string, startLine: string, endLine: string) {
    const allLines = mainResponse.split('\n');
    const lines = (startLine && endLine) ? allLines.slice(parseInt(startLine) - 1, parseInt(endLine)).join('\n') : allLines.join('\n');
    return lines;
}

function matchUrl(url: string) {
    const permalinkLinesRegex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)#L(\d+)-L(\d+)/;
    let match = url.match(permalinkLinesRegex);
    if (match) {
        if (match[4].length !== 40) {
            console.error("Invalid URL: Permalink with lines URL should include a commit hash.");
            let err = "Error: Invalid GitHub URL with commit hash and lines format.";
            return ["Error", err];
        }
        return match;
    }
    const repoFileRegex = /(https:\/\/github\.com\/)([^\/]+)\/([^\/]+)\/blob\/([^\/]+)\/([^#]+)/;
    match = url.match(repoFileRegex);
    if (!match) {
        console.error("Invalid URL: Check the URL format.");
        let err = "Error: Invalid GitHub URL format.";
        return ["Error", err];
    }
    return match;
}

function determineLanguage(filePath: string) {
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
