package discoself

import (
	"strings"

	"github.com/krishnassh/discoself/discord"
	"github.com/krishnassh/discoself/types"
)

type Client struct {
	Selfbot *discord.Selfbot
	Gateway *discord.Gateway
	Config  *types.Config
}

func NewClient(token string, config *types.Config) *Client {
	token = strings.TrimSpace(token)
	selfbot := discord.Selfbot{Token: token}
	gateway := discord.CreateGateway(&selfbot, config)
	return &Client{&selfbot, gateway, config}
}

func (client *Client) Connect() error {
	return client.Gateway.Connect()
}

func (client *Client) AddHandler(event string, function any) error {
	return client.Gateway.Handlers.Add(event, function)
}

func (client *Client) GetMembers(guildId string, memberIds []string) error {
	return client.Gateway.GetMembers(guildId, memberIds)
}

func (client *Client) SendMessage(channelID string, content string) bool {
	return discord.SendMessage(client.Gateway, channelID, content)
}

func (client *Client) DeleteMessage(channelID string, messageID string) bool {
	return discord.DeleteMessage(client.Gateway, channelID, messageID)
}

func (client *Client) EditMessage(channelID string, messageID string, content string) bool {
	return discord.EditMessage(client.Gateway, channelID, messageID, content)
}

func (client *Client) SendTyping(channelID string) bool {
	return discord.SendTyping(client.Gateway, channelID)
}

func (client *Client) AddReaction(channelID string, messageID string, emoji string) bool {
	return discord.AddReaction(client.Gateway, channelID, messageID, emoji)
}

func (client *Client) SendMessageWithReply(channelID string, content string, replyMessageID string) bool {
	return discord.SendMessageWithReply(client.Gateway, channelID, content, replyMessageID)
}

func (client *Client) SendSlashCommand(channelID string, guildID string, command types.ApplicationCommand) bool {
	return discord.SendSlashCommand(client.Gateway, channelID, guildID, command)
}

func (client *Client) Close() {
	client.Gateway.Close()
}
