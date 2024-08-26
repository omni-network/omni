import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';
import remarkMath from 'remark-math';
import rehypeKatex from 'rehype-katex';

const config: Config = {
  title: 'Omni Devs | Docs',
  tagline: 'Cross-chain dapps made easy',
  favicon: 'img/favicon.svg',

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
  onBrokenMarkdownLinks: 'throw',

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
          path: "../content",
          routeBasePath: "/",
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/omni-network/omni/tree/main/docs/website',
            remarkPlugins: [remarkMath],
            rehypePlugins: [rehypeKatex],
        },
        theme: {
          customCss: './src/css/custom.css',
        },
        gtag: {
          trackingID: 'G-P82Q1TL0SX',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: "img/omni-banner.png",
    docs: {
      sidebar: {
        autoCollapseCategories: true,
      },
    },
    navbar: {
      title: 'Omni Developers',
      logo: {
        alt: 'Omni Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          href: 'https://tally.so/r/wAJ2EB',
          label: 'Build with us',
          position: 'right',
        },
        {
          href: 'https://github.com/omni-network/omni',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      links: [
        {
          title: 'TECH',
          items: [
            {
              label: 'GitHub',
              to: 'https://github.com/omni-network/omni',
            },
            {
              label: 'Explorer',
              to: 'https://omega.omniscan.network/',
            },
            {
              label: 'Status page',
              to: 'https://status.omni.network/',
            },
          ],
        },
        {
          title: 'COMMUNITY',
          items: [
            {
              label: 'Telegram for developers',
              to: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              to: 'https://twitter.com/docusaurus',
            },
            {
              label: 'Discord',
              to: 'https://discordapp.com/invite/docusaurus',
            },
          ],
        },
        {
          title: 'JOIN US',
          items: [
            {
              label: 'Careers',
              to: 'https://boards.greenhouse.io/omnifdn',
            },
          ],
        },
        {
          title: 'LEGAL',
          items: [
            {
              label: 'Terms of Service',
              to: 'https://docs.omni.network/tos.pdf',
            },
            {
              label: 'Privacy Policy',
              to: 'https://docs.omni.network/privacy-policy.pdf',
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
