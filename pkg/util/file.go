package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckExtension(name, ext string) bool {
	return strings.HasSuffix(name, fmt.Sprintf(".%s", ext))
}

func CreateFolderIfNotExits(name string) {
	if len(name) == 0 {
		name = "."
	}
	if _, err := os.Stat(name); os.IsNotExist(err) {
		log.Printf("Creating directory '%s' since not present\n", name)
		if err := os.Mkdir(name, os.FileMode(0755)); err != nil {
			log.Fatalf("Directory '%s' creation failed\n", name)
		}
	}
}

func CreateFile(name string) *os.File {
	file, err := os.Create(name)

	if err != nil {
		log.Fatalf("Cannot create file: %s\n", name)
	}

	return file
}
