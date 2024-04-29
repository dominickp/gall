package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Randomize            bool `short:"r" long:"randomize" description:"Randomize image ordering in the gallery"`
	LaunchDefaultBrowser bool `short:"b" long:"browser" description:"Launch the default browser after creating the gallery"`
	LaunchFirefox        bool `short:"f" long:"firefox" description:"Launch Firefox after creating the gallery"`
}

func createHTMLGallery(template, directoryAbsolutePath string, images []fs.DirEntry) string {
	// Create a new file in the directory
	galleryFileName := filepath.Join(directoryAbsolutePath, "gal.html")
	file, err := os.Create(galleryFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	galleryTitle := filepath.Base(directoryAbsolutePath)

	galleryContents := ""
	for _, image := range images {
		galleryContents += fmt.Sprintf(`
		<figure class='card'>
			<img src='%s' />
		</figure>`, image.Name())
	}

	// Replace placeholders in the template
	template = strings.Replace(template, "<!-- GALLERY_CONTENTS -->", galleryContents, 1)
	template = strings.Replace(template, "<!-- GALLERY_TITLE -->", galleryTitle, 1)

	file.WriteString(template)

	log.Printf("Gallery created: %s", galleryFileName)

	return galleryFileName
}

//go:embed template.html
var template string

func main() {

	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}

	// log.Printf("Randomize: %v\n", opts.Randomize)
	// log.Printf("LaunchDefaultBrowser: %t\n", opts.LaunchDefaultBrowser)
	// log.Printf("LaunchFirefox: %t\n", opts.LaunchFirefox)

	// Get the target directory
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("Usage: %s <directory>", os.Args[0])
		return
	}
	dir := args[0]
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

	// Gather images in the target directory
	images := getImagesInDirectory(dir)

	// Create a new HTML file
	galleryFileName := createHTMLGallery(template, dir, images)

	// Open the gallery in the browser
	if opts.LaunchDefaultBrowser || opts.LaunchFirefox {
		var browser string
		if opts.LaunchFirefox {
			browser = "firefox"
		}
		if opts.LaunchDefaultBrowser {
			browser = "default"
		}

		err = openBrowser(fmt.Sprintf("file://%s", galleryFileName), browser)
		if err != nil {
			panic(err)
		}
	}

}
