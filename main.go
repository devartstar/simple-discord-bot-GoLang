package main

import (
	"fmt"

	"github.com/devartstar/discordBot/bot"
	"github.com/devartstar/discordBot/config"
)

func main() {
	// funciton in another file
	err := config.ReadConfig()
	
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	// functio in another file to start
	bot.Start()

	<-make(chan struct{})
	return
}