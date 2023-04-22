package image

import (
	"fmt"
	"image"
	"image/png"
	"log"
)

func MakePng(filename string, image image.Image) {
	file := createFile(filename)
	defer file.Close()

	if err := png.Encode(file, image); err != nil {
		log.Fatalln("Cannot create png file from image")
	}
}

func MakePngs(prefix string, count int) {
	for i := 0; i < count; i++ {
		filename := fmt.Sprintf("%s_%07d.jpg", prefix, i)
		image := GenerateImage(1000, 1000)
		MakePng(filename, image)
	}
}
