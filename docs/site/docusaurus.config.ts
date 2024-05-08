import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";
import remarkMath from 'remark-math';
import rehypeKatex from 'rehype-katex';

const config: Config = {
  title: "Omni Docs",
  tagline: "Documentation for the Omni Network",
  favicon: "img/favicon.svg",

  // Set the production url of your site here
  url: "https://docs.omni.network/",
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "throw",

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  clientModules: ["src/client-modules/index.ts"],

  presets: [
    [
      "classic",
      {
        docs: {
          routeBasePath: "/",
          sidebarPath: "./sidebars.ts",
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl: "https://github.com/omni-network/omni/tree/main/docs/site",
          remarkPlugins: [remarkMath],
          rehypePlugins: [rehypeKatex],
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
        gtag: {
          trackingID: 'G-P82Q1TL0SX',
        },
      } satisfies Preset.Options,
    ],
  ],
  stylesheets: [
    {
      href: 'https://cdn.jsdelivr.net/npm/katex@0.13.24/dist/katex.min.css',
      type: 'text/css',
      integrity: 'sha384-odtC+0UGzzFL/6PNoE8rX/SPcQDXBJ+uRepguP4QkPCm2LBxH3FA3y+fKSiJ+AmM', // pragma: allowlist secret
      crossorigin: 'anonymous',
    },
  ],

  themeConfig: {
    // Replace with your project's social card
    image: "img/omni-banner.png",
    navbar: {
      title: "Omni Docs",
      logo: {
        alt: "Omni Logo",
        src: "img/logo.svg",
      },
      items: [
        {
          position: "left",
          label: "Learn",
          to: "/learn/introduction",
        },
        {
          position: "left",
          label: "Protocol",
          to: "/protocol/introduction",
        },
        {
          position: "left",
          label: "Octane",
          to: "/octane/background/introduction",
        },
        {
          position: "left",
          label: "Develop",
          to: "/develop/introduction",
        },
        {
          position: "left",
          label: "Operate",
          to: "/operate/introduction",
        },
        {
          position: "right",
          label: "Tools",
          to: "/tools",
        },
        {
          href: "https://github.com/omni-network/omni",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      // style: "dark",
      links: [
        {
          label: "Main Site",
          href: "https://omni.network",
        },
        {
          label: "Discord",
          href: "https://discord.gg/bKNXmaX9VD",
        },
        {
          label: "Twitter",
          href: "https://twitter.com/OmniFDN",
        },
        {
          label: "Telegram",
          href: "https://t.me/omnifdn",
        },
        {
          label: "GitHub",
          href: "https://github.com/omni-network/omni/",
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} The Omni Network`,
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
