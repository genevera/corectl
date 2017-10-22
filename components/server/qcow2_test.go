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
	"os"
	"testing"
)

func TestValidateQcow2(t *testing.T) {
	type args struct {
		fh *os.File
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
			if err := ValidateQcow2(tt.args.fh); (err != nil) != tt.wantErr {
				t.Errorf("ValidateQcow2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_be32(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := be32(tt.args.b); got != tt.want {
				t.Errorf("be32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_be64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := be64(tt.args.b); got != tt.want {
				t.Errorf("be64() = %v, want %v", got, tt.want)
			}
		})
	}
}
