package image

import (
	"fmt"
	"log"

	"github.com/carck/libheif/go/heif"
	"github.com/prajeenrg/spawn/pkg/util"
)

type HeifGenerator struct {
	Quality int
}

func (h *HeifGenerator) SingleImage(name string, d *Dimens) {
	image := generateImage(d)
	ctx, err := heif.EncodeFromImage(image, heif.CompressionHEVC, h.Quality, heif.LosslessModeEnabled, heif.LoggingLevelFull)

	if err != nil {
		log.Fatalln("Cannot create heif file from image")
	}

	if !util.CheckExtension(name, "heic") || !util.CheckExtension(name, "heif") {
		name = fmt.Sprintf("%s.heic", name)
	}

	if err := ctx.WriteToFile(name); err != nil {
		log.Fatalln("Failed to write heif image")
	}
}

func (h *HeifGenerator) MultipleImages(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	bar := util.GetProgressBar(count, "Generating heif files")
	for i := uint(1); i <= count; i++ {
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.heic", directory, prefix, d.Width, d.Height, i)
		h.SingleImage(filename, d)
		util.Increment(&bar)
	}
}
