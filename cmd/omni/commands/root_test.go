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
	conf "github.com/omni-network/omni/pkg/config"
	"gopkg.in/yaml.v3"
	"gotest.tools/assert"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestConfigPrecedence(t *testing.T) {
	// Set argument from command line
	t.Run("set flag", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "*")
		if err != nil {
			defer func(path string) {
				err := os.RemoveAll(path)
				if err != nil {

				}
			}(tmpDir)
			t.Fatal(err)
		}

		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}

		rootCmd.SetArgs([]string{"--home-dir", tmpDir})
		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		gotOutput, err := rootCmd.Flags().GetString(optionNameHomeDir)
		if err != nil {
			t.Fatal(err)
		}

		wantOutput := tmpDir
		assert.Equal(t, wantOutput, gotOutput, "expected %v but got %v", wantOutput, gotOutput)
	})

	//  Check if the command line arguments are written in the file properly
	t.Run("read from file", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "*")
		if err != nil {
			defer func(path string) {
				err := os.RemoveAll(path)
				if err != nil {

				}
			}(tmpDir)
			t.Fatal(err)
		}

		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}

		rootCmd.SetArgs([]string{"--home-dir", tmpDir})
		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		cFileName := filepath.Join(config.omniConfig.OmniConfigDir, config.omniConfig.OmniConfigFileName+".yml")
		stat, err := os.Stat(cFileName)
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != config.omniConfig.OmniConfigFileName+".yml" {
			t.Fatal(err)
		}

		yamlFile, err := os.ReadFile(cFileName)
		if err != nil {
			t.Fatal(err)
		}

		gotConfig := &conf.OmniConfig{}
		err = yaml.Unmarshal(yamlFile, gotConfig)
		if err != nil {
			t.Fatal(err)
		}

		res := reflect.DeepEqual(config.omniConfig, gotConfig)
		assert.Equal(t, true, res, "expected %v but got %v", true, res)

	})

	// Set arguments from environment variables
	t.Run("read env var", func(t *testing.T) {
		err := os.Setenv("OMNI_LOG-LEVEL", "myLogLevel")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err := os.Unsetenv("OMNI_LOG-LEVEL")
			if err != nil {

			}
		}()

		rootCmd, err := NewRootCommand()
		if err != nil {
			t.Fatal(err)
		}

		err = rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		gotOutput, err := rootCmd.Flags().GetString(optionNameLogLevel)
		if err != nil {
			t.Fatal(err)
		}

		wantOutput := "myLogLevel"
		assert.Equal(t, wantOutput, gotOutput, "expected %v but got %v", wantOutput, gotOutput)
	})
}
