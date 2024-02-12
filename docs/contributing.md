# Contributing to Omni Network

Thank you for considering contributing to the Omni Network! We appreciate contributions of all forms, from code to documentation. Here's how you can help:

## Branching Model

We use Trunk Based Development for managing branches, please refer to the [Branching Model](./branching.md) for more details.

## How to Contribute

### Testing

- **Unit Testing:** Isolate tests to specific code paths or failure modes within functions. Ensure each function's failure modes and code paths are thoroughly tested.
- **Integration Testing:** Focus on testing sophisticated interactions between components.

### Issues

- Use the search tool to avoid duplicate issues.
- Provide detailed reports including source code and commit SHA for bugs.
- Engage with existing issues by providing feedback or reactions.

### Pull Requests

- Open PRs against the `main` branch only.
- Include tests for new or modified code.
- Follow the Go guidelines shown in the [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) and this [Solidity Style Guide](https://gist.github.com/lucas-manuel/a43da80cdd4c3f37a2f3151d3774b8e0).
- Ensure all exported types outside of the internal package are well-documented.
- Run linters and tests locally before submitting your PR, install the pre-commit hooks by running `make install-pre-commit`.
