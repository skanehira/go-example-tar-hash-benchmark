package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
}

func calcHashAndCopyFile() {
	f, err := os.Open("example.tgz")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		panic(err)
	}

	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	_ = hash

	if _, err := f.Seek(0, 0); err != nil {
		panic(err)
	}

	out, err := os.Create("example2.tgz")
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(out, f); err != nil {
		panic(err)
	}

}
