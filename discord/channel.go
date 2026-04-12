package discord

import (
	"fmt"
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
