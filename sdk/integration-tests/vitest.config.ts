import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: ['./suites/**/*.{ts,tsx}'],
    environment: 'happy-dom',
  },
})
