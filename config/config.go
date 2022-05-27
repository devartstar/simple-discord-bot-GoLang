package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string
	config *configStruct // user degined datatype
)

// defining the data type (to read data from config.json) as struct
type configStruct struct {
	Token string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// function to read the config file from ../config.json
func ReadConfig() error{
	fmt.Println("Reading config file...")

	// ioutil to read a file
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	fmt.Println(string(file))

	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println((err.Error()))
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}