package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dshardmanager"
)

// Variables used for command line parameters
var (
	FlagToken      string
	FlagLogChannel string
)

func main() {

	flag.StringVar(&FlagToken, "t", "", "Discord token")
	flag.StringVar(&FlagLogChannel, "c", "", "Log channel, optional")
	flag.Parse()

	log.Println("Starting v" + dshardmanager.VersionString)
	if FlagToken == "" {
		FlagToken = os.Getenv("DG_TOKEN")
		if FlagToken == "" {
			log.Fatal("No token specified")
		}
	}

	if !strings.HasPrefix(FlagToken, "Bot ") {
		log.Fatal("dshardmanager only works on bot accounts, did you forget to add `Bot ` before the token?")
	}

	manager := dshardmanager.New(FlagToken)
	manager.Name = "Brihaspati"
	manager.LogChannel = FlagLogChannel
	manager.StatusMessageChannel = FlagLogChannel

	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("users", 0755)
		os.MkdirAll("guilds", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	recommended, err := manager.GetRecommendedCount()
	if err != nil {
		log.Fatal("Failed getting recommended shard count")
	}
	if recommended < 2 {
		manager.SetNumShards(5)
	}
	log.Println("Starting shard manager")
	err = manager.Start()
	if err != nil {
		log.Fatal("Failed to start: ", err)
	}
	//dg.Identify.Intents = 4679
	log.Println("Started!")
	manager.AddHandler(messageCreate)
	manager.AddHandler(ListenForAction)
	manager.AddHandler(Moderate)
	manager.AddHandler(funStuff)
	//log.Fatal(http.ListenAndServe(":7441", nil))
	//select {}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	manager.StopAll()
}

//discordgo.IntentsAll //<- unneedeed for the time being
//https://ziad87.net/intents/ helped me get the intents number.

//s.StateEnabled = true
//s.State.MaxMessageCount = 50

// Open a websocket connection to Discord and begin listening.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "^help" {
		fileUrl := "https://raw.githubusercontent.com/fisik-yum/brihaspati/main/help.txt"
		err := DownloadFile("help.txt", fileUrl)
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadFile("help.txt")
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		s.ChannelMessageSend(m.ChannelID, string(data))
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
