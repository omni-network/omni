import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    testTimeout: 60000, // 1m, orders take time to process
    include: ['./suites/**/*.{ts,tsx}'],
    environment: 'happy-dom',
  },
})
