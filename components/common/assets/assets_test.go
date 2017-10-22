package assets

import "testing"

func TestContents(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name  string
		args  args
		wantT string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := Contents(tt.args.file); gotT != tt.wantT {
				t.Errorf("Contents() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
