package golangmdtty

import (
	"log"
	"os"
	"strings"
)

// Run takes in a pathname to read and convert to terminal output
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
