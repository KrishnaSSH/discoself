package discord

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

func SendMessage(gateway *Gateway, channelID string, content string) bool {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Chrome OS\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("referrer", fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages")
	req.SetBodyString(fmt.Sprintf("{\"mobile_network_type\":\"wifi\",\"content\":\"%s\",\"nonce\":\"%s\",\"tts\":false,\"flags\":0}", content, GenerateNonce()))
	resp := fasthttp.AcquireResponse()
	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode() == 200 {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
		return true
	} else {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
		return false
	}
}

func SendTyping(gateway *Gateway, channelID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("content-length", "0")
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/typing")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

func DeleteMessage(gateway *Gateway, channelID string, messageID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

func EditMessage(gateway *Gateway, channelID string, messageID string, content string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)
	req.SetBodyString(fmt.Sprintf(`{"content":"%s"}`, content))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

func SendMessageWithReply(gateway *Gateway, channelID string, content string, replyMessageID string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages")
	req.SetBodyString(fmt.Sprintf(
		`{"content":"%s","nonce":"%s","tts":false,"flags":0,"message_reference":{"channel_id":"%s","message_id":"%s"}}`,
		content, GenerateNonce(), channelID, replyMessageID,
	))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

func AddReaction(gateway *Gateway, channelID string, messageID string, emoji string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions/" + emoji + "/@me")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// RemoveReaction removes your own reaction from a message.
func RemoveReaction(gateway *Gateway, channelID string, messageID string, emoji string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions/" + emoji + "/@me")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// DeleteAllReactions removes all reactions from a message.
func DeleteAllReactions(gateway *Gateway, channelID string, messageID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// GetMessage fetches a single message by ID.
func GetMessage(gateway *Gateway, channelID string, messageID string) (types.MessageData, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.MessageData{}, err
	}

	var msg types.MessageData
	if err = json.Unmarshal(resp.Body(), &msg); err != nil {
		fmt.Println("Error:", err)
		return types.MessageData{}, err
	}
	return msg, nil
}

// GetMessages fetches up to 100 messages from a channel. Pass limit 1-100.
func GetMessages(gateway *Gateway, channelID string, limit int) ([]types.MessageData, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if limit < 1 {
		limit = 1
	} else if limit > 100 {
		limit = 100
	}

	req.Header.SetMethod("GET")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI(fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?limit=%d", channelID, limit))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var msgs []types.MessageData
	if err = json.Unmarshal(resp.Body(), &msgs); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return msgs, nil
}

// GetPinnedMessages fetches all pinned messages in a channel.
func GetPinnedMessages(gateway *Gateway, channelID string) ([]types.MessageData, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins")

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var msgs []types.MessageData
	if err = json.Unmarshal(resp.Body(), &msgs); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return msgs, nil
}

// PinMessage pins a message in a channel.
func PinMessage(gateway *Gateway, channelID string, messageID string) bool {
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
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins/" + messageID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// UnpinMessage unpins a message in a channel.
func UnpinMessage(gateway *Gateway, channelID string, messageID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins/" + messageID)

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// CreateThread creates a public thread from an existing message.
func CreateThread(gateway *Gateway, channelID string, messageID string, name string) (types.Channel, error) {
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
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/threads")
	req.SetBodyString(fmt.Sprintf(`{"name":"%s","auto_archive_duration":1440}`, name))

	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.Channel{}, err
	}

	var thread types.Channel
	if err = json.Unmarshal(resp.Body(), &thread); err != nil {
		fmt.Println("Error:", err)
		return types.Channel{}, err
	}
	return thread, nil
}