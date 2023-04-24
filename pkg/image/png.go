package image

import (
	"fmt"
	"image/png"
	"log"

	"github.com/prajeenrg/spawn/pkg/util"
)

type PngGenerator struct{}

func (p *PngGenerator) SingleImage(filename string, d *Dimens) {
	file := util.CreateFile(filename)
	defer file.Close()

	image := generateImage(d)

	if err := png.Encode(file, image); err != nil {
		log.Fatalln("Cannot create png file from image")
	}
}

func (p *PngGenerator) MultipleImages(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	bar := util.GetProgressBar(count, "Generating PNG files")
	for i := uint(1); i <= count; i++ {
		util.Increment(&bar)
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.png", directory, prefix, d.Width, d.Height, i)
		p.SingleImage(filename, d)
	}
}
