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

package session

import (
	"os/user"
	"reflect"
	"testing"

	"github.com/TheNewNormal/corectl/release"
	"github.com/spf13/viper"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantCtx *Context
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtx, tt.wantCtx) {
				t.Errorf("New() = %v, want %v", gotCtx, tt.wantCtx)
			}
		})
	}
}

func TestNetwork_Base(t *testing.T) {
	type fields struct {
		Address string
		Mask    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Network{
				Address: tt.fields.Address,
				Mask:    tt.fields.Mask,
			}
			if got := n.Base(); got != tt.want {
				t.Errorf("Network.Base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_Debug(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.Debug(); got != tt.want {
				t.Errorf("Context.Debug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_JSON(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.JSON(); got != tt.want {
				t.Errorf("Context.JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_ConfigDir(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.ConfigDir(); got != tt.want {
				t.Errorf("Context.ConfigDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_ImageStore(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.ImageStore(); got != tt.want {
				t.Errorf("Context.ImageStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_RunDir(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.RunDir(); got != tt.want {
				t.Errorf("Context.RunDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_TmpDir(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.TmpDir(); got != tt.want {
				t.Errorf("Context.TmpDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_EtcDir(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if got := ctx.EtcDir(); got != tt.want {
				t.Errorf("Context.EtcDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext_NormalizeOnDiskLayout(t *testing.T) {
	type fields struct {
		Privileged    bool
		Meta          *release.Info
		CmdLine       *viper.Viper
		ServerAddress string
		User          *user.User
		Network       *Network
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				Privileged:    tt.fields.Privileged,
				Meta:          tt.fields.Meta,
				CmdLine:       tt.fields.CmdLine,
				ServerAddress: tt.fields.ServerAddress,
				User:          tt.fields.User,
				Network:       tt.fields.Network,
			}
			if err := ctx.NormalizeOnDiskLayout(); (err != nil) != tt.wantErr {
				t.Errorf("Context.NormalizeOnDiskLayout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetwork_SetContext(t *testing.T) {
	type fields struct {
		Address string
		Mask    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Network{
				Address: tt.fields.Address,
				Mask:    tt.fields.Mask,
			}
			if err := ctx.SetContext(); (err != nil) != tt.wantErr {
				t.Errorf("Network.SetContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecutable(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Executable(); got != tt.want {
				t.Errorf("Executable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppName(); got != tt.want {
				t.Errorf("AppName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutableFolder(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecutableFolder(); got != tt.want {
				t.Errorf("ExecutableFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
