# Omni Release Process

## Overview

This document outlines the release process for the Omni monorepo.

## Components Covered

Note that a single semantic version is used for the entire repository to ensure compatibility across all components:

- Halo
- Relayer
- Monitor
- Contracts
- CLI

## Release Process

- Releases are cut from release branches, named `release/v{X.Y.Z}` (not the main branch).
- Then push a version upgrade commit to update `lib/buildinfo/buildinfo.go`.
- Versions are released without `-alpha`, `-beta`, or similar tags. Bugs are addressed by swiftly releasing a new patch version.
- After tagging the version with a `git tag`, the CI pipeline drafts the release and publishes Docker containers on DockerHub.

## Release Notes Format

Release notes provide a comprehensive overview of updates, including new features, bug fixes, and breaking changes for each component.
