package golangmdtty

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Convert takes in a pathname to read and convert to terminal output. This is the only entry point into the package.
func Convert(path string) {
	fileType := strings.Split(path, ".")[1]
	if fileType != "md" {
		log.Println("Please only use a md file")
		return
	}

	isPathExists, _ := isFileExists(path)
	if !isPathExists {
		log.Println("The file to output does not even exist")
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	p := parser{scanner: *bufio.NewScanner(file)}

	// Parses through the input file and gets all the blocks
	p.getNextLine()
}

func isFileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
