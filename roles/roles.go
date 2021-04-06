package roles

/*
This is to handle code for keeping track of roles(e.g: muteroles), and creating overrides for guild channels
*/
import (
	"brihaspati/colors"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func writeTo(guildID string, muteroleID string) { //extracted to a function so I can make life easier.

	records := [][]string{
		{guildID, muteroleID},
	}

	f, err := os.Create(filepath.Join(".", "guilds", guildID+".csv"))
	if err != nil {
		log.Fatalln("failed to open file", err)
		return
	}

	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
		return
	}

}

func readDataFromFile(id string) (record []string) { //id is to find the said .csv file
	f, err := os.Open(filepath.Join(".", "guilds", id+".csv"))

	if err != nil {

		return record
	}

	r := csv.NewReader(f)

	record, err = r.Read()
	if err == io.EOF {
		log.Fatal(err)
	}

	return record
}

func ReadItem(id string, index int) string { // makes it easy to read stored data
	return readDataFromFile(id)[index]
}

func CheckIfMuteExists(guildID string, roles []*discordgo.Role) bool {
	if checkIfRecordExists(guildID) {
		return checkIfIDExists(ReadItem(guildID, 1), roles)
	}
	return false
}

func checkIfRecordExists(id string) bool { //check if the csv file exists
	if _, err := os.Stat(filepath.Join(".", "guilds", id+".csv")); err == nil {
		// path/to/whatever exists
		return true
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Println(err)
		return false
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

	}
}

func checkIfIDExists(id string, roles []*discordgo.Role) bool { //check if the role ID is still valid
	for x := range roles {
		if id == roles[x].ID {
			return true
		}
	}
	return false
}

func ApplyChannelOverrides(roleID string, channels []*discordgo.Channel, s *discordgo.Session) {
	for x := range channels { //cycle through each channel and apply overrides.
		s.ChannelPermissionSet(channels[x].ID, roleID, discordgo.PermissionOverwriteTypeRole, discordgo.PermissionReadMessageHistory, 2048) //2048 is the override value
	}
}

//TODO: add functions to create roles and apply permission overrides to all channels in a server

func CreateMuteRole(guildID string, s *discordgo.Session) bool {
	role, err := s.GuildRoleCreate(guildID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	roleID := role.ID
	s.GuildRoleEdit(guildID, roleID, "Muted", 16111426, false, 0, false)
	writeTo(guildID, roleID)
	return true

}

func CreateNewRole(message string, guildID string, cMap map[string]int, s *discordgo.Session) bool { // format ^createrole name,color. include map of colors to avoid reading all the time.
	items := strings.Split(message, ",")
	loadC := colors.LoadColors()
	if len(items) < 2 || !loadC.State {
		return false
	}
	name := items[0]                                                //dont strip spaces
	color := strings.ReplaceAll(strings.ToLower(items[1]), " ", "") //strip all spaces

	if val, ok := cMap[color]; !ok {
		val = 8421504
		cRole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return false
		}
		s.GuildRoleEdit(guildID, cRole.ID, name, val, true, 0, true)
		return true
	} else {
		cRole, err := s.GuildRoleCreate(guildID)
		if err != nil {
			return false
		}
		s.GuildRoleEdit(guildID, cRole.ID, name, val, true, 0, true)
		return true
	}

}
