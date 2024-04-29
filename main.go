package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	flags "github.com/jessevdk/go-flags"
	minify "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

type Options struct {
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
	galleryInfo := fmt.Sprintf("Created on %s - %d images", time.Now().Format("January 2, 2006 at 3:04 PM"), len(images))

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
	template = strings.Replace(template, "<!-- GALLERY_INFO -->", galleryInfo, 1)

	// Minify the HTML
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	template, err = m.String("text/html", template)
	if err != nil {
		panic(err)
	}

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
