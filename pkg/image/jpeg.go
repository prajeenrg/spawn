package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
)

func MakeJpeg(filename string, image image.Image) {
	file := createFile(filename)
	defer file.Close()

	err := jpeg.Encode(file, image, &jpeg.Options{
		Quality: 76,
	})

	if err != nil {
		log.Fatalln("Cannot create jpg file from image")
	}
}

func MakeJpegs(prefix string, d *Dimens, count uint) {
	for i := uint(0); i < count; i++ {
		filename := fmt.Sprintf("%s_%7d.jpg", prefix, i)
		image := GenerateImage(d)
		MakeJpeg(filename, image)
	}
}
