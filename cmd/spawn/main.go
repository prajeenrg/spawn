package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	defer app.Run(os.Args)

	app.Name = "spawn"
	app.Usage = "generate dummy data files for testing"
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{{
		Name:    "image",
		Aliases: []string{"i"},
		Usage:   "used to generate dummy image files",
		Flags: []cli.Flag{
			&cli.UintFlag{Name: "width", Aliases: []string{"iw"}, Value: 500},
			&cli.UintFlag{Name: "height", Aliases: []string{"ih"}, Value: 500},
			&cli.UintFlag{Name: "count", Aliases: []string{"c"}, Value: 1},
			&cli.StringFlag{Name: "type", Aliases: []string{"t"}, Value: "png"},
			&cli.StringFlag{Name: "prefix", Aliases: []string{"p"}, Value: "spawn"},
			&cli.StringFlag{Name: "directory", Aliases: []string{"d"}},
		},
		Action: func(ctx *cli.Context) error {
			switch v := ctx.String("type"); v {
			case "png":
				log.Println("PNG type is used")
			case "jpg", "jpeg":
				log.Println("JPG type is used")
			default:
				log.Fatalf("Invalid image mime type '%s' used", v)
			}
			return nil
		},
	}}
}
