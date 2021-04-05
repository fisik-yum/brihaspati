package main

import (
	"brihaspati/auth"
	"brihaspati/utils"

	"github.com/bwmarrin/discordgo"
)

func ListenForAction(s *discordgo.Session, m *discordgo.MessageCreate) { //temporarily exported
	if m.Content == "^getcode" {

		if utils.ChannelInGuild(m.ChannelID, m.GuildID, s) {
			s.ChannelMessageSend(m.ChannelID, "`DMed your code!`")
		}
		dmChannel, _ := s.UserChannelCreate(m.Author.ID)
		_, _ = s.ChannelMessageSend(dmChannel.ID, "`Code "+auth.CreateCode(m.Author.ID)+"`")
		return
	}

	if utils.Prefix(m.Content, "^authcode") {
		dmChannel, _ := s.UserChannelCreate(m.Author.ID)

		if m.ChannelID != dmChannel.ID {
			s.ChannelMessageSend(m.ChannelID, "`Verify in your DMs!`")
			return
		} else if auth.ValidateCode(m.Author.ID, m.Content) {
			dmChannel, _ := s.UserChannelCreate(m.Author.ID)
			_, _ = s.ChannelMessageSend(dmChannel.ID, "`Verified!`")
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		}
	}
}
