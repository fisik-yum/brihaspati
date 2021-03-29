package main

import (
	"brihaspati/auth"
	"brihaspati/moderation"
	"brihaspati/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Moderate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if utils.StartsWith(m.Content, "^batchkick") { //allow batch kicking
		toKick := m.Mentions
		if len(toKick) == 0 {
			return
		}
		state := moderation.CheckForPerms(m.Member.Roles, m.GuildID, discordgo.PermissionKickMembers, s)
		if auth.CodeState(m.Author.ID) && state { //check if user has permissions to kick
			if moderation.BatchKick(toKick, m.GuildID, s) {
				s.ChannelMessageSend(m.ChannelID, "`Batch Kick Success!`")
			}
		}
	}

	if utils.StartsWith(m.Content, "^kick") { //allow single kicking
		toKick := m.Mentions
		if len(toKick) == 0 {
			return
		}
		state := moderation.CheckForPerms(m.Member.Roles, m.GuildID, discordgo.PermissionKickMembers, s)
		if auth.CodeState(m.Author.ID) && state { //check if user has permissions to kick
			s.GuildMemberDelete(m.GuildID, toKick[0].ID)
			s.ChannelMessageSend(m.ChannelID, ("`Kicked User " + toKick[0].Username + "`"))
		} else {
			fmt.Println("err")
		}

	}
	if utils.StartsWith(m.Content, "^ban") { //allow single ban
		toBan := m.Mentions
		if len(toBan) == 0 {
			return
		}
		fmt.Println(toBan)
		state := moderation.CheckForPerms(m.Member.Roles, m.GuildID, discordgo.PermissionBanMembers, s)
		if auth.CodeState(m.Author.ID) && state { //check if user has permissions to kick
			s.ChannelMessageSend(m.ChannelID, ("`Banned User " + toBan[0].Username + "`"))
			s.GuildBanCreate(m.GuildID, toBan[0].ID, 0)
		}
	}

	if utils.StartsWith(m.Content, "^mute") {
		state := moderation.CheckForPerms(m.Member.Roles, m.GuildID, discordgo.PermissionKickMembers, s)
		mentions := m.Mentions
		if len(mentions) == 0 {
			return
		}
		if auth.CodeState(m.Author.ID) && state {
			if moderation.Mute(m.ChannelID, mentions[0].ID, m.GuildID, s) {
				s.ChannelMessageSend(m.ChannelID, "`Mute Successful!`")
				return
			} else {
				s.ChannelMessageSend(m.ChannelID, "`Mute Unsuccessful`")
			}

		}
	}

}

//https://discord.com/channels/118456055842734083/155361364909621248/815116075552866334
// These two ^ \/ conversations helped me a lot.                                        //discord gophers server
//https://discord.com/channels/118456055842734083/155361364909621248/820247438249426954
