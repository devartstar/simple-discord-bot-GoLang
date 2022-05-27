package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/devartstar/discordBot/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	// creating an instance of bot
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	// logic of handeling the inputs
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignores the text enterd by bot
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}