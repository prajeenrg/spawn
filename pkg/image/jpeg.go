package image

import (
	"fmt"
	"image/jpeg"
	"log"

	"github.com/prajeenrg/spawn/pkg/util"
)

type JpegGenerator struct {
	Quality int
}

func (j *JpegGenerator) SingleImage(name string, d *Dimens) {
	if !util.CheckExtension(name, "jpg") || !util.CheckExtension(name, "jpeg") {
		name = fmt.Sprintf("%s.jpg", name)
	}

	file := util.CreateFile(name)
	defer file.Close()

	image := generateImage(d)

	err := jpeg.Encode(file, image, &jpeg.Options{
		Quality: j.Quality,
	})

	if err != nil {
		log.Fatalln("Cannot create jpg file from image")
	}
}

func (j *JpegGenerator) MultipleImages(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	bar := util.GetProgressBar(count, "Generating JPEG files")
	for i := uint(1); i <= count; i++ {
		util.Increment(&bar)
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.jpg", directory, prefix, d.Width, d.Height, i)
		j.SingleImage(filename, d)
	}
}
