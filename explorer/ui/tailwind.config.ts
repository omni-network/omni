import type { Config } from 'tailwindcss'

export default {
  content: ['./app/**/*.{html,js,ts,jsx,tsx}', '~/components/**/*.{html,js,ts,tsx,jsx}'],
  theme: {
    extend: {
      colors: {
        'default': 'var(--color-text)',
        'subtle': 'var(--color-text-subtle)',
        'subtlest': 'var(--color-text-subtlest)',
        'inverse': 'var(--color-text-inverse)',
        'primary': 'var(--color-text-primary)',
        'critical': 'var(--color-text-critical)',
        'moderate': 'var(--color-text-moderate)',
        'positive': 'var(--color-text-positive)',
        'disabled': 'var(--color-text-disabled)',
        'white': 'var(--color-text-white)',

        'link': 'var(--color-link)',
        'link-hover': 'var(--color-link-hover)',
        'link-pressed': 'var(--color-link-pressed)',
        'link-visited': 'var(--color-link-visited)',

        'icon-positive': 'var(--color-icon-primary)',
        'icon-critical': 'var(--color-icon-critical)',
        'icon-moderate': 'var(--color-icon-moderate)',

        'border-default': 'var(--color-border-default)',
        'border-subtle': 'var(--color-border-subtle)',
        'border-disabled': 'var(--color-border-disabled)',

        'surface': 'var(--color-surface)',
        'raised': 'var(--color-raised)',
        'overlay':'var(--color-overlay)',

        'bg-positive': 'var(--color-bg-positive)',
        'bg-critical': 'var(--color-bg-critical)',
        'bg-subtle': 'var(--color-bg-subtle)',
        'bg-moderate': 'var(--color-bg-moderate)',

        'bg-interactive-default': 'var(--color-bg-interactive-default)',
        'bg-interactive-hover': 'var(--color-bg-interactive-hover)',
        'bg-interactive-active': 'var(--color-bg-interactive-active)',

        'bg-hover': 'var(--color-bg-hover)',
        'bg-active': 'var(--color-bg-active)',

        'bg-input-default': 'var(--color-bg-input-default)',
        'bg-input-hover': 'var(--color-bg-input-hover)',
        'bg-input-active': 'var(--color-bg-input-active)',

        'bg-search-default': 'var(--color-bg-search-default)',
        'bg-search-hover': 'var(--color-bg-search-hover)',
        'bg-search-active': 'var(--color-bg-search-active)',
      },
    },
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
  daisyui: {
    themes: [
      {
        light: {
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          ...require('daisyui/src/theming/themes')['light'],
          fontFamily: 'Manrope',
          '--color-text': '#142442',
          '--color-text-subtle': '#394866',
          '--color-text-subtlest': '#69748B',
          '--color-text-inverse': '#F7F7F7',
          '--color-text-primary': '#1B52B8',
          '--color-text-critical': '#BB372B',
          '--color-text-positive': '#2E7B41',
          '--color-text-moderate': '#80720B',
          '--color-text-disabled': '#B0B7C3',
          '--color-text-white': '#FDFDFD',
          '--color-link': '#0564FF',
          '--color-link-hover': '#1B52B8',
          '--color-link-pressed': '#28467C',
          '--color-link-visited': '#7C48BD',

          '--color-icon-positive': '#2E7B41',
          '--color-icon-critical': '#BB372B',
          '--color-icon-moderate': '#CCB612',

          '--color-border-default': '#DCDFE3',
          '--color-border-subtle': '#EEEFF0',
          '--color-border-disabled': '#B0B7C3',

          '--color-surface':'#F7F7F7',
          '--color-raised':'#FDFDFD',
          '--color-overlay': '#FEFEFE',

          '--color-bg-positive': '#2E7B4126',
          '--color-bg-critical': '#BB372B26',
          '--color-bg-subtle':'#F7F7F7',
          '--color-bg-moderate':'#FFE31624',
          '--color-bg-hover': '#1023460D',
          '--color-bg-active': '#1023460D',

          '--color-bg-interactive-default': '#1023460D',
          '--color-bg-interactive-hover': '#1023460D',
          '--color-bg-interactive-active': '#1023460D',

          '--color-bg-input-default': '#1023460D',
          '--color-bg-input-hover': '#1023460D',
          '--color-bg-input-active': '#1023460D',

          '--color-bg-search-default': '#EEEFF0',
          '--color-bg-search-hover': '#1023460D',
          '--color-bg-search-active': '#DCDFE3',
        },
        dark: {
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          ...require('daisyui/src/theming/themes')['dark'],
          fontFamily: 'Manrope',
          '--color-text': '#F7F7F7',
          '--color-text-subtle': '#EEEFF0',
          '--color-text-subtlest': '#B0B7C3',
          '--color-text-inverse': '#0B1A37',
          '--color-text-primary': '#7CAEFF',
          '--color-text-critical': '#F6D5D5',
          '--color-text-positive': '#C7E6D5',
          '--color-text-moderate': '#FFFCE1',
          '--color-text-disabled': '#B0B7C3',
          '--color-text-white': '#FDFDFD',
          '--color-link': '#7CAEFF',
          '--color-link-hover': '#5588E4',
          '--color-link-pressed': '#5588E4',
          '--color-link-visited': '#C29EEF',

          '--color-icon-positive': '#5588E4',
          '--color-icon-critical': '#F6D5D5',
          '--color-icon-moderate': '#CCB612',

          '--color-border-default': '#69748B',
          '--color-border-subtle': '#394866',
          '--color-border-disabled': '#B0B7C3',

          '--color-surface':'#142442',
          '--color-raised':'#22314C',
          '--color-overlay': '#394866',

          '--color-bg-positive': '#2E7B4166',
          '--color-bg-critical': '#BB372B66',
          '--color-bg-subtle':'#FBFBFC1F',
          '--color-bg-moderate':'#FFE31633',
          '--color-bg-hover': '#FBFBFC0D',
          '--color-bg-active': '#FBFBFC0D',

          '--color-bg-input-default': '#FBFBFC0D',
          '--color-bg-input-hover': '#FBFBFC14',
          '--color-bg-input-active': '#394866',

          '--color-bg-search-default': '#FBFBFC0D',
          '--color-bg-search-hover': '#1023460D',
          '--color-bg-search-active': '#DCDFE3',

        },
      },
    ],
    // base: true, // applies background color and foreground color for root element by default
    // styled: true, // include daisyUI colors and design decisions for all components
    // utils: true, // adds responsive and modifier utility classes
    // prefix: '', // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
    // logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
    // themeRoot: ':root', // The element that receives theme color CSS variables
    // rtl: false, // Enable RTL support
  },
} satisfies Config
