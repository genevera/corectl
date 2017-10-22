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
	"net/http"
	"reflect"
	"testing"

	"github.com/TheNewNormal/corectl/release"
)

func Test_rpcServiceSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpcServiceSetup()
		})
	}
}

func TestRPCservice_Echo(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Echo(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.Echo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_HandlesNFS(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.HandlesNFS(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.HandlesNFS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_AvailableImages(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AvailableImages(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.AvailableImages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_RemoveImage(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.RemoveImage(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.RemoveImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_UUIDtoMACaddr(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UUIDtoMACaddr(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.UUIDtoMACaddr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_Run(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Run(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_Stop(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Stop(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_ActiveVMs(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.ActiveVMs(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.ActiveVMs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCservice_StopVMs(t *testing.T) {
	type args struct {
		r     *http.Request
		args  *RPCquery
		reply *RPCreply
	}
	tests := []struct {
		name    string
		s       *RPCservice
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.StopVMs(tt.args.r, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("RPCservice.StopVMs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRPCQuery(t *testing.T) {
	type args struct {
		f    string
		args *RPCquery
	}
	tests := []struct {
		name      string
		args      args
		wantReply *RPCreply
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReply, err := RPCQuery(tt.args.f, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("RPCQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReply, tt.wantReply) {
				t.Errorf("RPCQuery() = %v, want %v", gotReply, tt.wantReply)
			}
		})
	}
}

func TestServerContext_Running(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ServerContext
		wantI   *release.Info
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := tt.cfg.Running()
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerContext.Running() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("ServerContext.Running() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
