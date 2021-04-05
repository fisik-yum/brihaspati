package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func LoadColors() ColorData {
	colorN, err := os.Open(filepath.Join(".", "colorN.csv"))

	if err != nil {
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}

	colorD, err := os.Open(filepath.Join(".", "colorD.csv"))

	if err != nil {
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}
	r := csv.NewReader(colorN)
	names, err := r.Read()

	if err == io.EOF {
		log.Fatal(err)
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}

	r = csv.NewReader(colorD)
	values, err := r.Read()

	if err == io.EOF || (len(names) != len(values)) {
		log.Fatal(err)
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}

	m := make(map[string]int)
	for x := range names { //load all colors into a map
		val, err := strconv.Atoi(values[x])
		if err != nil {
			break
		}
		m[names[x]] = val
	}
	x := ColorData{
		Colors: m,
		State:  true,
	}
	return x
}

//define data structures below this point
type ColorData struct {
	Colors map[string]int
	State  bool
}
