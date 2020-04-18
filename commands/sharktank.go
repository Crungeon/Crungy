package commands

import "github.com/bwmarrin/discordgo"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func SharkTank(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "shark tank" {
		s.ChannelMessageSend(m.ChannelID, "Mike has learned alot of valuable business information on shark tank primarily the following.")
		s.ChannelMessageSend(m.ChannelID, "1) Has a proper valuation based on current business not potential business or you look like an idiot.")
		s.ChannelMessageSend(m.ChannelID, "2) Building off #1 don't assume your shitty start up is going to get 20% market share of anything in a year.")
		s.ChannelMessageSend(m.ChannelID, "3) Have a plan for what you're doing with the investor's money and be able to articulate what that plan is.")
		s.ChannelMessageSend(m.ChannelID, "4) Patents are important. If you don't have one don't bother showing up. Also a registered design is not a patent you dumb dumb.")
		s.ChannelMessageSend(m.ChannelID, "5) Don't claim your shitty idea has medical benefits if you haven't done a clinical trial. It makes you look like a dingus who's about to get sued.")
	}

}
