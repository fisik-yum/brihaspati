package auth

/*
The go module that is used to generate, handle and compare auth tokens for the bot.
The RNG is not crypto-secure, and I'll probably fix that in the future
Chungus Code

I reimplemented most of this code to work with csv files so that it will be easier to manage authorization. It seems to clear up a lot of clutter.

*/
import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type userDetails struct {
	id    string
	code  string
	aTime string
	bTime string
}

func getRandomNumber() string {
	rand.Seed(time.Now().UnixNano()) //seed for random number
	return strconv.Itoa(rand.Intn(9999-1000) + 1000)
}

func CreateCode(id string) string { //Make our lives easier.
	uTime := strconv.FormatInt(time.Now().Unix(), 10)
	writeTo(id, getRandomNumber(), uTime, uTime)
	return readItem(id).code //return id, which is stored at index 1

}

func writeTo(id string, code string, aTime string, bTime string) { //extracted to a function so I can make life easier.

	records := [][]string{
		{id, code, aTime, bTime},
	}

	f, err := os.Create(filepath.Join(".", "users", id+".csv"))
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

func readItem(id string) userDetails { // makes it easy to read stored data
	f, err := os.Open(filepath.Join(".", "users", id+".csv"))

	if err != nil {

		return userDetails{}
	}

	r := csv.NewReader(f)

	record, err := r.Read()
	if err == io.EOF {
		log.Fatal(err)
	}

	x := userDetails{
		id:    record[0],
		code:  record[1],
		aTime: record[2],
		bTime: record[3],
	}
	return x
}

func ValidateCode(id string, message string) bool {
	if _, err := os.Stat(filepath.Join(".", "users", id+".csv")); os.IsNotExist(err) {
		return false
	}

	s := strings.Fields(message)
	if len(s) < 2 {
		return false
	}
	aTime, err := strconv.Atoi(readItem(id).aTime) //time of code generation

	if err != nil {
		return false
	}
	bTime, err := strconv.Atoi(readItem(id).bTime) //time of code authentication
	vTime := int(time.Now().Unix())

	if err != nil {
		return false
	}

	state := (readItem(id).code == s[1]) && ((aTime + 300) > vTime) && (aTime == bTime)
	if state {
		writeTo(id, readItem(id).code, strconv.Itoa(aTime), strconv.Itoa(vTime))
		return true //return bool to be used in an if statement
	} else {
		return false
	}
}

func CodeState(id string) bool {
	if _, err := os.Stat(filepath.Join(".", "users", id+".csv")); os.IsNotExist(err) {
		return false
	}
	aTime, err := strconv.Atoi(readItem(id).aTime) //generation time
	if err != nil {
		return false
	}
	bTime, err := strconv.Atoi(readItem(id).bTime) //authentication time
	vTime := int(time.Now().Unix())                //current time
	if err != nil {
		return false
	}

	if (bTime+300 > vTime) && (aTime != bTime) {
		return true
	} else {
		return false
	}
}
