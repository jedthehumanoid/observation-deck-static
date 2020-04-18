package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

func loadToml(file string, i interface{}) error {
	d, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(d), i)
	return err
}

func toJSON(in interface{}) string {
	b, err := json.MarshalIndent(in, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
