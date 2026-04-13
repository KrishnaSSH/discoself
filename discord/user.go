package discord

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

// GetUser fetches a user's profile by ID.
func GetUser(gateway *Gateway, userID string) (types.User, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/users/" + userID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.User{}, err
	}

	var user types.User
	if err = json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return types.User{}, err
	}

	return user, nil
}

// GetProfile fetches the full profile of a user in a guild context.
func GetProfile(gateway *Gateway, userID string, guildID string) (types.User, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI(fmt.Sprintf("https://discord.com/api/v9/users/%s/profile?guild_id=%s", userID, guildID))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.User{}, err
	}

	var user types.User
	if err = json.Unmarshal(resp.Body(), &user); err != nil {
		fmt.Println("Error:", err)
		return types.User{}, err
	}

	return user, nil
}

// ModifyUsername changes the account's username. Requires the account password.
func ModifyUsername(gateway *Gateway, username string, password string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/users/@me")
	req.SetBodyString(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// SetStatus sets the online status. Valid values: "online", "idle", "dnd", "invisible".
func SetStatus(gateway *Gateway, status string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/users/@me/settings")
	req.SetBodyString(fmt.Sprintf(`{"status":"%s"}`, status))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// SetCustomStatus sets a custom status message and optional emoji.
// Pass an empty string for emoji to set a text-only status.
func SetCustomStatus(gateway *Gateway, text string, emoji string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/users/@me/settings")

	var body string
	if emoji != "" {
		body = fmt.Sprintf(`{"custom_status":{"text":"%s","emoji_name":"%s"}}`, text, emoji)
	} else {
		body = fmt.Sprintf(`{"custom_status":{"text":"%s"}}`, text)
	}
	req.SetBodyString(body)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// ClearCustomStatus removes the custom status.
func ClearCustomStatus(gateway *Gateway) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/users/@me/settings")
	req.SetBodyString(`{"custom_status":null}`)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// SetNickname changes the account's nickname in a guild.
// Pass an empty string to reset the nickname.
func SetNickname(gateway *Gateway, guildID string, nickname string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/members/@me")
	req.SetBodyString(fmt.Sprintf(`{"nick":"%s"}`, nickname))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// SendFriendRequest sends a friend request to a user by username.
func SendFriendRequest(gateway *Gateway, username string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/users/@me/relationships")
	req.SetBodyString(fmt.Sprintf(`{"username":"%s","discriminator":null}`, username))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// RemoveFriend removes a friend or cancels an outgoing friend request by user ID.
func RemoveFriend(gateway *Gateway, userID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/users/@me/relationships/" + userID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// BlockUser blocks a user by ID.
func BlockUser(gateway *Gateway, userID string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/users/@me/relationships/" + userID)
	req.SetBodyString(`{"type":2}`)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}
