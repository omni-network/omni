import { defineConfig } from 'vocs';

export default defineConfig({
  title: 'Omni Devs | Docs',
  rootDir: './docs/',
  baseUrl: '/',
  description: 'Documentation for the Omni Network SolverNet SDK and concepts.',
  logoUrl: '/img/logo.svg',
  iconUrl: '/img/favicon.svg',
  editLink: {
    pattern: 'https://github.com/omni-network/omni/tree/main/docs/docs/pages/:path',
    text: 'Edit this page on GitHub',
  },
  socials: [
    {
      icon: 'github',
      link: 'https://github.com/omni-network/omni',
    },
    {
      icon: 'x',
      link: 'https://twitter.com/OmniFDN',
    },
    {
      icon: 'discord',
      link: 'https://discord.com/invite/bKNXmaX9VD',
    },
  ],
  sidebar: [
    {
      text: 'Learn',
      collapsed: true,
      items: [
        { text: 'Welcome', link: '/' },
        { text: 'Omni Overview', link: '/introduction/omni-overview' },
        { text: 'What is SolverNet?', link: '/introduction/what-is-solvernet' },
        { text: 'Whitepaper', link: '/introduction/whitepaper' },
        { text: 'Omni Token', link: '/introduction/omni-token' },
      ],
    },
    {
      text: 'SolverNet Concepts',
      collapsed: true,
      items: [
        { text: 'Problem: Fragmentation & UX', link: '/concepts/the-problem' },
        { text: 'Solution: Intents & SolverNet', link: '/concepts/the-solution' },
        { text: 'Intent Mechanism', link: '/concepts/intent-mechanism' },
        { text: 'Single Chain Deployment', link: '/concepts/single-chain-deployment' },
      ],
    },
    {
      text: 'SolverNet SDK',
      collapsed: true,
      items: [
        { text: 'Getting Started', link: '/sdk/getting-started' },
        { text: 'useQuote', link: '/sdk/hooks/useQuote' },
        { text: 'useOrder', link: '/sdk/hooks/useOrder' },
        { text: 'withExecAndTransfer', link: '/sdk/utils/withExecAndTransfer' },
        { text: 'Supported Assets', link: '/sdk/supported-assets' },
      ],
    },
    {
      text: 'SolverNet Guides',
      collapsed: true,
      items: [
        { text: 'Basic Deposit', link: '/guides/basic-deposit' },
        { text: 'Contracts without onBehalfOf', link: '/guides/contracts-without-onbehalfof' },
        {
          text: 'Examples',
          collapsed: true,
          items: [
            { text: 'Symbiotic', link: '/guides/examples/symbiotic' },
            { text: 'EigenLayer', link: '/guides/examples/eigenlayer' },
            { text: 'Rocketpool', link: '/guides/examples/rocketpool' },
          ],
        },
      ],
    },
    { text: 'Resources', link: '/resources' },
  ],
  topNav: [
    {
      text: 'Build with us',
      link: 'https://tally.so/r/wAJ2EB',
    },
  ],
});
