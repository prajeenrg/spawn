package image

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type Dimens struct {
	width  int
	height int
}

func GenerateImage(width, height int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	b := make([]byte, width*height*3)
	n, err := rand.Read(b)

	if err != nil {
		log.Fatalln("Random pixel generation failed")
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
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

func MakeJpegs(prefix string, count int) {
	for i := 0; i < count; i++ {
		filename := fmt.Sprintf("%s_%7d.jpg", prefix, i)
		image := GenerateImage(1000, 1000)
		MakeJpeg(filename, image)
	}
}

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
