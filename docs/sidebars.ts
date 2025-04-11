import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// Define the sidebar array following the new structure
const solvernetSidebarItems = [
    {
      type: 'category',
      label: 'Learn', // General Omni Info
      collapsed: true,
      items: [
        'index', // Welcome
        'introduction/omni-overview',
        'introduction/what-is-solvernet', // Can maybe merge into overview later
        'introduction/whitepaper',
        'introduction/omni-token',
      ],
    },
    {
      type: 'category',
      label: 'SolverNet Concepts',
      collapsed: true,
      items: [
        {
          type: 'doc',
          id: 'concepts/the-problem',
          label: 'Problem: Fragmentation & UX'
        },
        {
          type: 'doc',
          id: 'concepts/the-solution',
          label: 'Solution: Intents & SolverNet'
        },
        'concepts/intent-mechanism',
        'concepts/single-chain-deployment',
      ],
    },
    {
      type: 'category',
      label: 'SolverNet SDK',
      collapsed: true,
      items: [
        'sdk/getting-started',
        'sdk/hooks/useQuote',
        'sdk/hooks/useOrder',
        'sdk/utils/withExecAndTransfer',
        'sdk/supported-assets',
      ],
    },
    {
      type: 'category',
      label: 'SolverNet Guides',
      collapsed: true,
      items: [
        'guides/basic-deposit',
        'guides/contracts-without-onbehalfof',
        {
          type: 'category',
          label: 'Demos',
          collapsed: true,
          items: [
            'guides/demos/symbiotic',
            'guides/demos/eigen-layer',
            'guides/demos/rocketpool',
          ],
        },
      ],
    },
    {
      type: 'doc',
      id: 'resources',
      label: 'Resources',
    }
  ];

// Export an object containing the named sidebar (type inferred)
const sidebars = {
  solvernetSidebar: solvernetSidebarItems,
};

export default sidebars;
