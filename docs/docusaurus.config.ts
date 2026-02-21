import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import { themes as prismThemes } from "prism-react-renderer";

const config: Config = {
  title: "S3 Controller",
  tagline: "A backend-agnostic Kubernetes controller for S3 resources",

  future: {
    v4: true,
  },

  url: "https://benfiola.github.io",
  baseUrl: "/s3-controller/",

  organizationName: "benfiola",
  projectName: "s3-controller",
  trailingSlash: false,

  onBrokenLinks: "throw",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        blog: false,
        docs: {
          routeBasePath: "/",
          sidebarPath: "./sidebars.ts",
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    colorMode: {
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: "S3 Controller",
      items: [
        {
          href: "https://github.com/benfiola/s3-controller/tree/main",
          label: "Code",
          position: "right",
        },
      ],
      logo: {
        src: "/img/icon.png",
      },
    },
    footer: {},
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ["bash"],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
