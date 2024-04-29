package main

import (
	"io/fs"
	"os/exec"
	"runtime"
	"strings"
)

// FIXME: opens only Firefox in private mode
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", "firefox", "-private-window"}
	default:
		cmd = "firefox"
		args = []string{"-private-window"}
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func fileIsImage(file fs.DirEntry) bool {
	// Check if the file is an image
	imageFileTypes := []string{
		".apng",                                    // APNG
		".avif",                                    // AVIF
		".gif",                                     // GIF
		".jpg", ".jpeg", ".jfif", ".pjpeg", ".pjp", // JPEG
		".png",         // PNG
		".svg",         // SVG
		".webp",        // WebP
		".bmp",         // BMP
		".ico", ".cur", // ICO
		".tiff", ".tif", // TIFF

	}
	for _, fileType := range imageFileTypes {
		if strings.HasSuffix(file.Name(), fileType) {
			return true
		}
	}
	return false
}
