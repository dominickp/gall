package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func getImagesInDirectory(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return nil
	}
	images := []fs.DirEntry{}
	for _, file := range files {
		if fileIsImage(file) {
			images = append(images, file)
		}
	}
	log.Printf("Found %d images in the directory", len(images))
	return images
}
