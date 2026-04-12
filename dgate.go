package discoself

import (
	"github.com/krishnassh/discoself/discord"
	"github.com/krishnassh/discoself/types"
)

type Client struct {
	Selfbot *discord.Selfbot
	Gateway *discord.Gateway
	Config  *types.Config
}

func NewClient(token string, config *types.Config) *Client {
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
