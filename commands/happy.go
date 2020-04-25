package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func FindUser(users []*discordgo.User, user *discordgo.User) bool {
	for _, u := range users {
		if u.ID == user.ID {
			return true
		}
	}
	return false
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func Happy(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !FindUser(m.Mentions, s.State.User) {
		return
	}

	if !strings.Contains("make me happy", strings.ToLower(m.Content)) {
		return
	}

	happyThoughts := [10]string{
		"You're beautiful!",
		"I like your butt.",
		"You're smart af!",
		"Your friends think you're great.",
		"You're really good at what you do!",
		"I think about you when I touch myself.",
		"Made you look. Oops, that wasn't nice. Fuck you, be happy.",
		"If everyone in the world was like you, I'd probably kill myself because I'd feel so bad in comparison. You're too great.",
		"You're amazing!",
		"I love you.",
	}

	rand.Seed(time.Now().Unix())
	s.ChannelMessageSend(m.ChannelID, happyThoughts[rand.Intn(len(happyThoughts))])
}
