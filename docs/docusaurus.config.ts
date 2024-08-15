import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Omni',
  tagline: 'Cross-chain dapps made easy',
  favicon: 'img/favicon.ico',

  // Set the production url of your site here
  url: 'https://docs.omni.network',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'omni-network', // Usually your GitHub org/user name.
  projectName: 'omni', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/omni-network/omni/tree/main/packages/create-docusaurus/templates/shared/',
        },
        // blog: {
        //   showReadingTime: true,
        //   feedOptions: {
        //     type: ['rss', 'atom'],
        //     xslt: true,
        //   },
        //   // Please change this to your repo.
        //   // Remove this to remove the "edit this page" links.
        //   editUrl:
        //     'https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/',
        // },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: 'img/docusaurus-social-card.jpg',
    navbar: {
      title: 'Omni Builders',
      logo: {
        alt: 'Omni Logo',
        src: 'img/logo.svg',
      },
      items: [
        // {
        //   type: 'docSidebar',
        //   sidebarId: 'tutorialSidebar',
        //   position: 'left',
        //   label: 'Build',
        // },
        // {to: '/blog', label: 'Blog', position: 'left'},
        {
          href: 'https://github.com/facebook/docusaurus',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Code',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/facebook/docusaurus',
            },
            // {
            //   label: 'Build',
            //   to: '/docs/build',
            // },
            // {
            //   label: 'Run a node',
            //   to: '/docs/intro',
            // },
            // {
            //   label: '$OMNI token',
            //   to: '/docs/intro',
            // },
            // {
            //   label: 'FAQ',
            //   to: '/docs/intro',
            // },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Telegram for devs',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Discord',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/docusaurus',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Whitepaper',
              to: 'https://docs.omni.network/whitepaper.pdf',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/facebook/docusaurus',
            },
          ],
        },
      ],
      // copyright: `Copyright © ${new Date().getFullYear()}`,
    },
    prism: {
      theme: prismThemes.vsLight,
      darkTheme: prismThemes.vsDark,
      additionalLanguages: ["solidity", "bash"],
    },    
    algolia: {
      // The application ID provided by Algolia
      appId: "YGLZ6VW4T5", // pragma: allowlist secret
      // Public API key: it is safe to commit it
      apiKey: "64557e587da746830ff903f126eb134b", // pragma: allowlist secret
      indexName: "omni",
      contextualSearch: false,
      searchParameters: {
        clickAnalytics: true,
        analytics: true,
        enableReRanking: true,
      },
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
