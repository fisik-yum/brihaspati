package main

import (
	"brihaspati/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func funStuff(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "^rolldice" {
		rand.Seed(time.Now().UnixNano())
		s.ChannelMessageSend(m.ChannelID, "`"+strconv.Itoa(rand.Intn(6)))

	}

	if utils.Prefix(m.Content, "^ghostping") {
		if len(m.Mentions) != 0 {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		} else {
			s.ChannelMessageSend(m.ChannelID, "`Mention someone to use this function!`")
		}
	}

	if utils.IsMentioned(s.State.User.ID, m.Mentions) {
		s.ChannelMessageSend(m.ChannelID, "`I HAVE BEEN SUMMONED`")
	}
}
