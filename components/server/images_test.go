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

	"github.com/blang/semver"
)

func Test_timeSlice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    timeSlice
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("timeSlice.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeSlice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    timeSlice
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_timeSlice_Len(t *testing.T) {
	tests := []struct {
		name string
		s    timeSlice
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("timeSlice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localImages(t *testing.T) {
	tests := []struct {
		name      string
		wantLocal map[string]semver.Versions
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocal, err := localImages()
			if (err != nil) != tt.wantErr {
				t.Errorf("localImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLocal, tt.wantLocal) {
				t.Errorf("localImages() = %v, want %v", gotLocal, tt.wantLocal)
			}
		})
	}
}

func TestPullImage(t *testing.T) {
	type args struct {
		channel     string
		version     string
		override    bool
		preferLocal bool
	}
	tests := []struct {
		name    string
		args    args
		wantV   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, err := PullImage(tt.args.channel, tt.args.version, tt.args.override, tt.args.preferLocal)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotV != tt.wantV {
				t.Errorf("PullImage() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func Test_localize(t *testing.T) {
	type args struct {
		channel string
		version string
	}
	tests := []struct {
		name    string
		args    args
		wantB   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := localize(tt.args.channel, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("localize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("localize() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_downloadAndVerify(t *testing.T) {
	type args struct {
		channel string
		version string
	}
	tests := []struct {
		name    string
		args    args
		wantL   map[string]string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := downloadAndVerify(tt.args.channel, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadAndVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("downloadAndVerify() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
