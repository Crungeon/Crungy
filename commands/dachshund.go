package commands

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

type dogApiRandom struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func Dachshund(s *discordgo.Session, m *discordgo.MessageCreate) {

	client := resty.New()
	response := &dogApiRandom{}
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	// https://dog.ceo/api/breed/dachshund/images/random
	if m.Author.ID == s.State.User.ID {
		return
	}

	resp, err := client.R().
		EnableTrace().
		Get("https://dog.ceo/api/breed/dachshund/images/random")
	if err != nil {
		log.Error(err)
	}

	json.Unmarshal(resp.Body(), response)
	if strings.ToLower(m.Content) == "wiener me" {
		s.ChannelMessageSend(m.ChannelID, response.Message)
	}
}