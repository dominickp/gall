package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"runtime"
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

func createHTMLPage(dir string, images []fs.DirEntry) string {
	// Create a new file in the directory
	fileName := dir + "/gal.html"
	// Get absolute path to the file
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fileName = absPath + "/" + fileName
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	// Write the HTML content to the file
	file.WriteString(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Gallery</title>
		<style>
			main { 
				display: grid;
				grid-template-columns: repeat(auto-fill, minmax(8rem, 1fr) minmax(14rem, 2fr)) minmax(8rem, 1fr);
				grid-template-rows: masonry; 
				gap: 1rem;
			}
			figure {
				position: relative;
				counter-increment: item-counter;
				margin: 0;
			}
			img {
				width: 100%;
				height: auto;
				display: block;
			}
		</style>
		</head>
		<body>
		<main class="grid">
	`)
	for _, image := range images {
		file.WriteString("<figure class='card'><img src='" + image.Name() + "' /></figure>")
	}
	file.WriteString("</main></body></html>")
	return fileName
}

func main() {

	dir := os.Args[1:]

	fmt.Println("Directory to be scanned: ", dir)

	// Check if the directory exists
	if _, err := os.Stat(dir[0]); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		return
	}

	// List all files of a valid image filetype in the directory
	files, err := os.ReadDir(dir[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print all files in the directory
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Create a new HTML file
	htmlFile := createHTMLPage(dir[0], files)

	// Replace with your URL
	xerr := openBrowser(fmt.Sprintf("file://%s", htmlFile))
	if xerr != nil {
		panic(err)
	}

}
