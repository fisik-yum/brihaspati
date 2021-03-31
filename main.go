package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dshardmanager"
)

// Variables used for command line parameters
var (
	FlagToken      string
	FlagLogChannel string
)

/*func main() {

// Create a new Discord session using the provided bot token.
dg, err := discordgo.New("Bot " + Token)
if err != nil {
	fmt.Println("error creating Discord session,", err)
	return
}
_, err = os.Stat("test")

if os.IsNotExist(err) {
	errDir := os.MkdirAll("users", 0755)
	os.MkdirAll("guilds", 0755)
	if errDir != nil {
		log.Fatal(err)
	}
}*/
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
	manager.Name = "ExampleBot"
	manager.LogChannel = FlagLogChannel
	manager.StatusMessageChannel = FlagLogChannel

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
	log.Fatal(http.ListenAndServe(":7441", nil))
	select {}
}

// Register the messageCreate func as a callback for MessageCreate events.

// In this example, we only care about receiving message events.
//dg.Identify.Intents = discordgo.IntentsGuildMessages nopeeee
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
}
