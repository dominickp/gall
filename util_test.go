package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/afero"
)

type MockFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (m *MockFileInfo) Name() string       { return m.name }
func (m *MockFileInfo) Size() int64        { return m.size }
func (m *MockFileInfo) Mode() os.FileMode  { return m.mode }
func (m *MockFileInfo) ModTime() time.Time { return m.modTime }
func (m *MockFileInfo) IsDir() bool        { return m.isDir }
func (m *MockFileInfo) Sys() interface{}   { return nil }

func Test_fileIsImage(t *testing.T) {
	tests := []struct {
		name string
		file fs.FileInfo
		want bool
	}{
		{name: "jpg is an image", file: &MockFileInfo{name: "image.jpg"}, want: true},
		{name: "jpeg is an image", file: &MockFileInfo{name: "image.jpeg"}, want: true},
		{name: "gif is an image", file: &MockFileInfo{name: "image.gif"}, want: true},
		{name: "avif is an image", file: &MockFileInfo{name: "image.avif"}, want: true},
		{name: "zip is not an image", file: &MockFileInfo{name: "image.zip"}, want: false},
		{name: "html is not an image", file: &MockFileInfo{name: "gal.html"}, want: false},
		{name: "a directory is not an image", file: &MockFileInfo{name: "example/", isDir: true}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileIsImage(tt.file); got != tt.want {
				t.Errorf("fileIsImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImagesInDirectory(t *testing.T) {
	// Create a new in-memory filesystem
	fs := afero.NewMemMapFs()

	tests := []struct {
		name      string
		fileNames []string
		want      []string
	}{
		{
			name:      "all JPEGs are present",
			fileNames: []string{"image1.jpg", "image2.jpeg"},
			want:      []string{"image1.jpg", "image2.jpeg"},
		},
		{
			name:      "non-images are not present",
			fileNames: []string{"image1.jpg", "image2.jpg", "archive.zip"},
			want:      []string{"image1.jpg", "image2.jpg"},
		},
	}
	for ti, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new directory
			dir := fmt.Sprintf("/testDir-%d", ti)
			err := fs.MkdirAll(dir, 0755)
			if err != nil {
				t.Fatal(err)
			}

			// Create new files
			for _, file := range tt.fileNames {
				afero.WriteFile(fs, dir+"/"+file, []byte{}, 0644)
			}

			// Call getImagesInDirectory
			images := getImagesInDirectory(fs, dir)

			if len(images) != len(tt.want) {
				t.Errorf("Expected %d images, got %d", len(tt.want), len(images))
			}

			for index, file := range images {
				// Assuming `expected` is the expected file path and `file` is the fs.FileInfo object
				if filepath.Base(tt.want[index]) != file.Name() {
					t.Errorf("Expected %v, got %v", filepath.Base(tt.want[index]), file.Name())
				}
			}
		})
	}
}
