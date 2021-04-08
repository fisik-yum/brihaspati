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

	//however these errors are far more important, and I really want to catch them
	if err != nil {
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}

	r := csv.NewReader(colorsF)
	colors, err := r.Read()

	if err == io.EOF {
		log.Fatal(err)
		x := ColorData{
			Colors: make(map[string]int),
			State:  false,
		}
		return x
	}
	m := make(map[string]int)
	for x := 0; x < len(colors); x += 2 { //load all colors into a map. Each color is an even index[inc. 0], while each color value is an odd index.
		val, _ := strconv.Atoi(colors[x+1])
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

func ListColors(cd ColorData) []string {
	keys := make([]string, 0, len(cd.Colors))
	values := make([]int, 0, len(cd.Colors))

	for k, v := range cd.Colors {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys
}
