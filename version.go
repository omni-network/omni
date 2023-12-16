// Copyright 2020 The Omni Network Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package omni

var (
	version    string // set from Makefile during build
	commitHash string // set from Makefile during build

	Version = func() string {
		// can happen if the binary is compiled without using the Makefile
		if version == "" {
			version = "dev"
		}
		if commitHash != "" {
			return version + "-" + commitHash
		}
		return version
	}()
)
