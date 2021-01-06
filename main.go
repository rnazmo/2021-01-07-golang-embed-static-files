package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/rakyll/statik/fs"
	_ "github.com/rnazmo/sandbox__go_statik/statik"
)

type Config struct {
	Foo Foo `toml:"foo"`
	Bar Bar `toml:"bar"`
}

type Foo struct {
	Words string `toml:"words"`
}
type Bar struct {
	Number int `toml:"number"`
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/default.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conf := new(Config)
	_, err = toml.DecodeReader(f, conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
	fmt.Println(conf.Foo)
	fmt.Println(conf.Foo.Words)
}
