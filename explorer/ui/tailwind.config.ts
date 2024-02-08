import type { Config } from "tailwindcss";
const typography = require('@tailwindcss/typography')
const daisyui = require('daisyui')

export default ({
  content: [
    "./app/**/*.{html,js,ts,jsx,tsx}",
    "~/components/**/*.{html,js,ts,tsx,jsx}"
  ],
  theme: {
    extend: {
      fontFamily: {
        urbanist: ["Urbanist", "sans-serif"],
      },
    },
  },
  plugins: [typography, daisyui],
  daisyui: {
    themes: true,
    darkTheme: "dark", // name of one of the included themes for dark mode
    base: true, // applies background color and foreground color for root element by default
    styled: true, // include daisyUI colors and design decisions for all components
    utils: true, // adds responsive and modifier utility classes
    prefix: "", // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
    logs: true, // Shows info about daisyUI version and used config in the console when building your CSS
    themeRoot: ":root", // The element that receives theme color CSS variables
    rtl: false, // Enable RTL support
  },
} satisfies Config);
