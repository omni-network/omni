// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

// Command verifypr provides a tool to verify charon PRs against the template defined in docs/contibuting.md.
package main

import (
	"testing"
)

func TestVerify(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		commit  string
		wantErr bool
	}{
		{
			name: "valid no issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: none`,
		},
		{
			name:   "valid mixed line endings", // Github uses \r\n
			commit: "feat(*): add foo bar\n\nfoo bar baz\r\n\r\nissue: none",
		},
		{
			name: "valid single scope",
			commit: `feat(foo): add foo bar

foo bar baz

issue: none`,
		},
		{
			name: "invalid double scope",
			commit: `feat(foo/baz): add foo bar

foo bar baz

issue: none`,
		},
		{
			name:    "invalid too much scope",
			wantErr: true,
			commit: `feat(foo/baz/bar): add foo bar

foo bar baz

issue: none`,
		},
		{
			name:    "invalid type",
			wantErr: true,
			commit: `foo(*): add foo bar

foo bar baz

issue: none`,
		},
		{
			name: "valid full github issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: https://github.com/omni-network/omni/issues/1334`,
		},
		{
			name: "valid short github issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: #1334`,
		},
		{
			name: "fix valid short github issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: fix #1334`,
		},
		{
			name: "resolves valid short github issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: resolves #1334`,
		},
		{
			name: "closed full github issue",
			commit: `feat(*): add foo bar

foo bar baz

issue: closed https://github.com/omni-network/omni/issues/1334`,
		},
		{
			name:    "invalid description title case",
			wantErr: true,
			commit: `feat(*): Add foo bar

foo bar baz

issue: none`,
		},
		{
			name:    "invalid link short github issue",
			wantErr: true,
			commit: `feat(*): add foo bar

foo bar baz

issue: ifix #1334`,
		},
		{
			name:    "invalid description punctuation",
			wantErr: true,
			commit: `feat(*): foo, baz, bar.

foo bar baz

issue: none`,
		},
		{
			name:    "invalid no scope",
			wantErr: true,
			commit: `feat: foo baz bar

foo bar baz

issue: none`,
		},
		{
			name:    "invalid no body",
			wantErr: true,
			commit: `feat: foo baz bar

issue: none`,
		},
		{
			name: "valid other type",
			commit: `ci(*): foo baz bar

foo bar baz

issue: none`,
		},
		{
			name: "valid example",
			commit: `ci(github): add verifypr action

Adds a 'verifypr' github action that ensures all PR adhere to the omni style conventional commit template.

issue: https://github.com/omni-network/omni/issues/1334`,
		},
		{
			name: "valid description with dashes",
			commit: `ci(github/workflows): add pre-commit and golangci-lint actions

Adds two github actions:
 - golangci-lint: Runs the go linter so issues are added inline to PR.
 - pre-commit: Runs pre-commit hooks (excluding golangci-lint)

issue: none`,
		},
		{
			name: "valid space around none issue footer",
			commit: `ci(*): foo baz bar

foo bar baz

issue:  none `,
		},
		{
			name: "valid multiple footers",
			commit: `ci(*): foo baz bar

foo bar baz

more: stuff
issue:  none `,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := verify(tt.commit)
			if (err != nil) != tt.wantErr {
				t.Fatalf("verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
