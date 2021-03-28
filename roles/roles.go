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
)

func writeTo(guildID string, roleID string) { //extracted to a function so I can make life easier.

	records := [][]string{
		{roleID},
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
