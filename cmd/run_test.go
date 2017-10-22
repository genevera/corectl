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
	"reflect"
	"testing"

	"github.com/TheNewNormal/corectl/components/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Test_runCommand(t *testing.T) {
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
			if err := runCommand(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bootIt(t *testing.T) {
	type args struct {
		vm *server.VMInfo
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
			if err := bootIt(tt.args.vm); (err != nil) != tt.wantErr {
				t.Errorf("bootIt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_vmBootstrap(t *testing.T) {
	type args struct {
		args *viper.Viper
	}
	tests := []struct {
		name    string
		args    args
		wantVm  *server.VMInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVm, err := vmBootstrap(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("vmBootstrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVm, tt.wantVm) {
				t.Errorf("vmBootstrap() = %v, want %v", gotVm, tt.wantVm)
			}
		})
	}
}

func Test_runFlagsDefaults(t *testing.T) {
	type args struct {
		setFlag *pflag.FlagSet
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runFlagsDefaults(tt.args.setFlag)
		})
	}
}
