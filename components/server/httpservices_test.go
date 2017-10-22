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
	"testing"
)

func Test_httpServiceSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpServiceSetup()
		})
	}
}

func Test_remoteIP(t *testing.T) {
	type args struct {
		s string
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
			if got := remoteIP(tt.args.s); got != tt.want {
				t.Errorf("remoteIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLoopback(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLoopback(tt.args.s); got != tt.want {
				t.Errorf("isLoopback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpError(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		status int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpError(tt.args.w, tt.args.status)
		})
	}
}

func Test_acceptableRequest(t *testing.T) {
	type args struct {
		r *http.Request
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := acceptableRequest(tt.args.r, tt.args.w); got != tt.want {
				t.Errorf("acceptableRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpInstanceCloudConfig(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpInstanceCloudConfig(tt.args.w, tt.args.r)
		})
	}
}

func Test_isPortOpen(t *testing.T) {
	type args struct {
		t      string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPortOpen(tt.args.t, tt.args.target); got != tt.want {
				t.Errorf("isPortOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpInstanceIgnitionConfig(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpInstanceIgnitionConfig(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpInstanceExternalConnectivity(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpInstanceExternalConnectivity(tt.args.w, tt.args.r)
		})
	}
}

func Test_httpInstanceCallback(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpInstanceCallback(tt.args.w, tt.args.r)
		})
	}
}
