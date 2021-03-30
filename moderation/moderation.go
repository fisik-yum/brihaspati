package moderation

import (
	"brihaspati/roles"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CheckForPerms(roles []string, guildID string, cInt int, s *discordgo.Session) bool {
	for x := 0; x < len(roles); x++ {
		pInt, err := s.State.Role(guildID, roles[x])
		if err != nil {
			fmt.Println("ERR getting role")
			fmt.Println(err)

			return false
		}
		if (pInt.Permissions)&int64(cInt) == int64(cInt) {
			return true
		}
	}
	fmt.Println("ERR with CHECK")
	return false
}

func BatchKick(toKick []*discordgo.User, guildID string, s *discordgo.Session) bool {
	for x := 0; x < len(toKick); x++ { //recursion to kick all mentioned users
		s.GuildMemberDelete(guildID, toKick[x].ID)
	}
	return true
}

func Mute(ChannelID, userID, gID string, s *discordgo.Session) bool {
	role, err := s.GuildRoles(gID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if roles.CheckIfMuteExists(gID, role) {
		channels, err := s.GuildChannels(gID)
		if err != nil {
			fmt.Println(err)
			return false
		}
		roles.ApplyChannelOverrides(roles.ReadItem(gID, 1), channels, s)
		s.GuildMemberRoleAdd(gID, userID, roles.ReadItem(gID, 1))
		return true
	} else {
		fmt.Println("couldn't find role")
		if roles.CreateMuteRole(gID, s) {
			channels, err := s.GuildChannels(gID)
			if err != nil {
				fmt.Println(err)
				return false
			}
			roles.ApplyChannelOverrides(roles.ReadItem(gID, 1), channels, s)
			s.GuildMemberRoleAdd(gID, userID, roles.ReadItem(gID, 1))
			return true
		}
		return false

	}
}

func Unmute(ChannelID, userID, gID string, s *discordgo.Session) bool {
	role, err := s.GuildRoles(gID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if roles.CheckIfMuteExists(gID, role) {
		s.GuildMemberRoleRemove(gID, userID, roles.ReadItem(gID, 1))
		return true
	} else {
		fmt.Println("couldn't find role")
		return false

	}
}
