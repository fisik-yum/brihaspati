package utils

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StartsWith(target string, command string) bool {
	return (strings.Fields(target)[0] == command)

}

func ChannelInGuild(channelID, guildID string, s *discordgo.Session) bool {

	// Get channels for this guild
	channels, _ := s.GuildChannels(guildID)
	for x := 0; x < len(channels); x++ {
		if channels[x].ID == channelID {
			return true
		}
	}
	return false
}
