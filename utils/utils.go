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
	if len(strings.Fields(target)) == 0 {
		return false
	}
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

func Prefix(content string, prefix string) bool {
	if len(content) <= len(prefix) {
		return false
	} else if content[0:len(prefix)] == prefix {
		return true
	}
	return false
}

func IsMentioned(id string, mentions []*discordgo.User) bool {
	for x := range mentions {
		if mentions[x].ID == id {
			return true
		}
	}
	return false
}
