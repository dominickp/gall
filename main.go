package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func createHTMLGallery(template, directoryAbsolutePath string, images []fs.DirEntry) string {
	// Create a new file in the directory
	galleryFileName := filepath.Join(directoryAbsolutePath, "gal.html")
	file, err := os.Create(galleryFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	galleryContents := ""
	for _, image := range images {
		galleryContents += fmt.Sprintf(`
		<figure class='card'>
			<a href='%s' target='_blank'>
				<img src='%s' />
			</a>
		</figure>`, image.Name(), image.Name())
	}

	// Find the placeholder "<!-- GALLERY_CONTENTS -->" and replace it in the template
	template = strings.Replace(template, "<!-- GALLERY_CONTENTS -->", galleryContents, 1)

	file.WriteString(template)

	log.Printf("Gallery created: %s", galleryFileName)

	return galleryFileName
}

//go:embed template.html
var template string

func main() {

	dir := os.Args[1:][0]
	log.Printf("Directory to be scanned: %s", dir)

	// Check if the directory exists
	dirInfo, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return
	}
	if !dirInfo.IsDir() {
		log.Fatalf("'%s' is not a directory", dir)
		return
	}

	dir, err = filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Error resolving absolute path: %s", err)
		return
	}

	// List all files of a valid image filetype in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
		return
	}
	images := []fs.DirEntry{}

	// Print all files in the directory
	for _, file := range files {
		if fileIsImage(file) {
			images = append(images, file)
		}
	}

	log.Printf("Found %d images", len(images))

	// Create a new HTML file
	galleryFileName := createHTMLGallery(template, dir, images)

	// FIXME: open browser conditionally
	err = openBrowser(fmt.Sprintf("file://%s", galleryFileName))
	if err != nil {
		panic(err)
	}

}
