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

package commands

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"testing"
)

func TestPrintConfigCmd(t *testing.T) {
	// create the command
	rootCmd := &cobra.Command{
		Use: "omni",
	}
	config.cmd = rootCmd
	cmd := PrintConfigCmd()

	// execute the command and get the output
	got := runCommandAndCaptureStdio(t, cmd)

	//check the output with the expected results (except the node name which is created at random every time)
	want, err := json.MarshalIndent(config.omniConfig, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want)+"\n" {
		t.Errorf("got output %q, want %q", got, want)
	}
}

func TestPrintConfigCmdWithVerbosity(t *testing.T) {
	// create the command
	rootCmd := &cobra.Command{
		Use: "omni",
	}
	config.cmd = rootCmd
	cmd := PrintConfigCmd()

	// set the verbosity to true
	config.viperConfig = viper.New()
	config.viperConfig.Set(optionNameVerbose, "true")

	// execute the command and get the output
	got := runCommandAndCaptureStdio(t, cmd)

	//check the output with the expected results (except the node name which is created at random every time)
	want1, err := json.MarshalIndent(config.omniConfig, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	want2, err := json.MarshalIndent(config.cometBFTConfig, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	want := string(want1) + "\n" + string(want2) + "\n"
	if got != want {
		t.Errorf("got output %q, want %q", got, want)
	}
}
