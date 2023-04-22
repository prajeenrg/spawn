package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"

	"github.com/prajeenrg/spawn/pkg/util"
)

func MakeJpeg(filename string, image image.Image) {
	file := util.CreateFile(filename)
	defer file.Close()

	err := jpeg.Encode(file, image, &jpeg.Options{
		Quality: 76,
	})

	if err != nil {
		log.Fatalln("Cannot create jpg file from image")
	}
}

func MakeJpegs(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	for i := uint(1); i <= count; i++ {
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.jpg", directory, prefix, d.Width, d.Height, i)
		image := GenerateImage(d)
		MakeJpeg(filename, image)
	}
}