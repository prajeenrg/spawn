package image

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/prajeenrg/spawn/pkg/util"
)

type Dimens struct {
	Width  uint
	Height uint
}

type Generator interface {
	SingleImage(string, *Dimens)
	MultipleImages(string, string, *Dimens, uint)
}

func convertNrgbaToRgba(im image.Image) image.Image {
	b := im.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(rgba, rgba.Bounds(), im, b.Min, draw.Src)
	return rgba
}

func generateImage(d *Dimens) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, int(d.Width), int(d.Height)))
	n, b := util.GetRandomBytes(d.Width * d.Height * 3)

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
