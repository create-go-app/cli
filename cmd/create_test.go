/*
Package cmd includes all of the Create Go App CLI commands.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func newRunCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Test for create command",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), args[0])
			return nil
		},
	}
}

func Test_runCreateCommand(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"successfully",
			args{
				cmd:  newRunCreateCommand(),
				args: []string{},
			},
		},
		{
			"successfully",
			args{
				cmd:  newRunCreateCommand(),
				args: []string{"--use-config"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runInitCommand(tt.args.cmd, tt.args.args)
		})

		os.RemoveAll("./.cgapp.yml")
	}

	cmd := newRunCreateCommand()
	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.SetArgs([]string{"create"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != "create" {
		t.Fatalf("expected \"%s\" got \"%s\"", "create", string(out))
	}

	if currentDir == "" {
		t.Fatalf("expected \"%s\" got \"%s\"", "create", currentDir)
	}
}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConfig()
		})
	}
}
