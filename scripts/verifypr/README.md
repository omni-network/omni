# verifypr

Simple script that is called by [verifypr](../../.github/workflows/verifypr.yml) github action
that verifies omni PRs against the conventional commit template.

See supported types [here](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional#type-enum)

> Note it has its own go.mod since it depends on conventionalcommit.org's go library that isn't required by the main omni codebase.
