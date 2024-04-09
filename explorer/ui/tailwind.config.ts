import type { Config } from 'tailwindcss'

export default {
  content: ['./app/**/*.{html,js,ts,jsx,tsx}', '~/components/**/*.{html,js,ts,tsx,jsx}'],
  theme: {
    extend: {
      fontFamily: {
        manrope: ['Manrope'],
      },
      fontSize: {
        h1: [
          '2.625rem',
          {
            lineHeight: '58px',
            letterSpacing: '0.5px',
            fontWeight: 500,
          },
        ],
        h2: [
          '2.25rem',
          {
            lineHeight: '50px',
            letterSpacing: '0.5px',
            fontWeight: 500,
          },
        ],
        h3: [
          '1.875rem',
          {
            lineHeight: '42px',
            letterSpacing: '0.5px',
            fontWeight: 500,
          },
        ],
        h4: [
          '1.5rem',
          {
            lineHeight: '34px',
            letterSpacing: '0.5px',
            fontWeight: 500,
          },
        ],
        h5: [
          '1.25rem',
          {
            lineHeight: '28px',
            letterSpacing: '0.5px',
            fontWeight: 500,
          },
        ],
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
        },
        dark: {
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          ...require("daisyui/src/theming/themes")["dark"],
        },
      },
    ],
    base: true, // applies background color and foreground color for root element by default
    styled: true, // include daisyUI colors and design decisions for all components
    utils: true, // adds responsive and modifier utility classes
    prefix: '', // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
    logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
    themeRoot: ':root', // The element that receives theme color CSS variables
    rtl: false, // Enable RTL support
  },
} satisfies Config
