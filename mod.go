package main

import (
	"brihaspati/auth"
	"brihaspati/moderation"
	"brihaspati/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Moderate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if utils.Contains(strings.Fields(m.Content), "^kick") {
		toKick := m.Mentions
		fmt.Println(toKick)
		state := moderation.CheckForPerms(m.Member.Roles, m.GuildID, discordgo.PermissionKickMembers, s)
		if auth.CodeState(m.Author.ID) && state { //check if user has permissions to kick
			s.ChannelMessageSend(m.ChannelID, "`Success!`")
			for x := 0; x < len(toKick); x++ {
				s.GuildMemberDelete(m.GuildID, toKick[x].ID)
			}
		}
	}
}

//https://discord.com/channels/118456055842734083/155361364909621248/815116075552866334
// These two ^ \/ links helped me a lot.
//https://discord.com/channels/118456055842734083/155361364909621248/820247438249426954
