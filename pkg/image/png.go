package image

import (
	"fmt"
	"image"
	"image/png"
	"log"

	"github.com/prajeenrg/spawn/pkg/util"
)

func MakePng(filename string, image image.Image) {
	file := util.CreateFile(filename)
	defer file.Close()

	if err := png.Encode(file, image); err != nil {
		log.Fatalln("Cannot create png file from image")
	}
}

func MakePngs(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	bar := util.GetProgressBar(count, "Generating PNG files")
	for i := uint(1); i <= count; i++ {
		bar.Add(1)
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.png", directory, prefix, d.Width, d.Height, i)
		image := GenerateImage(d)
		MakePng(filename, image)
	}
}
