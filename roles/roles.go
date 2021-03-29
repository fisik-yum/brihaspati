package roles

/*
This is to handle code for keeping track of roles(e.g: muteroles), and creating overrides for guild channels
*/
import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func writeTo(guildID string, muteroleID string) { //extracted to a function so I can make life easier.

	records := [][]string{
		{guildID, muteroleID},
	}

	f, err := os.Create(filepath.Join(".", "guilds", guildID+".csv"))
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
		return
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
		return
	}

}

func readDataFromFile(id string) (record []string) { //id is to find the said .csv file
	f, err := os.Open(filepath.Join(".", "users", id+".csv"))

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

func readItem(id string, index int) string { // makes it easy to read stored data
	return readDataFromFile(id)[index]
}

func checkIfMuteExists(guildID string, roles []*discordgo.Role) bool {
	if checkIfRecordExists(guildID) {
		if checkIfIDExists(readItem(guildID, 1), roles) {
			return true
		}
		return false
	}
	return false
}

func checkIfRecordExists(id string) bool { //check if the csv file exists
	if _, err := os.Stat(filepath.Join(".", "users", id+".csv")); err == nil {
		// path/to/whatever exists
		return true
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		return false
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

	}
}

func checkIfIDExists(id string, roles []*discordgo.Role) bool { //check if the role ID is still valid
	for x := 0; x < len(roles); x++ {
		if id == roles[x].ID {
			return true
		}
	}
	return false
}

//TODO: add functions to create roles and apply permission overrides to all channels in a server
