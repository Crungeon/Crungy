package commands

import "github.com/bwmarrin/discordgo"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func Roasted(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "gottem" {
		s.ChannelMessageSend(m.ChannelID, ":boom: Roasted!")
	}

}