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

package server

import (
	"reflect"
	"testing"
)

func TestVMInfo_ValidateCDROM(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		vm      *VMInfo
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.ValidateCDROM(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.ValidateCDROM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_ValidateVolumes(t *testing.T) {
	type args struct {
		volumes []string
		root    bool
	}
	tests := []struct {
		name    string
		vm      *VMInfo
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.ValidateVolumes(tt.args.volumes, tt.args.root); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.ValidateVolumes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_ValidateCloudConfig(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		vm      *VMInfo
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.ValidateCloudConfig(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.ValidateCloudConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_SSHkeyGen(t *testing.T) {
	tests := []struct {
		name    string
		vm      *VMInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.SSHkeyGen(); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.SSHkeyGen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_assembleBootPayload(t *testing.T) {
	tests := []struct {
		name      string
		vm        *VMInfo
		wantXArgs []string
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXArgs, err := tt.vm.assembleBootPayload()
			if (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.assembleBootPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotXArgs, tt.wantXArgs) {
				t.Errorf("VMInfo.assembleBootPayload() = %v, want %v", gotXArgs, tt.wantXArgs)
			}
		})
	}
}

func TestVMs_gracefullyShutdown(t *testing.T) {
	tests := []struct {
		name string
		list VMs
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.gracefullyShutdown()
		})
	}
}

func TestVMInfo_kill(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vm.kill()
		})
	}
}

func TestVMInfo_gracefullyShutdown(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vm.gracefullyShutdown()
		})
	}
}

func TestVMInfo_lookup(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.lookup(); got != tt.want {
				t.Errorf("VMInfo.lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMInfo_register(t *testing.T) {
	tests := []struct {
		name    string
		vm      *VMInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.register(); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_deregister(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vm.deregister()
		})
	}
}

func TestVMInfo_endpoint(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.endpoint(); got != tt.want {
				t.Errorf("VMInfo.endpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMInfo_RunDir(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.RunDir(); got != tt.want {
				t.Errorf("VMInfo.RunDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMInfo_MkRunDir(t *testing.T) {
	tests := []struct {
		name    string
		vm      *VMInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vm.MkRunDir(); (err != nil) != tt.wantErr {
				t.Errorf("VMInfo.MkRunDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVMInfo_Log(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.Log(); got != tt.want {
				t.Errorf("VMInfo.Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMInfo_TTY(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.TTY(); got != tt.want {
				t.Errorf("VMInfo.TTY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMInfo_PrettyPrint(t *testing.T) {
	tests := []struct {
		name string
		vm   *VMInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vm.PrettyPrint()
		})
	}
}

func TestStorageAssets_PrettyPrint(t *testing.T) {
	type args struct {
		root bool
	}
	tests := []struct {
		name    string
		volumes *StorageAssets
		args    args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.volumes.PrettyPrint(tt.args.root)
		})
	}
}

func TestVMs_Len(t *testing.T) {
	tests := []struct {
		name string
		run  VMs
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.run.Len(); got != tt.want {
				t.Errorf("VMs.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMs_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		run  VMs
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.run.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestVMs_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		run  VMs
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.run.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("VMs.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVMmap_array(t *testing.T) {
	tests := []struct {
		name    string
		in      VMmap
		wantOut VMs
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.in.array(); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("VMmap.array() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
