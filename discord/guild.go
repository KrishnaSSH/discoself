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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.Guild{}, err
	}

	var guild types.Guild
	if err = json.Unmarshal(resp.Body(), &guild); err != nil {
		fmt.Println("Error:", err)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/channels")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var channels []types.Channel
	if err = json.Unmarshal(resp.Body(), &channels); err != nil {
		fmt.Println("Error:", err)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/roles")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var roles []types.Role
	if err = json.Unmarshal(resp.Body(), &roles); err != nil {
		fmt.Println("Error:", err)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/members/" + userID)

	if err := requestClient.Do(req, resp); err != nil {
		return err
	}
	if resp.StatusCode() != 204 {
		return parseError(resp)
	}
	return nil
}

// BanMember bans a member from a guild.
func BanMember(gateway *Gateway, guildID string, userID string, deleteMessageSeconds int) error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PUT")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("content-length", "0")
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/users/@me/guilds/" + guildID)

	err := requestClient.Do(req, resp)
	if err != nil {
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
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID)
	req.SetBodyString(fmt.Sprintf(`{"rate_limit_per_user":%d}`, seconds))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}