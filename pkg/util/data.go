package util

import (
	"crypto/rand"
	"log"
)

func GetRandomBytes(size uint) (int, []byte) {
	b := make([]byte, size)
	n, err := rand.Read(b)

	if err != nil {
		log.Fatalln("Cryptographic random data generation failed")
	}

	return n, b
}
