package main

import (
	"log"
	"os"

	"github.com/prajeenrg/spawn/pkg/file"
	"github.com/prajeenrg/spawn/pkg/image"
	"github.com/urfave/cli/v2"
)

var version = "latest"

func main() {
	app := cli.NewApp()

	app.Name = "spawn"
	app.Usage = "generate dummy data files for testing purposes"
	app.Authors = []*cli.Author{{Name: "Prajeen Govardhanam", Email: "prajeenrg@gmail.com"}}
	app.Version = version
	app.UseShortOptionHandling = true
	app.HideHelpCommand = true
	app.Copyright = "Copyright 2023 Prajeen Govardhanam"
	app.Commands = []*cli.Command{
		fileCmd(),
		imageCmd(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Println("Failed initializing cli app")
		log.Fatalln(err)
	}
}

func imageCmd() *cli.Command {
	image := &cli.Command{
		Name:  "image",
		Usage: "generate dummy images",
		Flags: []cli.Flag{
			&cli.UintFlag{Name: "width", Aliases: []string{"iw"}, Value: 500},
			&cli.UintFlag{Name: "height", Aliases: []string{"ih"}, Value: 500},
			&cli.UintFlag{Name: "count", Aliases: []string{"c"}, Value: 1},
			&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Value: "png"},
			&cli.StringFlag{Name: "prefix", Aliases: []string{"p"}, Value: "spawn"},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "spawn"},
			&cli.StringFlag{Name: "directory", Aliases: []string{"d"}},
			&cli.UintFlag{Name: "quality", Aliases: []string{"q"}, Value: 76},
		},
		Action: func(ctx *cli.Context) error {
			count := ctx.Uint("count")
			if count == 0 {
				return nil
			}

			dimen := &image.Dimens{Width: ctx.Uint("width"), Height: ctx.Uint("height")}
			imgType := ctx.String("type")
			q := ctx.Uint("quality")
			if q > 100 {
				q = 100
			}

			var generator image.Generator

			switch imgType {
			case "png":
				generator = &image.PngGenerator{}
			case "jpg", "jpeg":
				generator = &image.JpegGenerator{Quality: int(q)}
			case "webp":
				generator = &image.WebpGenerator{Quality: float32(q)}
			case "heic":
				generator = &image.HeifGenerator{Quality: int(q)}
			default:
				log.Fatalf("Invalid image mime type '%s' used", imgType)
			}

			if count == 1 {
				name := ctx.String("name")
				generator.SingleImage(name, dimen)
			} else {
				prefix := ctx.String("prefix")
				dir := ctx.String("directory")
				generator.MultipleImages(dir, prefix, dimen, count)
			}

			return nil
		},
	}
	return image
}

func fileCmd() *cli.Command {
	text := &cli.Command{
		Name:  "file",
		Usage: "generate dummy file",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
			&cli.UintFlag{Name: "size", Aliases: []string{"s"}, Value: 100},
		},
		Action: func(ctx *cli.Context) error {
			name := ctx.String("name")
			size := ctx.Uint("size")
			file.MakeDummyFile(name, size)
			return nil
		},
	}
	return text
}
