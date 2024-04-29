package main

import (
	"io/fs"
	"testing"
)

// MockDirEntry implements fs.DirEntry interface
type MockDirEntry struct {
	name  string
	isDir bool
}

// Implement fs.DirEntry methods
func (m MockDirEntry) Name() string {
	return m.name
}

func (m MockDirEntry) IsDir() bool {
	return m.isDir
}

func (m MockDirEntry) Type() fs.FileMode {
	return 0
}

func (m MockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}

func Test_fileIsImage(t *testing.T) {
	tests := []struct {
		name string
		file fs.DirEntry
		want bool
	}{
		{name: "jpg is an image", file: MockDirEntry{name: "image.jpg"}, want: true},
		{name: "jpeg is an image", file: MockDirEntry{name: "image.jpeg"}, want: true},
		{name: "gif is an image", file: MockDirEntry{name: "image.gif"}, want: true},
		{name: "avif is an image", file: MockDirEntry{name: "image.avif"}, want: true},
		{name: "zip is not an image", file: MockDirEntry{name: "image.zip"}, want: false},
		{name: "html is not an image", file: MockDirEntry{name: "gal.html"}, want: false},
		{name: "a directory is not an image", file: MockDirEntry{name: "example/", isDir: true}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileIsImage(tt.file); got != tt.want {
				t.Errorf("fileIsImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
