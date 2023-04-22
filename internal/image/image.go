package image

import (
	"crypto/rand"
	"image"
	"image/color"
	"log"
	"os"
)

type Dimens struct {
	width  int
	height int
}

func GenerateImage(d *Dimens) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, d.width, d.height))
	b := make([]byte, d.width*d.height*3)
	n, err := rand.Read(b)

	if err != nil {
		log.Fatalln("Random pixel generation failed")
	}

	for x := 0; x < d.width; x++ {
		for y := 0; y < d.height; y++ {
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

func createFile(filename string) *os.File {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Cannot create file: %s\n", filename)
	}

	return file
}
