package util

import (
	"log"
	"os"
)

func CreateFolderIfNotExits(name string) {
	if _, err := os.Stat(name); os.IsExist(err) {
		return
	}

	log.Printf("Creating directory '%s' since not present\n", name)
	if err := os.Mkdir(name, os.ModeDir); err != nil {
		log.Fatalf("Directory '%s' creation failed\n", name)
	}
}

func CreateFile(name string) *os.File {
	file, err := os.Create(name)

	if err != nil {
		log.Fatalf("Cannot create file: %s\n", name)
	}

	return file
}
