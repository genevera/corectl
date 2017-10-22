//
// quick hack based in https://github.com/kelseyhightower/cpic/image
//

package image

import (
	"bytes"
	"compress/gzip"
	"io"
	"reflect"
	"testing"
)

func TestNewReader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Reader
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReader(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Close(t *testing.T) {
	type fields struct {
		z *gzip.Reader
		c *cpio.Reader
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
			r := &Reader{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Reader.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewWriter(t *testing.T) {
	tests := []struct {
		name    string
		want    *Writer
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got, err := NewWriter(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewWriter() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestWriter_Close(t *testing.T) {
	type fields struct {
		z *gzip.Writer
		c *cpio.Writer
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
			w := &Writer{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			if err := w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Writer.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_Write(t *testing.T) {
	type fields struct {
		z *gzip.Writer
		c *cpio.Writer
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			got, err := w.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Writer.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_WriteHeader(t *testing.T) {
	type fields struct {
		z *gzip.Writer
		c *cpio.Writer
	}
	type args struct {
		hdr *cpio.Header
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			if err := w.WriteHeader(tt.args.hdr); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		dst *Writer
		src *Reader
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
			if err := Copy(tt.args.dst, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteToFile(t *testing.T) {
	type fields struct {
		z *gzip.Writer
		c *cpio.Writer
	}
	type args struct {
		content *bytes.Buffer
		name    string
		mode    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			if err := w.WriteToFile(tt.args.content, tt.args.name, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriter_WriteDir(t *testing.T) {
	type fields struct {
		z *gzip.Writer
		c *cpio.Writer
	}
	type args struct {
		name string
		mode int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				z: tt.fields.z,
				c: tt.fields.c,
			}
			if err := w.WriteDir(tt.args.name, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("Writer.WriteDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
