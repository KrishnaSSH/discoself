package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krishnassh/discoself"
	"github.com/krishnassh/discoself/types"
)

func init() {
	flag.StringVar(&token, "t", "", "User Token")
	flag.StringVar(&logFile, "f", "mentions.log", "Log file path")
	flag.Parse()
}

var token string
var logFile string
var client *discoself.Client

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: mentionlogger -t <user token>")
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
	fmt.Printf("Logging mentions to: %s\n", logFile)
}

func onMessage(e *types.MessageEventData) {
	for _, mention := range e.Mentions {
		if mention.ID == client.Gateway.Selfbot.User.ID {
			logMention(e)
			return
		}
	}
}

func logMention(e *types.MessageEventData) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	line := fmt.Sprintf("[%s] guild:%s channel:%s author:%s#%s said: %s\n",
		timestamp,
		e.GuildID,
		e.ChannelID,
		e.Author.Username,
		e.Author.Discriminator,
		e.Content,
	)

	if _, err := f.WriteString(line); err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}

	fmt.Print(line)
}
