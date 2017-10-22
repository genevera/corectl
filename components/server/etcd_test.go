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

//  adapted from github.com/kubernetes/minikube/pkg/localkube/etcd.go
//

package server

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/coreos/etcd/pkg/types"
)

func TestServerContext_NewEtcd(t *testing.T) {
	type args struct {
		clientURLStrs []string
		peerURLStrs   []string
		name          string
		dataDirectory string
	}
	tests := []struct {
		name    string
		d       *ServerContext
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.NewEtcd(tt.args.clientURLStrs, tt.args.peerURLStrs, tt.args.name, tt.args.dataDirectory); (err != nil) != tt.wantErr {
				t.Errorf("ServerContext.NewEtcd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEtcdServer_Start(t *testing.T) {
	tests := []struct {
		name    string
		e       *EtcdServer
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Start(); (err != nil) != tt.wantErr {
				t.Errorf("EtcdServer.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEtcdServer_Stop(t *testing.T) {
	tests := []struct {
		name string
		e    *EtcdServer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Stop()
		})
	}
}

func TestEtcdServer_requestTimeout(t *testing.T) {
	tests := []struct {
		name string
		e    *EtcdServer
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.requestTimeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtcdServer.requestTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createListeners(t *testing.T) {
	type args struct {
		urls types.URLs
	}
	tests := []struct {
		name          string
		args          args
		wantListeners []net.Listener
		wantErr       bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotListeners, err := createListeners(tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("createListeners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotListeners, tt.wantListeners) {
				t.Errorf("createListeners() = %v, want %v", gotListeners, tt.wantListeners)
			}
		})
	}
}
