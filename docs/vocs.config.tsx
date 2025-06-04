import { defineConfig } from 'vocs';

export default defineConfig({
  title: 'Omni Devs | Docs',
  rootDir: './docs/',
  baseUrl: '/',
  description: 'Documentation for the Omni SDK, SolverNet, and related concepts.',
  logoUrl: '/img/logo.svg',
  iconUrl: '/img/favicon.svg',
  ogImageUrl: 'https://docs.omni.network/img/omni-banner.png',
  editLink: {
    pattern: 'https://github.com/omni-network/omni/tree/main/docs/docs/pages/:path',
    text: 'Suggest changes to this page on GitHub',
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
      items: [
        { text: 'Welcome', link: '/' },
        { text: 'Omni Overview', link: '/introduction/omni-overview' },
        { text: 'What is SolverNet?', link: '/introduction/what-is-solvernet' },
        { text: 'Whitepaper', link: '/introduction/whitepaper' },
        { text: 'Omni Token', link: '/introduction/omni-token' },
      ],
    },
    {
      text: 'SolverNet',
      items: [
        { text: 'The Problem', link: '/concepts/the-problem' },
        { text: 'The Solution', link: '/concepts/the-solution' },
        { text: 'Intents Mechanism', link: '/concepts/intent-mechanism' },
        { text: 'Single Chain Deployments', link: '/concepts/single-chain-deployment' },
      ],
    },
    {
      text: 'SolverNet SDK',
      items: [
        {
          text: 'Getting Started',
          link: '/sdk/getting-started',
          collapsed: false,
          items: [
            { text: 'With React', link: '/sdk/getting-started/react' },
            { text: 'Without framework', link: '/sdk/getting-started/core' },
            { text: 'Basic Deposit', link: '/sdk/getting-started/basic-deposit' },
          ],
        },
        { text: 'Supported Assets', link: '/sdk/supported-assets' },
        {
          text: 'React hooks',
          collapsed: true,
          items: [
            { text: 'useQuote', link: '/sdk/hooks/useQuote' },
            { text: 'useOrder', link: '/sdk/hooks/useOrder' },
            { text: 'useOmniContracts', link: '/sdk/hooks/useOmniContracts' },
            { text: 'useOmniAssets', link: '/sdk/hooks/useOmniAssets' },
          ],
        },
        {
          text: 'Core functions',
          collapsed: true,
          items: [
            {
              text: 'getContracts',
              link: '/sdk/core/getContracts',
            },
            {
              text: 'getQuote',
              link: '/sdk/core/getQuote',
            },
            {
              text: 'validateOrder',
              link: '/sdk/core/validateOrder',
            },
            {
              text: 'openOrder',
              link: '/sdk/core/openOrder',
            },
            {
              text: 'generateOrder',
              link: '/sdk/core/generateOrder',
            },
            {
              text: 'waitForOrderOpen',
              link: '/sdk/core/waitForOrderOpen',
            },
            {
              text: 'waitForOrderClose',
              link: '/sdk/core/waitForOrderClose',
            },
            {
              text: 'watchDidFill',
              link: '/sdk/core/watchDidFill',
            },
          ],
        },
        {
          text: 'Utility functions',
          items: [
            {
              text: 'withExecAndTransfer',
              link: '/sdk/utils/withExecAndTransfer',
            },
          ],
        },
        { text: 'Swaps', link: '/sdk/swaps' },
        {
          text: 'Advanced Guides',
          collapsed: true,
          items: [
            { text: 'Cross Chain Transfers', link: '/guides/transfers' },
            { text: 'Multi-Step Deposit', link: '/guides/multistep-deposit' },
            { text: 'Contracts without onBehalfOf', link: '/guides/contracts-without-onbehalfof' },
            { text: 'Debugging destination calls', link: '/guides/debugging-destination-calls' },
            { text: 'Symbiotic', link: '/guides/examples/symbiotic' },
            { text: 'Rocketpool', link: '/guides/examples/rocketpool' },
          ],
        },
        { text: 'FAQ', link: '/sdk/faq' },
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
