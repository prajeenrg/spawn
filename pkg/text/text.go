package text

import (
	"log"

	"github.com/prajeenrg/spawn/pkg/util"
)

func MakeDummyFile(name string, size uint) {
	file := util.CreateFile(name)
	defer file.Close()

	_, b := util.GetRandomBytes(size)
	_, err := file.Write(b)

	if err != nil {
		log.Fatalf("Writing to %s failed", name)
	}
}
