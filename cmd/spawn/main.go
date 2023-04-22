package main

import (
	"github.com/prajeenrg/spawn/internal/image"
)

func main() {
	filename := "sample"
	image.MakePngs(filename, 10)
}
