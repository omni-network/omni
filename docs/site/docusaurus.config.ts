import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";

const config: Config = {
  title: "Omni Docs",
  tagline: "Documentation for the Omni Network",
  favicon: "img/favicon-dark.png",

  // Set the production url of your site here
  url: "https://your-docusaurus-site.example.com",
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "omni-network", // Usually your GitHub org/user name.
  projectName: "docs", // Usually your repo name.

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl: "https://github.com/omni-network/omni/docs/",
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: "img/docusaurus-social-card.jpg",
    navbar: {
      title: "My Site",
      logo: {
        alt: "My Site Logo",
        src: "img/logo.svg",
      },
      items: [
        {
          to: "docs/home",
          sidebarId: "oldSidebar",
          position: "left",
          label: "Omni",
        },
        // { to: "/blog", label: "Blog", position: "left" },
        {
          href: "https://github.com/omni-network/omni",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      style: "dark",
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
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ["solidity"],
    },
    algolia: {
      appId: "<NEW_APP_ID>", // pragma: allowlist secret
      apiKey: "<NEW_API_KEY>", // pragma: allowlist secret
      indexName: "index-name",
      contextualSearch: true,
      searchParameters: {
        clickAnalytics: true,
        analytics: true,
        enableReRanking: true,
      },
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
