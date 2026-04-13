package discord

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

// SendMessage sends a message to a channel.
func SendMessage(gateway *Gateway, channelID string, content string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages")
	req.SetBodyString(fmt.Sprintf(`{"mobile_network_type":"wifi","content":"%s","nonce":"%s","tts":false,"flags":0}`, content, GenerateNonce()))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// SendMessageWithReply sends a message that replies to another message.
func SendMessageWithReply(gateway *Gateway, channelID string, content string, replyMessageID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages")
	req.SetBodyString(fmt.Sprintf(
		`{"content":"%s","nonce":"%s","tts":false,"flags":0,"message_reference":{"channel_id":"%s","message_id":"%s"}}`,
		content, GenerateNonce(), channelID, replyMessageID,
	))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// EditMessage edits a message by ID.
func EditMessage(gateway *Gateway, channelID string, messageID string, content string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PATCH")
	req.Header.SetContentType("application/json")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)
	req.SetBodyString(fmt.Sprintf(`{"content":"%s"}`, content))

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 200
}

// DeleteMessage deletes a message by ID.
func DeleteMessage(gateway *Gateway, channelID string, messageID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("DELETE")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// SendTyping sends a typing indicator in a channel.
func SendTyping(gateway *Gateway, channelID string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("POST")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/typing")

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// AddReaction adds a reaction to a message. Emoji must be in format "emojiName:emojiID" for custom emojis or just the emoji character for unicode emojis.
func AddReaction(gateway *Gateway, channelID string, messageID string, emoji string) bool {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PUT")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions/" + emoji + "/@me")

	if err := requestClient.Do(req, resp); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions/" + emoji + "/@me")

	if err := requestClient.Do(req, resp); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/reactions")

	if err := requestClient.Do(req, resp); err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return resp.StatusCode() == 204
}

// GetChannel fetches a channel by ID.
func GetChannel(gateway *Gateway, channelID string) (types.Channel, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, "https://discord.com/channels/@me")
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID)

	if err := requestClient.Do(req, resp); err != nil {
		return types.Channel{}, err
	}
	if resp.StatusCode() != 200 {
		return types.Channel{}, parseError(resp)
	}

	var channel types.Channel
	if err := json.Unmarshal(resp.Body(), &channel); err != nil {
		return types.Channel{}, err
	}
	return channel, nil
}

// GetMessage fetches a single message by ID.
func GetMessage(gateway *Gateway, channelID string, messageID string) (types.MessageData, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID)

	if err := requestClient.Do(req, resp); err != nil {
		return types.MessageData{}, err
	}
	if resp.StatusCode() != 200 {
		return types.MessageData{}, parseError(resp)
	}

	var msg types.MessageData
	if err := json.Unmarshal(resp.Body(), &msg); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI(fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?limit=%d", channelID, limit))

	if err := requestClient.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, parseError(resp)
	}

	var msgs []types.MessageData
	if err := json.Unmarshal(resp.Body(), &msgs); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins")

	if err := requestClient.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, parseError(resp)
	}

	var msgs []types.MessageData
	if err := json.Unmarshal(resp.Body(), &msgs); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins/" + messageID)

	if err := requestClient.Do(req, resp); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/pins/" + messageID)

	if err := requestClient.Do(req, resp); err != nil {
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
	setCommonHeaders(req, gateway, fmt.Sprintf("https://discord.com/channels/%s/%s", gateway.Selfbot.User.ID, channelID))
	req.SetRequestURI("https://discord.com/api/v9/channels/" + channelID + "/messages/" + messageID + "/threads")
	req.SetBodyString(fmt.Sprintf(`{"name":"%s","auto_archive_duration":1440}`, name))

	if err := requestClient.Do(req, resp); err != nil {
		return types.Channel{}, err
	}
	if resp.StatusCode() != 200 {
		return types.Channel{}, parseError(resp)
	}

	var thread types.Channel
	if err := json.Unmarshal(resp.Body(), &thread); err != nil {
		return types.Channel{}, err
	}
	return thread, nil
}