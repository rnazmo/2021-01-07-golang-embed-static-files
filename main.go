package main

import (
	"io"
	"log"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/rnazmo/sandbox__go_statik/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/hello.sh")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}
