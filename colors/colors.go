package colors

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func LoadColors() ColorData {

	/*if _, err := os.Stat(filepath.Join(".", "users", "colors.csv")); os.IsNotExist(err) {
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}*/

	colorsF, err := os.Open(filepath.Join(".", "resources", "colors.csv"))

	if err != nil {
		x := ColorData{
			Colors: make(map[string]int),
			State:  true, //ensure that even if the file doesn't exist,we use a fallback color
		}
		return x
	}

	r := csv.NewReader(colorsF)
	colors, err := r.Read()

	//however these errors are far more important, and I really want to catch them
	if err == io.EOF {
		log.Fatal(err)
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}

	m := make(map[string]int)
	for x := 0; x < (len(colors)); x = x + 2 { //load all colors into a map. Each color is an even index[inc. 0], while each color value is an odd index.
		val, err := strconv.Atoi(colors[x+1])
		if err != nil {
			break
		}
		m[colors[x]] = val
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
