import type { Config } from 'tailwindcss'

export default {
  content: ['./app/**/*.{html,js,ts,jsx,tsx}', '~/components/**/*.{html,js,ts,tsx,jsx}'],
  theme: {
    extend: {
      colors: {
        'color-text': 'var(--color-text)',
        'color-text-subtle': 'var(--color-text-subtle)',
        'color-text-subtlest': 'var(--color-text-subtlest)',
        'color-text-inverse': 'var(--color-text-inverse)',
        'color-text-primary': 'var(--color-text-primary)',
        'color-text-critical': 'var(--color-text-critical)',
        'color-text-positive': 'var(--color-text-positive)',
        'color-text-disabled': 'var(--color-text-disabled)',
        'color-text-white': 'var(--color-text-white)',

        'color-link': 'var(--color-link)',
        'color-link-hover': 'var(--color-link-hover)',
        'color-link-pressed': 'var(--color-link-pressed)',
        'color-link-visited': 'var(--color-link-visited)',

        'color-icon-primary': 'var(--color-icon-primary)',
        'color-icon-critical': 'var(--color-icon-critical)',

        'color-border': 'var(--color-border)',
        'color-border-subtle': 'var(--color-border-subtle)',
        'color-border-disabled': 'var(--color-border-disabled)',
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
          '--color-text-disabled': '#B0B7C3',
          '--color-text-white': '#FDFDFD',
          '--color-link': '#0564FF',
          '--color-link-hover': '#1B52B8',
          '--color-link-pressed': '#28467C',
          '--color-link-visited': '#7C48BD',

          '--color-icon-primary': '#2E7B41',
          '--color-icon-critical': '#BB372B',

          '--color-border': '#DCDFE3',
          '--color-border-subtle': '#EEEFF0',
          '--color-border-disabled': '#B0B7C3',
        },
        dark: {
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          ...require('daisyui/src/theming/themes')['dark'],
          fontFamily: 'Manrope',
          '--color-text': '#142442',
          '--color-text-subtle': '#EEEFF0',
          '--color-text-subtlest': '#B0B7C3',
          '--color-text-inverse': '#0B1A37',
          '--color-text-primary': '#7CAEFF',
          '--color-text-critical': '#F6D5D5',
          '--color-text-positive': '#C7E6D5',
          '--color-text-disabled': '#B0B7C3',
          '--color-text-white': '#FDFDFD',
          '--color-link': '#7CAEFF',
          '--color-link-hover': '#5588E4',
          '--color-link-pressed': '#5588E4',
          '--color-link-visited': '#C29EEF',

          '--color-icon-primary': '#5588E4',
          '--color-icon-critical': '#F6D5D5',

          '--color-border': '#69748B',
          '--color-border-subtle': '#394866',
          '--color-border-disabled': '#B0B7C3',
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
