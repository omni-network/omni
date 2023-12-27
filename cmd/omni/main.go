// Copyright 2023 The Omni Authors. All rights reserved.
// This file is part of the omni library.
//
// The omni library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The omni library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public License
// along with the omni library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	cmd "github.com/omni-network/omni/cmd/omni/commands"
)

func main() {
	// creat the root command
	rootCmd, err := cmd.GetRootCommand()
	if err != nil {
		panic(err)
	}

	// add other commands
	rootCmd.AddCommand(cmd.PrintVersionCmd())
	rootCmd.AddCommand(cmd.PrintConfigCmd())

	// execute the respective command
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
