package types

type Config struct {
	Presence       string
	ApiVersion     string
	Browser        string
	BrowserVersion string
	Capabilities   int64
	Device         string
	Os             string
	OsVersion      string
	UserAgent      string
}

var DefaultConfig = Config{
	Presence:       "offline",
	ApiVersion:     "9",
	Browser:        "Chrome",
	BrowserVersion: "124.0.0.0",
	Capabilities:   4093,
	Device:         "",
	Os:             "",
	OsVersion:      "",
	UserAgent:      "Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
}
