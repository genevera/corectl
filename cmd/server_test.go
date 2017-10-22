// Copyright (c) 2016 by António Meireles  <antonio.meireles@reformi.st>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func Test_uuidToMacCommand(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := uuidToMacCommand(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("uuidToMacCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shutdownCommand(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := shutdownCommand(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("shutdownCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serverStartCommand(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := serverStartCommand(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("serverStartCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
