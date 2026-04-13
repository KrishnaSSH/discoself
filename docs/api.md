# discoself API Reference

exported functions and types are documented here.

---

## Table of Contents

- [Client (dgate.go)](#client-dgatego)
- [Event Constants (types/subscribable.go)](#event-constants)
- [Handler Signatures](#handler-signatures)
- [discord/channel.go](#discordchannelgo)
- [discord/interactions.go](#discordinteractionsgo)
- [discord/gateway.go](#discordgatewaygo)
- [discord/handlers.go](#discordhandlersgo)
- [discord/utils.go](#discordutilsgo)
- [types/config.go](#typesconfiggo)
- [types/events.go](#typeseventgo)
- [types/discord.go](#typesdiscordgo)

---

## Client (dgate.go)

The `Client` struct is the main entrypoint.

```go
type Client struct {
    Selfbot *discord.Selfbot
    Gateway *discord.Gateway
    Config  *types.Config
}
```

The logged-in user is available at `client.Gateway.Selfbot.User` after `Connect` returns. This is a `types.User` struct populated from the READY event before any handlers fire.

---

### `NewClient`

```go
func NewClient(token string, config *types.Config) *Client
```

Creates a new client. `config` must not be nil -- pass `&types.DefaultConfig` to use the defaults.

```go
client := discoself.NewClient("your-user-token", &types.DefaultConfig)
```

---

### `Connect`

```go
func (client *Client) Connect() error
```

Opens the WebSocket connection, runs the login handshake, and starts the event loop. Blocks until the connection is established. The READY event fires and handlers are called before this returns.

```go
if err := client.Connect(); err != nil {
    fmt.Println("Error connecting:", err)
    return
}
```

---

### `Close`

```go
func (client *Client) Close()
```

Closes the gateway connection and stops the event loop.

```go
client.Close()
```

---

### `AddHandler`

```go
func (client *Client) AddHandler(event string, function any) error
```

Registers a callback for a gateway event. Must be called before `Connect`. The function signature must exactly match the event type -- passing the wrong signature silently fails (the handler is not registered). Use the constants from `types/subscribable.go` for event names.

Supported events and their required function signatures:

| Event constant | Function signature |
|---|---|
| `types.GatewayEventReady` | `func(data *types.ReadyEventData)` |
| `types.GatewayEventMessageCreate` | `func(data *types.MessageEventData)` |
| `types.GatewayEventMessageUpdate` | `func(data *types.MessageEventData)` |
| `types.GatewayEventMessageDelete` | `func(data *types.MessageDeleteEventData)` |
| `types.GatewayEventTypingStart` | `func(data *types.TypingStartEventData)` |
| `types.GatewayEventPresenceUpdate` | `func(data *types.PresenceUpdateEventData)` |
| `types.GatewayGuildMembersChunk` | `func(data *types.GuildMembersChunkEventData)` |
| `types.GatewayEventReconnect` | `func()` |
| `types.GatewayEventInvalidated` | `func()` |

```go
client.AddHandler(types.GatewayEventReady, func(data *types.ReadyEventData) {
    fmt.Printf("Logged in as: %s\n", data.User.Username)
})

client.AddHandler(types.GatewayEventMessageCreate, func(data *types.MessageEventData) {
    fmt.Println(data.Content)
})
```

---

### `SendMessage`

```go
func (client *Client) SendMessage(channelID string, content string) bool
```

Sends a text message to a channel. Returns `true` on success.

```go
client.SendMessage("1234567890123456789", "hello")
```

---

### `SendMessageWithReply`

```go
func (client *Client) SendMessageWithReply(channelID string, content string, replyMessageID string) bool
```

Sends a message as a reply to `replyMessageID`. Returns `true` on success.

```go
client.SendMessageWithReply("1234567890123456789", "got it", "9876543210987654321")
```

---

### `EditMessage`

```go
func (client *Client) EditMessage(channelID string, messageID string, content string) bool
```

Edits one of your own messages. Returns `true` on success.

```go
client.EditMessage("1234567890123456789", "1122334455667788990", "updated content")
```

---

### `DeleteMessage`

```go
func (client *Client) DeleteMessage(channelID string, messageID string) bool
```

Deletes one of your own messages. Returns `true` on success.

```go
client.DeleteMessage("1234567890123456789", "1122334455667788990")
```

---

### `SendTyping`

```go
func (client *Client) SendTyping(channelID string) bool
```

Sends a typing indicator to a channel. Returns `true` on success.

```go
client.SendTyping("1234567890123456789")
```

---

### `AddReaction`

```go
func (client *Client) AddReaction(channelID string, messageID string, emoji string) bool
```

Adds a reaction to a message. For Unicode emoji pass the character directly. For custom emoji use `name:id` format. Returns `true` on success.

```go
client.AddReaction("1234567890123456789", "1122334455667788990", "🐢")
client.AddReaction("1234567890123456789", "1122334455667788990", "myemoji:9988776655443322110")
```

---

### `SendSlashCommand`

```go
func (client *Client) SendSlashCommand(channelID string, guildID string, command types.ApplicationCommand) bool
```

Fires a slash command with no options. Get a valid `types.ApplicationCommand` from `discord.GetSlashCommands` first. Returns `true` on success.

```go
client.SendSlashCommand("1234567890123456789", "9876543210987654321", command)
```

---

### `GetMembers`

```go
func (client *Client) GetMembers(guildID string, memberIDs []string) error
```

Requests member data for a list of user IDs in a guild. Results arrive asynchronously via the `types.GatewayGuildMembersChunk` event.

```go
client.GetMembers("9876543210987654321", []string{"111111111111111111", "222222222222222222"})

client.AddHandler(types.GatewayGuildMembersChunk, func(data *types.GuildMembersChunkEventData) {
    for _, member := range data.Members {
        fmt.Println(member.User.Username)
    }
})
```

---

## Event Constants

Defined in `types/subscribable.go`. Always use these constants with `AddHandler` rather than raw strings.

```go
const (
    GatewayEventReady          = "READY"
    GatewayEventMessageCreate  = "MESSAGE_CREATE"
    GatewayEventMessageUpdate  = "MESSAGE_UPDATE"
    GatewayEventMessageDelete  = "MESSAGE_DELETE"
    GatewayEventTypingStart    = "TYPING_START"
    GatewayEventPresenceUpdate = "PRESENCE_UPDATE"
    GatewayEventReconnect      = "RECONNECT"
    GatewayEventInvalidated    = "INVALIDATED"
    GatewayGuildMembersChunk   = "GUILD_MEMBERS_CHUNK"
)
```

---

## Handler Signatures

`AddHandler` uses a type assertion on the function value. If the signature does not match exactly, the handler is silently dropped and an error is returned. Always check the return value of `AddHandler` during development.

```go
if err := client.AddHandler(types.GatewayEventMessageCreate, myHandler); err != nil {
    log.Fatal("handler registration failed:", err)
}
```

---

## discord/channel.go

Lower-level channel functions that take a `*discord.Gateway` directly. The `Client` methods above call these internally.

```go
import "github.com/krishnassh/discoself/discord"
```

All functions mirror their `Client` equivalents but accept `*discord.Gateway` as the first argument instead.

```go
discord.SendMessage(gateway, channelID, content string) bool
discord.SendTyping(gateway, channelID string) bool
discord.DeleteMessage(gateway, channelID, messageID string) bool
discord.EditMessage(gateway, channelID, messageID, content string) bool
discord.SendMessageWithReply(gateway, channelID, content, replyMessageID string) bool
discord.AddReaction(gateway, channelID, messageID, emoji string) bool
```

---

## discord/interactions.go

---

### `GetSlashCommands`

```go
func GetSlashCommands(gateway *Gateway, guildID string) (types.ApplicationCommandIndex, error)
```

Fetches all slash commands available in a guild. Returns an `ApplicationCommandIndex` containing both `Applications []Application` and `ApplicationCommand []ApplicationCommand`.

```go
index, err := discord.GetSlashCommands(gateway, "9876543210987654321")
if err != nil {
    fmt.Println("Error:", err)
    return
}
for _, cmd := range index.ApplicationCommand {
    fmt.Println(cmd.Name, cmd.ApplicationID)
}
```

---

### `GetUserSlashCommands`

```go
func GetUserSlashCommands(gateway *Gateway) (types.ApplicationCommandIndex, error)
```

Fetches slash commands available globally to your user (outside any specific guild).

```go
index, err := discord.GetUserSlashCommands(gateway)
```

---

### `SendSlashCommand`

```go
func SendSlashCommand(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand) bool
```

Lower-level version of `client.SendSlashCommand`. Takes a `*Gateway` directly. Fires the command with an empty options list.

---

### `SendSlashCommandWithOptions`

```go
func SendSlashCommandWithOptions(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand, options []any) bool
```

Fires a slash command with option values. `options` is marshalled to JSON and sent as the command's `options` field.

```go
discord.SendSlashCommandWithOptions(gateway, channelID, guildID, command, []any{"arg1", "arg2"})
```

---

### `ClickButton`

```go
func ClickButton(gateway *Gateway, e *types.MessageEventData, interactionID string) bool
```

Clicks a message component button. `e` is the message event containing the button. `interactionID` is the `CustomID` field of the button component to click.

```go
client.AddHandler(types.GatewayEventMessageCreate, func(e *types.MessageEventData) {
    if len(e.Components) > 0 && len(e.Components[0].Buttons) > 0 {
        discord.ClickButton(client.Gateway, e, e.Components[0].Buttons[0].CustomID)
    }
})
```

Note: `MessageComponent.Buttons` is a `[]types.Buttons` -- the nested field is named `Buttons`, not `Components`.

---

## discord/gateway.go

Internal types and functions. Mostly will be needed directly.

```go
type Gateway struct {
    CloseChan         chan struct{}
    Closed            bool
    Config            *types.Config
    Connection        *websocket.Conn
    GatewayURL        string
    Handlers          Handlers
    LastSeq           int
    Selfbot           *Selfbot
    SessionID         string
    ClientBuildNumber string
}
```

---

### `CreateGateway`

```go
func CreateGateway(selfbot *Selfbot, config *types.Config) *Gateway
```

Creates a `*Gateway`. Called internally by `NewClient`. `config` must not be nil.

---

### `Connect`

```go
func (gateway *Gateway) Connect() error
```

Runs the full connection sequence: dials WebSocket, receives HELLO, starts heartbeat, sends IDENTIFY, waits for READY, then starts the event loop. Called internally by `client.Connect`.

---

### `Close`

```go
func (gateway *Gateway) Close() error
```

Closes the WebSocket connection and signals the event loop to stop.

---

### `GetMembers`

```go
func (gateway *Gateway) GetMembers(id string, ids []string) error
```

Sends a `REQUEST_GUILD_MEMBERS` payload. Called internally by `client.GetMembers`.

---

## discord/handlers.go

### `Add`

```go
func (handlers *Handlers) Add(event string, function any) error
```

Registers a handler on a `*Handlers` instance. Returns an error if the event name is unrecognised or the function signature does not match. Called internally by `client.AddHandler`.

---

## discord/utils.go

### `UtcNow`

```go
func UtcNow() time.Time
```

Returns the current time in UTC.

---

### `TimeSnowflake`

```go
func TimeSnowflake(dt time.Time, high bool) int64
```

Converts a `time.Time` to a Discord snowflake. `high=true` returns the top of that millisecond's range, `high=false` returns the bottom. Useful for timestamp-based message history pagination.

```go
snowflake := discord.TimeSnowflake(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), false)
```

---

### `GenerateNonce`

```go
func GenerateNonce() string
```

Generates a nonce string from the current time as a snowflake. Used internally when sending messages.

---

### `GenerateSessionID`

```go
func GenerateSessionID() string
```

Generates a random 16-character alphanumeric session ID. Used internally during the IDENTIFY handshake.

---

### `GenerateSuperProperties`

```go
func GenerateSuperProperties(gateway *Gateway) string
```

Returns a base64-encoded JSON string for the `X-Super-Properties` header. Built from the gateway's config and the logged-in user's locale. Used internally on every API request.

---

## types/config.go

```go
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
```

`DefaultConfig` is provided as a ready-to-use value:

```go
var DefaultConfig = Config{
    Presence:       "offline",
    ApiVersion:     "10",
    Browser:        "Chrome",
    BrowserVersion: "135.0.0.0",
    Capabilities:   4093,
    UserAgent:      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",
}
```

Pass `&types.DefaultConfig` to `NewClient`. Passing `nil` will panic.

---

## types/events.go

Key event data types passed to handlers:

### `ReadyEventData`

```go
type ReadyEventData struct {
    Version          int     `json:"v"`
    User             User    `json:"user"`
    Guilds           []Guild `json:"guilds"`
    SessionID        string  `json:"session_id"`
    ResumeGatewayURL string  `json:"resume_gateway_url"`
}
```

### `MessageEventData`

```go
type MessageEventData struct {
    MessageData
    ReferencedMessage MessageData `json:"referenced_message"`
}
```

Embeds `MessageData`:

```go
type MessageData struct {
	ID                string             `json:"id"`
	ChannelID         string             `json:"channel_id"`
	GuildID           string             `json:"guild_id"`
	Author            User               `json:"author"`
	Content           string             `json:"content"`
	Timestamp         string             `json:"timestamp"`
	EditedTimestamp   string             `json:"edited_timestamp"`
	Tts               bool               `json:"tts"`
	MentionEveryone   bool               `json:"mention_everyone"`
	Mentions          []User             `json:"mentions"`
	MentionRoles      []string           `json:"mention_roles"`
	MentionChannels   []string           `json:"mention_channels"`
	Attachments       []Attachment       `json:"attachments"`
	Components        []MessageComponent `json:"components"`
	Embeds            []Embed            `json:"embeds"`
	Reactions         []Reaction         `json:"reactions"`
	Nonce             string             `json:"nonce"`
	Pinned            bool               `json:"pinned"`
	WebhookID         string             `json:"webhook_id"`
	Type              int                `json:"type"`
	Activity          MessageActivity    `json:"activity"`
	Application       MessageApplication `json:"application"`
	Flags             int                `json:"flags"`
	ReferencedMessage MessageReference   `json:"referenced_message"`
	Interaction       MessageInteraction `json:"interaction"`
	Thread            Channel            `json:"thread"`
	StickerItems      []StickerItem      `json:"sticker_items"`
}
```

### `MessageDeleteEventData`

```go
type MessageDeleteEventData struct {
    ID        string `json:"id"`
    ChannelID string `json:"channel_id"`
    GuildID   string `json:"guild_id"`
}
```

### `TypingStartEventData`

```go
type TypingStartEventData struct {
    ChannelID string `json:"channel_id"`
    GuildID   string `json:"guild_id"`
    UserID    string `json:"user_id"`
    Timestamp int64  `json:"timestamp"`
}
```

### `PresenceUpdateEventData`

```go
type PresenceUpdateEventData struct {
    User    User   `json:"user"`
    GuildID string `json:"guild_id"`
    Status  string `json:"status"`
}
```

---

## types/discord.go

Key shared types:

### `User`

```go
type User struct {
    ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot,omitempty"`
	System        bool   `json:"system,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Banner        string `json:"banner,omitempty"`
	AccentColor   int    `json:"accent_color,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"`
	Flags         uint64 `json:"flag,omitempty"`
	PremiumType   uint64 `json:"premium_type,omitempty"`
	PublicFlags   uint64 `json:"public_flag,omitempty"`
}
```

### `MessageComponent` and `Buttons`

Message components (e.g. button rows) use these types:

```go
type MessageComponent struct {
    Type    int       `json:"type"`
    Buttons []Buttons `json:"components"`
}

type Buttons struct {
    Type     int         `json:"type,omitempty"`
    Style    int         `json:"style,omitempty"`
    Label    string      `json:"label,omitempty"`
    CustomID string      `json:"custom_id,omitempty"`
    Disabled bool        `json:"disabled,omitempty"`
    Emoji    ButtonEmoji `json:"emoji,omitempty"`
}
```

Note: the nested button slice is accessed as `.Buttons`, not `.Components`, despite the JSON key being `components`.

### `ApplicationCommand`

```go
type ApplicationCommand struct {
    ID            string `json:"id"`
    Type          int    `json:"type"`
    ApplicationID string `json:"application_id"`
    Version       string `json:"version"`
    Name          string `json:"name"`
    Description   string `json:"description"`
}
```

### `ApplicationCommandIndex`

```go
type ApplicationCommandIndex struct {
    Applications       []Application        `json:"applications"`
    ApplicationCommand []ApplicationCommand `json:"application_commands"`
    Version            *string              `json:"version,omitempty"`
}
```