package discord

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
	"github.com/valyala/fasthttp"
)

const DiscordEpoch = 1420070400000

func UtcNow() time.Time {
	return time.Now().UTC()
}

func TimeSnowflake(dt time.Time, high bool) int64 {
	discordMillis := dt.UnixNano()/1e6 - DiscordEpoch
	if high {
		return (discordMillis << 22) + (1<<22 - 1)
	}
	return discordMillis << 22
}

func GenerateNonce() string {
	return fmt.Sprintf("%d", TimeSnowflake(UtcNow(), false))
}

func GenerateSessionID() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 16
	var b strings.Builder
	for range length {
		b.WriteRune(chars[r.Intn(len(chars))])
	}
	return b.String()
}

func GenerateSuperProperties(gateway *Gateway) string {
	if gateway.superProperties != "" {
		return gateway.superProperties
	}

	super := &types.SuperProperties{
		OS:                     gateway.Config.Os,
		Browser:                gateway.Config.Browser,
		Device:                 gateway.Config.Device,
		SystemLocale:           clientLocale,
		BrowserUserAgent:       gateway.Config.UserAgent,
		BrowserVersion:         gateway.Config.BrowserVersion,
		OSVersion:              gateway.Config.OsVersion,
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      gateway.ClientBuildNumber,
		ClientEventSource:      nil,
	}

	jsonData, err := json.Marshal(super)
	if err != nil {
		fmt.Println("Error marshalling super properties:", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(jsonData)
}

func setCommonHeaders(req *fasthttp.Request, gateway *Gateway, referrer string) {
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"134\", \"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-arch", "\"x86\"")
	req.Header.Set("sec-ch-ua-bitness", "\"64\"")
	req.Header.Set("sec-ch-ua-full-version", "\"134.0.6998.118\"")
	req.Header.Set("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"134.0.6998.118\", \"Chromium\";v=\"134.0.6998.118\", \"Not:A-Brand\";v=\"24.0.0.0\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", "\"\"")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-ch-ua-platform-version", "\"19.0.0\"")
	req.Header.Set("sec-ch-ua-wow64", "?0")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("origin", "https://discord.com")
	req.Header.Set("referer", referrer)
	req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
	req.Header.Set("priority", "u=1, i")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
}
