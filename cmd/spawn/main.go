package main

import (
	"log"
	"os"

	"github.com/prajeenrg/spawn/pkg/image"
	"github.com/urfave/cli/v2"
)

var version = "latest"

func main() {
	app := cli.NewApp()
	defer app.Run(os.Args)

	app.Name = "spawn"
	app.Usage = "generate dummy data files for testing purposes"
	app.Authors = []*cli.Author{{Name: "Prajeen Govardhanam", Email: "prajeenrg@gmail.com"}}
	app.Version = version
	app.UseShortOptionHandling = true
	app.HideHelpCommand = true
	app.Copyright = "Copyright 2023 Prajeen Govardhanam"
	app.Commands = []*cli.Command{
		{
			Name:  "image",
			Usage: "used to generate dummy image files",
			Flags: []cli.Flag{
				&cli.UintFlag{Name: "width", Aliases: []string{"iw"}, Value: 500},
				&cli.UintFlag{Name: "height", Aliases: []string{"ih"}, Value: 500},
				&cli.UintFlag{Name: "count", Aliases: []string{"c"}, Value: 1},
				&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Value: "png"},
				&cli.StringFlag{Name: "prefix", Aliases: []string{"p"}, Value: "spawn"},
				&cli.StringFlag{Name: "directory", Aliases: []string{"d"}, Value: "."},
			},
			Action: func(ctx *cli.Context) error {
				prefix := ctx.String("prefix")
				dir := ctx.String("directory")
				dimen := &image.Dimens{Width: ctx.Uint("width"), Height: ctx.Uint("height")}
				imgType := ctx.String("type")
				count := ctx.Uint("count")
				switch imgType {
				case "png":
					image.MakePngs(dir, prefix, dimen, count)
				case "jpg", "jpeg":
					image.MakeJpegs(dir, prefix, dimen, count)
				default:
					log.Fatalf("Invalid image mime type '%s' used", imgType)
				}
				return nil
			},
		},
	}
}
