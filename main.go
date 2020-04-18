package main

import (
	"fmt"
	"github.com/jedthehumanoid/card-cabinet"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Src string `toml:"src"`
}

func loadConfig(file string) Config {
	var config Config
	err := loadToml(file, &config)
	if err != nil {
		fmt.Println("Couldn't load configuration file")
	}
	return config
}

func main() {

	config := loadConfig("cabinet.toml")

	if config.Src == "" {
		config.Src = "."
	}
	config.Src = filepath.Clean(config.Src) + "/"

	cards, boards := cardcabinet.ReadDir(config.Src)

	output := "let data = {\n"

	output += "cards: " + toJSON(cards) + ",\n"
	output += "boards: " + toJSON(boards) + "\n"
	output += "}\n"
	ioutil.WriteFile("data.js", []byte(output), 0644)
}
