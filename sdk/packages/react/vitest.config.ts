import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: ['src/**/*.test.{ts,tsx}'],
    environment: 'happy-dom',
    setupFiles: ['./test/unitSetup.ts'],
    typecheck: {
      enabled: true,
    },
    coverage: {
      reporter: ['text', 'json', 'html'],
      // TODO use lcov when running in CI
      // reporter: process.env.CI ? ['lcov'] : ['text', 'json', 'html'],
    },
  },
})
