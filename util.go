package main

import (
	"io/fs"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/afero"
)

func openBrowser(url string, browser string) error {
	validBrowsers := []string{"firefox", "default"}
	isValidBrowserOption := false
	for _, b := range validBrowsers {
		if b == browser {
			isValidBrowserOption = true
			break
		}
	}
	if !isValidBrowserOption {
		log.Fatalf("Invalid browser option: %s", browser)
		return nil
	}

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
		if browser == "firefox" {
			args = append(args, "firefox", "-private-window")
		}
	default:
		if browser == "firefox" {
			cmd = "firefox"
			args = []string{"-private-window"}
		} else {
			cmd = "xdg-open"
			args = []string{}
		}
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func fileIsVideo(file fs.FileInfo) bool {
	videoFileTypes := []string{
		".mp4",
		".webm",
		".ogg",
		".mov",
		".mpg",
	}
	for _, fileType := range videoFileTypes {
		if strings.HasSuffix(file.Name(), fileType) {
			return true
		}
	}
	return false
}

// fileIsImage checks if a file is a browser-compatible image based on its extension
func fileIsImage(file fs.FileInfo) bool {
	// https://developer.mozilla.org/en-US/docs/Web/Media/Formats/Image_types
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

// getImagesInDirectory reads the files in a directory and returns a list of images
func getImagesInDirectory(afs afero.Fs, dir string) []fs.FileInfo {
	files, err := afero.ReadDir(afs, dir)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return nil
	}
	images := []fs.FileInfo{}
	for _, file := range files {
		if fileIsImage(file) || fileIsVideo(file) {
			images = append(images, file)
		}
	}
	log.Printf("Found %d images in the directory", len(images))
	log.Printf("%d non-images were excluded from the gallery", len(files)-len(images))
	return images
}
