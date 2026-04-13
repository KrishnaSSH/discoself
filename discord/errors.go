package discord

import (
    "fmt"

    "github.com/goccy/go-json"
    "github.com/krishnassh/discoself/types"
    "github.com/valyala/fasthttp"
)

type DiscordError struct {
    types.DiscordError
}

func (e *DiscordError) Error() string {
    return fmt.Sprintf("discord error %d: %s", e.Code, e.Message)
}

func parseError(resp *fasthttp.Response) error {
    var e DiscordError
    if err := json.Unmarshal(resp.Body(), &e); err == nil && e.Code != 0 {
        return &e
    }
    return fmt.Errorf("http error %d", resp.StatusCode())
}