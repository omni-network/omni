import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    workspace: [
      {
        test: {
          name: 'unit',
          include: ['src/**/*.test.{ts,tsx}'],
          environment: 'happy-dom',
          setupFiles: ['./test/unitSetup.ts'],
          typecheck: {
            enabled: true,
          },
        },
      },
      {
        test: {
          name: 'integration',
          include: ['integration-tests/**/*.test.{ts,tsx}'],
          environment: 'browser',
          browser: {
            enabled: true,
            headless: true,
            provider: 'playwright',
            instances: [{ browser: 'chromium' }],
          },
        },
      },
    ],
    coverage: {
      reporter: ['text', 'json', 'html'],
      // TODO use lcov when running in CI
      // reporter: process.env.CI ? ['lcov'] : ['text', 'json', 'html'],
    },
  },
})
