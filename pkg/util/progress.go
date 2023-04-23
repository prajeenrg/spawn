package util

import (
	"log"

	pb "github.com/schollz/progressbar/v3"
)

func GetProgressBar(count uint, desp string) pb.ProgressBar {
	return *pb.Default(int64(count), desp)
}

func Increment(bar *pb.ProgressBar) {
	if err := bar.Add(1); err != nil {
		log.Println("Progress bar update failed")
	}
}
