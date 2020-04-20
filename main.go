package main

import (
	"fmt"
	"github.com/jedthehumanoid/card-cabinet"
	"io/ioutil"
	"path/filepath"
	"strings"
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

// load config
// getfiles
// get cards from files
// get boars from files
// remove path from cards
// remove path from boards
// --for all decks add path from board

func main() {

	config := loadConfig("cabinet.toml")

	if config.Src == "" {
		config.Src = "."
	}
	config.Src = filepath.Clean(config.Src) + "/"

	files := cardcabinet.FindFiles(config.Src)

	cards := cardcabinet.ReadCards(files)
	boards := cardcabinet.ReadBoards(files)

	for i, _ := range cards {
		cards[i].Title = strings.TrimPrefix(cards[i].Title, config.Src)
	}

	for i, _ := range boards {
		boards[i].Title = strings.TrimPrefix(boards[i].Title, config.Src)
	}

	// Add board dir for each card in deck
	for i, board := range boards {
		boarddir := cardcabinet.GetPath(board.Title)
		for j, deck := range board.Decks {
			for k, card := range deck.Cards {
				boards[i].Decks[j].Cards[k] = boarddir + card
			}
		}
	}

	for i, board := range boards {
		boards[i] = board.Get(cards)
	}

	labels := cardcabinet.GetLabels(cards)
	folders := cardcabinet.GetFolders(cards)

	output := "let data = {\n"

	output += "cards: " + toJSON(cards) + ",\n"
	output += "boards: " + toJSON(boards) + ",\n"
	output += "folders: " + toJSON(folders) + ",\n"
	output += "labels: " + toJSON(labels) + "\n"
	output += "}\n"
	ioutil.WriteFile("data.js", []byte(output), 0644)
}
