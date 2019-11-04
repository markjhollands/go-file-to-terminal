package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify a file name")
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
