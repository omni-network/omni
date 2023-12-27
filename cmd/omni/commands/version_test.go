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
	"bytes"
	"fmt"
	"github.com/cometbft/cometbft/version"
	"github.com/omni-network/omni"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	// set the version, just like it is set from ther Makefile
	omni.Version = "1.2.3-0xabcdef"

	// create the command
	rootCmd := &cobra.Command{
		Use: "omni",
	}
	config.cmd = rootCmd
	cmd := PrintVersionCmd()

	// execute the command and get the output
	got := runCommandAndCaptureStdio(t, cmd)

	//check the output with the expected results
	want := "Omni-Node    : 1.2.3-0xabcdef\nOmni-Halo    : 0.0.1\n"
	if got != want {
		t.Errorf("got output %q, want %q", got, want)
	}
}

func TestVersionCmdWithVerbosity(t *testing.T) {
	// set the version, just like it is set from ther Makefile
	omni.Version = "1.2.3-0xabcdef"

	// create the command
	rootCmd := &cobra.Command{
		Use: "omni",
	}
	config.cmd = rootCmd
	cmd := PrintVersionCmd()

	// set the verbosity to true
	config.viperConfig = viper.New()
	config.viperConfig.Set(optionNameVerbose, "true")

	// execute the command and get the output
	got := runCommandAndCaptureStdio(t, cmd)

	//check the output with the expected results
	want := fmt.Sprintf("Omni-Node    : 1.2.3-0xabcdef\nOmni-Halo    : 0.0.1\nCometBFT-Core: %v\nCometBFT-ABCI: %v\nCometBFT-P2P : %v\nCometBFT-BlockProtocol: %v\n",
		version.TMCoreSemVer,
		version.ABCIVersion,
		version.P2PProtocol,
		version.BlockProtocol)
	if got != want {
		t.Errorf("got output %q, want %q", got, want)
	}
}

func runCommandAndCaptureStdio(t *testing.T, cmd *cobra.Command) string {
	// replace stdio with a pipe
	backupOfStdIO := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// run the comand
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatal(err)
	}

	// cleanup
	w.Close()
	os.Stdout = backupOfStdIO

	// capture the screen output and send for verification
	out := <-outC
	return out
}
