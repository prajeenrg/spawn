package util

import pb "github.com/schollz/progressbar/v3"

func GetProgressBar(count uint, desp string) pb.ProgressBar {
	return *pb.Default(int64(count), desp)
}
