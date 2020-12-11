package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/thundagax/discord-miao/lib/rapid"
)

var BotID string
var goBot *discordgo.Session
var creatorID string

func main() {
	creatorID = os.Getenv("CREATOR_DISCORD_ID")
	token := os.Getenv("BOT_TOKEN")
	goBot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error Init Discord")
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println(m.Author.ID, m.Author.Username)
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "!me" {
		handleMeMessage(s, m.ChannelID, m.Author.ID)
	} else if m.Content == "!help" {
		handleHelpMessage(s, m.ChannelID, m.Author.ID)
	} else if m.Content == "!pet" {
		handlePetMessage(s, m.ChannelID, m.Author.ID)
	} else if m.Content == "!jokes" {
		handleJokesMessage(s, m.ChannelID, m.Author.ID)
	} else if strings.Contains(m.Content, "!match") {
		handleLoveCalculator(s, m)
	} else if strings.Contains(m.Content, "!bacotin") {

	}
}

func handleMeMessage(s *discordgo.Session, channelID string, authorID string) {
	if authorID == creatorID {
		worship, _ := os.Open("./resources/worship.gif")
		s.ChannelMessageSend(channelID, "yes kami-sama?")
		s.ChannelFileSend(channelID, "worship.gif", worship)
	} else {
		slap, _ := os.Open("./resources/slap.gif")
		s.ChannelMessageSend(channelID, "shut up hooman!?")
		s.ChannelFileSend(channelID, "slap.gif", slap)
	}
}

func handlePetMessage(s *discordgo.Session, channelID string, authorID string) bool {
	time := time.Now().Minute()

	if time < 15 || (time > 30 && time < 45) {
		s.ChannelMessageSend(channelID, "*angry noises*")
		return false
	} else {
		s.ChannelMessageSend(channelID, "*meow~*")
		return true
	}
}

func handleHelpMessage(s *discordgo.Session, channelID string, authorID string) {
	s.ChannelMessageSend(channelID, "Here is on the menu hooman.")
	s.ChannelMessageSend(channelID, `> "!me" to see who you are in my opinion`)
	s.ChannelMessageSend(channelID, `> "!pet" to check if i am in the mood`)
	s.ChannelMessageSend(channelID, `> "!jokes" i have tons of dad jokes`)
	s.ChannelMessageSend(channelID, `> "!match" mention someone, and i will give you my opinion about you being couple`)
}

func handleJokesMessage(s *discordgo.Session, channelID string, authorID string) {
	res := rapid.GetRandomJoke()
	if res == nil {
		s.ChannelMessageSend(channelID, "The limit has reached today hooman. ITS BECAUSE YOU USE THE FREE PLAN !!")
	} else {
		s.ChannelMessageSend(channelID, res.Setup)
		s.ChannelMessageSend(channelID, "> "+res.Punchline)
	}
}

func handleLoveCalculator(s *discordgo.Session, content *discordgo.MessageCreate) {
	mentions := content.Mentions
	log.Println(mentions)
	if len(mentions) < 1 {
		s.ChannelMessageSend(content.ChannelID, "Please mention the couple you want to match miao~ ")
		return
	} else if len(mentions) >= 2 {
		s.ChannelMessageSend(content.ChannelID, "Please mention one person only miao~ ")
		return
	} else if mentions[0].ID == BotID {
		s.ChannelMessageSend(content.ChannelID, "Miao dont like hooman ?!")
		return
	}

	owner := content.Author.Username
	target := content.Mentions[0].Username

	ownerId := content.Author.ID
	targetID := content.Mentions[0].ID

	total := 0
	for i := 0; i < len(owner); i++ {
		total = total + int(owner[i])
	}

	for i := 0; i < len(target); i++ {
		total = total + int(target[i])
	}

	percentage := total % 100

	s.ChannelMessageSend(content.ChannelID, fmt.Sprintf("<@%s> are %d%s compatible with <@%s>", ownerId, percentage, "%", targetID))
	if percentage < 30 {
		s.ChannelMessageSend(content.ChannelID, "> Haiya, u both sucks together miao~")
	} else if percentage < 60 {
		s.ChannelMessageSend(content.ChannelID, "> Hmm not bad miao~")
	} else if percentage < 90 {
		s.ChannelMessageSend(content.ChannelID, "> Hmmm i see lovebirds <3")
	} else {
		s.ChannelMessageSend(content.ChannelID, "> Send me those wedding invitation miao~")
	}
}
