# Omni Release Process

## Overview

This document outlines the release process for the Omni monorepo. Currently the process focuses on internal compatibility, simplicity, and rapid iteration.

## Versioning Strategy

- **Semantic Versioning**: A single semantic version is used for the entire repository to ensure intra-compatibility across all components.
- **No Pre-release Tags**: Versions are released without `-alpha`, `-beta`, or similar tags. Bugs are addressed by swiftly releasing a new patch version.
- **Release Schedule**: New versions are released by discussion and published Wednesdays, following a two-week sprint cycle or as needed for urgent features or fixes.

## Components Covered

- Halo
- Relayer
- Monitor
- Explorer-GraphQL
- Explorer-Indexer
- Indexer
- Contracts

## Release Process

1. **Internal Testing**: Perform a preliminary check on the release.
2. **Version Bump**: Update the version in `lib/buildinfo/buildinfo.go` to the desired `vX.Y.Z`, then push the branch and open a pull request.
3. **Draft Release Notes**: After tagging the version with a `git tag`, the CI pipeline drafts the release and publishes Docker containers on DockerHub.
4. **Discussion and Publishing**: Release draft is discussed internally for review before publication.

## Release Notes Format

Release notes provide a comprehensive overview of updates, including new features, bug fixes, and breaking changes for each component.

### CI for Release Notes

The release notes generation process is semi-automated through the CI, see [.goreleaser.yaml](../.goreleaser.yaml).
