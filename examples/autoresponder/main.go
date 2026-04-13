package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/krishnassh/discoself"
	"github.com/krishnassh/discoself/types"
)

func init() {
	flag.StringVar(&token, "t", "", "User Token")
	flag.StringVar(&channelID, "c", "", "Channel ID (optional, if not set responds in all channels)")
	flag.StringVar(&trigger, "r", "!api", "Trigger phrase to respond to")
	flag.StringVar(&response, "R", "read the api reference here https://github.com/krishnassh/discoself/blob/main/docs/api.md", "Response message")
	flag.Parse()
}

var token string
var channelID string
var trigger string
var response string
var client *discoself.Client

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: autoresponder -t <user token>")
		return
	}

	client = discoself.NewClient(token, &types.DefaultConfig)
	client.AddHandler(types.GatewayEventReady, onReady)
	client.AddHandler(types.GatewayEventMessageCreate, onMessage)

	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("Running. Press ctrl-c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	client.Close()
}

func onReady(e *types.ReadyEventData) {
	fmt.Printf("Logged in as: %s\n", e.User.Username)
	fmt.Printf("Trigger: '%s' -> Response: '%s'\n", trigger, response)
	if channelID != "" {
		fmt.Printf("Monitoring channel ID: %s\n", channelID)
	} else {
		fmt.Println("Monitoring all channels")
	}
}

func onMessage(e *types.MessageEventData) {
	if e.Author.ID == client.Gateway.Selfbot.User.ID {
		return
	}

	if channelID != "" && e.ChannelID != channelID {
		return
	}

	if strings.Contains(strings.ToLower(e.Content), strings.ToLower(trigger)) {
		sendResponse(e.ChannelID)
	}
}

func sendResponse(chID string) {
	ok := client.SendMessage(chID, response)
	if ok {
		fmt.Printf("[%s] Auto-responded in channel %s\n", time.Now().Format("2006-01-02 15:04:05"), chID)
	} else {
		fmt.Printf("[%s] Failed to send response\n", time.Now().Format("2006-01-02 15:04:05"))
	}
}
