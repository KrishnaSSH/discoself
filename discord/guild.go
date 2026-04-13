package discord

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

// GetGuild fetches a guild by ID.
func GetGuild(gateway *Gateway, guildID string) (types.Guild, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID)

	if err := requestClient.Do(req, resp); err != nil {
		return types.Guild{}, err
	}
	if resp.StatusCode() != 200 {
		return types.Guild{}, parseError(resp)
	}

	var guild types.Guild
	if err := json.Unmarshal(resp.Body(), &guild); err != nil {
		return types.Guild{}, err
	}
	return guild, nil
}

// GetGuildChannels fetches all channels in a guild.
func GetGuildChannels(gateway *Gateway, guildID string) ([]types.Channel, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/channels")

	if err := requestClient.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, parseError(resp)
	}

	var channels []types.Channel
	if err := json.Unmarshal(resp.Body(), &channels); err != nil {
		return nil, err
	}
	return channels, nil
}

// GetGuildRoles fetches all roles in a guild.
func GetGuildRoles(gateway *Gateway, guildID string) ([]types.Role, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/roles")

	if err := requestClient.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, parseError(resp)
	}

	var roles []types.Role
	if err := json.Unmarshal(resp.Body(), &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

// KickMember kicks a member from a guild.
func KickMember(gateway *Gateway, guildID string, userID string) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/members/" + userID)

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// BanMember bans a member from a guild. deleteMessageSeconds sets how many
// seconds of messages to delete (0 to 604800).
func BanMember(gateway *Gateway, guildID string, userID string, deleteMessageSeconds int) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PUT")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/bans/" + userID)
	req.SetBodyString(fmt.Sprintf(`{"delete_message_seconds":%d}`, deleteMessageSeconds))

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// UnbanMember removes a ban from a user in a guild.
func UnbanMember(gateway *Gateway, guildID string, userID string) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/bans/" + userID)

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// AddRole adds a role to a guild member.
func AddRole(gateway *Gateway, guildID string, userID string, roleID string) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PUT")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/members/" + userID + "/roles/" + roleID)

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// RemoveRole removes a role from a guild member.
func RemoveRole(gateway *Gateway, guildID string, userID string, roleID string) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/members/" + userID + "/roles/" + roleID)

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// LeaveGuild leaves a guild.
func LeaveGuild(gateway *Gateway, guildID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.SetRequestURI("https://discord.com/api/v9/users/@me/guilds/" + guildID)

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// SetSlowmode sets the slowmode delay (in seconds) for a channel. Pass 0 to disable.
func SetSlowmode(gateway *Gateway, channelID string, seconds int) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PATCH")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, "https://discord.com/channels/@me")
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID)
	req.SetBodyString(fmt.Sprintf(`{"rate_limit_per_user":%d}`, seconds))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}