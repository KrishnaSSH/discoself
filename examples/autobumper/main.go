package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krishnassh/discoself"
	"github.com/krishnassh/discoself/discord"
	"github.com/krishnassh/discoself/types"
)

func init() {
	flag.StringVar(&token, "t", "", "User Token")
	flag.StringVar(&guildID, "g", "", "Guild ID")
	flag.StringVar(&channelID, "c", "", "Channel ID")
	flag.Parse()
}

// Disboard's bot application ID
const disboardAppID = "302050872383242240"

var token string
var guildID string
var channelID string
var client *discoself.Client

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: bump -t <user token> -g <guild id> -c <channel id>")
		return
	}
	if guildID == "" {
		fmt.Println("No guild ID provided. Please run: bump -t <user token> -g <guild id> -c <channel id>")
		return
	}
	if channelID == "" {
		fmt.Println("No channel ID provided. Please run: bump -t <user token> -g <guild id> -c <channel id>")
		return
	}

	client = discoself.NewClient(token, &types.DefaultConfig)
	client.AddHandler(types.GatewayEventReady, onReady)

	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("Running. press ctrl-c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	client.Close()
}

func onReady(e *types.ReadyEventData) {
	fmt.Printf("Logged in as: %s\n", e.User.Username)

	go func() {
		for {
			sendBump()
			time.Sleep(2 * time.Hour)
		}
	}()
}

func sendBump() {
	cmds, err := discord.GetSlashCommands(client.Gateway, guildID)
	if err != nil {
		fmt.Println("Error fetching slash commands:", err)
		return
	}

	for _, cmd := range cmds.ApplicationCommand {
		// Match both name AND application ID to avoid hitting the wrong bot's bump command
		if cmd.Name == "bump" && cmd.ApplicationID == disboardAppID {
			ok := client.SendSlashCommand(channelID, guildID, cmd)
			if ok {
				fmt.Printf("[%s] /bump sent successfully!\n", time.Now().Format("2006-01-02 15:04:05"))
			} else {
				fmt.Printf("[%s] /bump failed\n", time.Now().Format("2006-01-02 15:04:05"))
			}
			return
		}
	}

	fmt.Println("Disboard bump command not found is Disboard in this server?")
}
