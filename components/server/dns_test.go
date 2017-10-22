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

//  adapted from github.com/kubernetes/minikube/pkg/localkube/dns.go
//

package server

import (
	"bytes"
	"testing"
	"time"
)

func TestServerContext_NewDNSServer(t *testing.T) {
	type args struct {
		root          string
		serverAddress string
		ns            []string
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
			if err := tt.d.NewDNSServer(tt.args.root, tt.args.serverAddress, tt.args.ns); (err != nil) != tt.wantErr {
				t.Errorf("ServerContext.NewDNSServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDNSServer_PortForward(t *testing.T) {
	tests := []struct {
		name    string
		dns     *DNSServer
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dns.PortForward(); (err != nil) != tt.wantErr {
				t.Errorf("DNSServer.PortForward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDNSServer_Start(t *testing.T) {
	tests := []struct {
		name string
		dns  *DNSServer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dns.Start()
		})
	}
}

func TestDNSServer_Stop(t *testing.T) {
	tests := []struct {
		name string
		dns  *DNSServer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dns.Stop()
		})
	}
}

func Test_teardownService(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			teardownService()
		})
	}
}

func Test_invertDomain(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := invertDomain(tt.args.in); gotOut != tt.wantOut {
				t.Errorf("invertDomain() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestDNSServer_addRecord(t *testing.T) {
	type args struct {
		hostName string
		ip       string
	}
	tests := []struct {
		name    string
		d       *DNSServer
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.addRecord(tt.args.hostName, tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("DNSServer.addRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDNSServer_rmRecord(t *testing.T) {
	type args struct {
		hostName string
		ip       string
	}
	tests := []struct {
		name    string
		d       *DNSServer
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.rmRecord(tt.args.hostName, tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("DNSServer.rmRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_until(t *testing.T) {
	type args struct {
		fn    func() error
		name  string
		sleep time.Duration
		done  <-chan struct{}
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			until(tt.args.fn, w, tt.args.name, tt.args.sleep, tt.args.done)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("until() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_pad(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pad(tt.args.str); got != tt.want {
				t.Errorf("pad() = %v, want %v", got, tt.want)
			}
		})
	}
}
