import type { ClientModule } from '@docusaurus/types';
import ExecutionEnvironment from '@docusaurus/ExecutionEnvironment';

if (ExecutionEnvironment.canUseDOM) {
  const COOKBOOK_API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NmJkMmUwZDZkMjk4YjBkZjY4OWI0ODgiLCJpYXQiOjE3MjM2NzQxMjUsImV4cCI6MjAzOTI1MDEyNX0.C--T1GtbF4lTg0402i34XQRKiTAKH6R-pruCsUPCsfg";

  window.addEventListener('load', () => {
    let element = document.getElementById('__cookbook');
    if (!element) {
      element = document.createElement('div');
      element.id = '__cookbook';
      element.dataset.apiKey = COOKBOOK_API_KEY;
      document.body.appendChild(element);
    }

    let script: HTMLScriptElement | null = document.getElementById('__cookbook-script') as HTMLScriptElement;
    if (!script) {
      script = document.createElement('script');
      script.src = 'https://cdn.jsdelivr.net/npm/@cookbookdev/docsbot/dist/standalone/index.cjs.js';
      script.id = '__cookbook-script';
      script.defer = true;
      document.body.appendChild(script);
    }
  });
}

const askCookbookModule: ClientModule = {};

export default askCookbookModule;
