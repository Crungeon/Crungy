package commands

import (
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func Love(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if matched, _ := regexp.MatchString("^luke time", m.Content); matched {
		re := regexp.MustCompile(`[0-9]{2}`)
		minstr := string(re.Find([]byte(m.Content)))
		minraw, _ := strconv.Atoi(minstr)

		minact = minraw*2 + 5

		if minact > (minraw + 45) {
			minact = (minraw + 45)
		}

		message := "Luke will arrive in" + strconv.Itoa(minact) + "min."
		s.ChannelMessageSend(m.ChannelID, message)
	}
}
