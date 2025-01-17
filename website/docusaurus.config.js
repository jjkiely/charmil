const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'Areogear-Charmil',
  tagline: 'Charmil provides ecosystem of tools to build production ready command line tools with Cobra Framework.',
  url: 'https://aerogear.github.io',
  baseUrl: '/charmil/',
  onBrokenLinks: 'ignore',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/logo.jpg',
  organizationName: 'aerogear',
  projectName: 'charmil',
  themeConfig: {
    navbar: {
      items: [
        {
          type: 'doc',
          docId: 'intro',
          position: 'left',
          label: 'Documentation',
        },
        {
          href: 'https://github.com/aerogear/charmil',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Documentation',
              to: 'docs/',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Discord',
              href: 'https://discord.com/invite/nAQBYZncvm',
            },
          ],
        },
        {
          title: 'More',
          items: [
            
            {
              label: 'GitHub',
              href: 'https://github.com/aerogear/charmil',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Aerogear Charmil.`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/aerogear/charmil',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
