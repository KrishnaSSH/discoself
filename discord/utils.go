package discord

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/krishnassh/discoself/types"
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
	super := &types.SuperProperties{
		OS:                     gateway.Config.Os,
		Browser:                gateway.Config.Browser,
		Device:                 gateway.Config.Device,
		SystemLocale:           gateway.Selfbot.User.Locale,
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
	base64Data := base64.StdEncoding.EncodeToString(jsonData)

	return base64Data
}
