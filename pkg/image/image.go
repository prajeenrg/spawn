package image

import (
	"crypto/rand"
	"image"
	"image/color"
	"log"
)

type Dimens struct {
	Width  uint
	Height uint
}

func GenerateImage(d *Dimens) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, int(d.Width), int(d.Height)))
	b := make([]byte, d.Width*d.Height*3)
	n, err := rand.Read(b)

	if err != nil {
		log.Fatalln("Random pixel generation failed")
	}

	for x := 0; x < int(d.Width); x++ {
		for y := 0; y < int(d.Height); y++ {
			img.Set(x, y, color.NRGBA{
				R: b[n-1],
				G: b[n-2],
				B: b[n-3],
				A: 100,
			})
			n -= 3
		}
	}

	return img
}
