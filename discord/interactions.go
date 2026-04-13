package discord

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

// GetSlashCommands fetches the slash command index for a guild.
func GetSlashCommands(gateway *Gateway, guildID string) (types.ApplicationCommandIndex, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/application-command-index")

	if err := requestClient.Do(req, resp); err != nil {
		return types.ApplicationCommandIndex{}, err
	}
	if resp.StatusCode() != 200 {
		return types.ApplicationCommandIndex{}, parseError(resp)
	}

	var index types.ApplicationCommandIndex
	if err := json.Unmarshal(resp.Body(), &index); err != nil {
		return types.ApplicationCommandIndex{}, err
	}
	return index, nil
}

// GetUserSlashCommands fetches the slash command index for the current user.
func GetUserSlashCommands(gateway *Gateway) (types.ApplicationCommandIndex, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, "https://discord.com/channels/@me")
	req.SetRequestURI("https://discord.com/api/v9/users/@me/application-command-index")

	if err := requestClient.Do(req, resp); err != nil {
		return types.ApplicationCommandIndex{}, err
	}
	if resp.StatusCode() != 200 {
		return types.ApplicationCommandIndex{}, parseError(resp)
	}

	var index types.ApplicationCommandIndex
	if err := json.Unmarshal(resp.Body(), &index); err != nil {
		return types.ApplicationCommandIndex{}, err
	}
	return index, nil
}

// SendSlashCommand sends a slash command interaction to a channel.
func SendSlashCommand(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand) bool {
	sessionID := gateway.SessionID
	if sessionID == "" {
		gateway.SessionID = GenerateSessionID()
		sessionID = gateway.SessionID
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", guildID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/interactions")
	req.SetBodyString(fmt.Sprintf(
		`{"type":2,"application_id":"%s","guild_id":"%s","channel_id":"%s","session_id":"%s","nonce":"%s","data":{"version":"%s","id":"%s","name":"%s","type":1,"options":[],"application_command":{"id":"%s","type":1,"application_id":"%s","version":"%s","name":"%s","description":"%s","dm_permission":true,"options":[],"integration_types":[0]},"attachments":[]},"analytics_location":"slash_ui"}`,
		command.ApplicationID, guildID, channelID, sessionID, GenerateNonce(),
		command.Version, command.ID, command.Name,
		command.ID, command.ApplicationID, command.Version, command.Name, command.Description,
	))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// SendSlashCommandWithOptions sends a slash command interaction with options to a channel.
func SendSlashCommandWithOptions(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand, options []any) bool {
	sessionID := gateway.SessionID
	if sessionID == "" {
		gateway.SessionID = GenerateSessionID()
		sessionID = gateway.SessionID
	}

	optionsJSON, err := json.Marshal(options)
	if err != nil {
		fmt.Println("Error marshalling options:", err)
		return false
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", guildID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/interactions")
	req.SetBodyString(fmt.Sprintf(
		`{"type":2,"application_id":"%s","guild_id":"%s","channel_id":"%s","session_id":"%s","nonce":"%s","data":{"version":"%s","id":"%s","name":"%s","type":1,"options":%s,"attachments":[]},"analytics_location":"slash_ui"}`,
		command.ApplicationID, guildID, channelID, sessionID, GenerateNonce(),
		command.Version, command.ID, command.Name, string(optionsJSON),
	))

	if err = requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// ClickButton clicks a message component button. interactionID is the custom_id of the button.
func ClickButton(gateway *Gateway, e *types.MessageEventData, interactionID string) bool {
	sessionID := gateway.SessionID
	if sessionID == "" {
		gateway.SessionID = GenerateSessionID()
		sessionID = gateway.SessionID
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", e.GuildID, e.ChannelID))
	req.SetRequestURI("https://discord.com/api/v9/interactions")
	req.SetBodyString(fmt.Sprintf(
		`{"type":3,"nonce":"%s","guild_id":"%s","channel_id":"%s","message_flags":0,"message_id":"%s","application_id":"%s","session_id":"%s","data":{"component_type":2,"custom_id":"%s"}}`,
		GenerateNonce(), e.GuildID, e.ChannelID, e.ID, e.Author.ID, sessionID, interactionID,
	))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}