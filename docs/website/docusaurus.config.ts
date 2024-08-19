import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';
import remarkMath from 'remark-math';
import rehypeKatex from 'rehype-katex';

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
    navbar: {
      title: 'Omni Developers',
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
      links: [
        {
          title: 'TECH',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/facebook/docusaurus',
            },
            {
              label: 'Status page',
              href: 'https://status.omni.network/',
            },
          ],
        },
        {
          title: 'COMMUNITY',
          items: [
            {
              label: 'Telegram for developers',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/docusaurus',
            },
            {
              label: 'Discord',
              href: 'https://discordapp.com/invite/docusaurus',
            },
          ],
        },
        {
          title: 'JOIN US',
          items: [
            {
              label: 'Career',
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
