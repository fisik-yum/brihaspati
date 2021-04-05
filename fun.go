package main

import (
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
}
