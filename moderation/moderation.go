package moderation

import (
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
