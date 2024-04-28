package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {

	dir := os.Args[1:]

	fmt.Println("Directory to be scanned: ", dir)

	// Check if the directory exists
	if _, err := os.Stat(dir[0]); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		return
	}

	// List all files in the directory
	files, err := os.ReadDir(dir[0])
	if err != nil {
		fmt.Println("Error reading directory")
		return
	}

	// Print all files in the directory
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Replace with your URL
	xerr := openBrowser("http://www.google.com")
	if xerr != nil {
		panic(err)
	}

}
