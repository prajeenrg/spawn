package image

import (
	"fmt"
	"log"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/prajeenrg/spawn/pkg/util"
)

type WebpGenerator struct {
	Quality float32
}

func (w *WebpGenerator) generateOptions() *encoder.Options {
	opts, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, w.Quality)
	if err != nil {
		log.Fatalln("WebP encoder options failed")
	}
	return opts
}

func (w *WebpGenerator) SingleImage(name string, d *Dimens) {
	if !util.CheckExtension(name, "webp") {
		name = fmt.Sprintf("%s.webp", name)
	}

	file := util.CreateFile(name)
	defer file.Close()

	image := generateImage(d)
	opts := w.generateOptions()

	if err := webp.Encode(file, image, opts); err != nil {
		log.Fatalln("Cannot create webp file from image")
	}
}

func (w *WebpGenerator) MultipleImages(directory, prefix string, d *Dimens, count uint) {
	util.CreateFolderIfNotExits(directory)
	bar := util.GetProgressBar(count, "Generating WebP files")
	for i := uint(1); i <= count; i++ {
		filename := fmt.Sprintf("%s/%s_%dx%d_%d.webp", directory, prefix, d.Width, d.Height, i)
		w.SingleImage(filename, d)
		util.Increment(&bar)
	}
}
