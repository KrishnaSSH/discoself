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

func (client *Client) Close() {
	client.Gateway.Close()
}

// channel

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

// interactions

func (client *Client) GetSlashCommands(guildID string) (types.ApplicationCommandIndex, error) {
	return discord.GetSlashCommands(client.Gateway, guildID)
}

func (client *Client) GetUserSlashCommands() (types.ApplicationCommandIndex, error) {
	return discord.GetUserSlashCommands(client.Gateway)
}

func (client *Client) SendSlashCommand(channelID string, guildID string, command types.ApplicationCommand) bool {
	return discord.SendSlashCommand(client.Gateway, channelID, guildID, command)
}

func (client *Client) SendSlashCommandWithOptions(channelID string, guildID string, command types.ApplicationCommand, options []any) bool {
	return discord.SendSlashCommandWithOptions(client.Gateway, channelID, guildID, command, options)
}

func (client *Client) ClickButton(e *types.MessageEventData, interactionID string) bool {
	return discord.ClickButton(client.Gateway, e, interactionID)
}

// user

func (client *Client) GetUser(userID string) (types.User, error) {
	return discord.GetUser(client.Gateway, userID)
}

func (client *Client) GetProfile(userID string, guildID string) (types.User, error) {
	return discord.GetProfile(client.Gateway, userID, guildID)
}

func (client *Client) ModifyUsername(username string, password string) bool {
	return discord.ModifyUsername(client.Gateway, username, password)
}

func (client *Client) SetStatus(status string) bool {
	return discord.SetStatus(client.Gateway, status)
}

func (client *Client) SetCustomStatus(text string, emoji string) bool {
	return discord.SetCustomStatus(client.Gateway, text, emoji)
}

func (client *Client) ClearCustomStatus() bool {
	return discord.ClearCustomStatus(client.Gateway)
}

func (client *Client) SetNickname(guildID string, nickname string) bool {
	return discord.SetNickname(client.Gateway, guildID, nickname)
}

func (client *Client) SendFriendRequest(username string) bool {
	return discord.SendFriendRequest(client.Gateway, username)
}

func (client *Client) RemoveFriend(userID string) bool {
	return discord.RemoveFriend(client.Gateway, userID)
}

func (client *Client) BlockUser(userID string) bool {
	return discord.BlockUser(client.Gateway, userID)
}