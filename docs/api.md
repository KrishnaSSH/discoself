# discoself API Reference

Every exported function in the project is documented here. Functions are grouped by file and package.

---

## Table of Contents

- [Client (dgate.go)](#client-dgatego)
  - [NewClient](#newclient)
  - [Connect](#connect)
  - [Close](#close)
  - [AddHandler](#addhandler)
  - [SendMessage](#sendmessage)
  - [SendMessageWithReply](#sendmessagewithreply)
  - [EditMessage](#editmessage)
  - [DeleteMessage](#deletemessage)
  - [SendTyping](#sendtyping)
  - [AddReaction](#addreaction)
  - [SendSlashCommand](#sendslashcommand)
  - [GetMembers](#getmembers)
- [discord/channel.go](#discordchannelgo)
  - [SendMessage](#sendmessage-1)
  - [SendTyping](#sendtyping-1)
  - [DeleteMessage](#deletemessage-1)
  - [EditMessage](#editmessage-1)
  - [SendMessageWithReply](#sendmessagewithreply-1)
  - [AddReaction](#addreaction-1)
- [discord/interactions.go](#discordinteractionsgo)
  - [GetSlashCommands](#getslashcommands)
  - [GetUserSlashCommands](#getuserslashcommands)
  - [SendSlashCommand](#sendslashcommand-1)
  - [SendSlashCommandWithOptions](#sendslashcommandwithoptions)
  - [ClickButton](#clickbutton)
- [discord/gateway.go](#discordgatewaygo)
  - [CreateGateway](#creategateway)
  - [Connect](#connect-1)
  - [Close](#close-1)
  - [GetMembers](#getmembers-1)
- [discord/handlers.go](#discordhandlersgo)
  - [Add](#add)
- [discord/utils.go](#discordutilsgo)
  - [UtcNow](#utcnow)
  - [TimeSnowflake](#timesnowflake)
  - [GenerateNonce](#generatenonce)
  - [GenerateSessionID](#generatesessionid)
  - [GenerateSuperProperties](#generatesuperproperties)

---

## Client (dgate.go)

The `Client` struct is the main entrypoint. All methods below are called on a `*Client` returned by `NewClient`.

---

### `NewClient`

```go
func NewClient(token string, config *types.Config) *Client
```

Creates a new client. `token` is your Discord user token. `config` is optional; pass `nil` to use defaults.

Does not open a connection. Call `Connect` after this.

```go
client := discoself.NewClient("your-user-token", nil)
```

---

### `Connect`

```go
func (client *Client) Connect() error
```

Opens the WebSocket connection to the Discord gateway and runs the login handshake. Returns an error if the connection fails.

```go
if err := client.Connect(); err != nil {
    log.Fatal(err)
}
```

---

### `Close`

```go
func (client *Client) Close()
```

Closes the gateway connection. Defer this right after `Connect`.

```go
defer client.Close()
```

---

### `AddHandler`

```go
func (client *Client) AddHandler(event string, function any) error
```

Registers a callback for a gateway event. `event` is the raw Discord event name. `function` is called each time that event fires.

Common event names:

- `MESSAGE_CREATE`
- `MESSAGE_UPDATE`
- `MESSAGE_DELETE`
- `GUILD_MEMBER_ADD`
- `GUILD_MEMBERS_CHUNK`
- `INTERACTION_CREATE`
- `TYPING_START`
- `READY`

Multiple handlers can be registered for the same event.

```go
client.AddHandler("MESSAGE_CREATE", func(data any) {
    fmt.Println(data)
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

Sends a message as a reply to `replyMessageID`. The reply target must be in the same channel. Returns `true` on success.

```go
client.SendMessageWithReply("1234567890123456789", "got it", "9876543210987654321")
```

---

### `EditMessage`

```go
func (client *Client) EditMessage(channelID string, messageID string, content string) bool
```

Edits the content of one of your own messages. Returns `true` on success.

```go
client.EditMessage("1234567890123456789", "1122334455667788990", "updated")
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

Sends a typing indicator to a channel. Clears automatically after a few seconds. Returns `true` on success.

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

Fires a slash command interaction in a channel. Use `GetSlashCommands` to get a valid `types.ApplicationCommand` to pass here. Returns `true` on success.

For commands that take options, use `discord.SendSlashCommandWithOptions` directly.

```go
client.SendSlashCommand("1234567890123456789", "9876543210987654321", command)
```

---

### `GetMembers`

```go
func (client *Client) GetMembers(guildId string, memberIds []string) error
```

Requests member data for a list of user IDs in a guild. Results come back asynchronously via the `GUILD_MEMBERS_CHUNK` event, not as a return value.

```go
client.GetMembers("9876543210987654321", []string{
    "111111111111111111",
    "222222222222222222",
})

client.AddHandler("GUILD_MEMBERS_CHUNK", func(data any) {
    fmt.Println(data)
})
```

---

## discord/channel.go

Lower-level channel functions that take a `*Gateway` directly. The `Client` methods above call these internally. Use these if you are working at the `discord` package level.

Import:

```go
import "github.com/krishnassh/discoself/discord"
```

---

### `SendMessage`

```go
func SendMessage(gateway *Gateway, channelID string, content string) bool
```

Same as `client.SendMessage` but takes a `*Gateway` directly.

---

### `SendTyping`

```go
func SendTyping(gateway *Gateway, channelID string) bool
```

Same as `client.SendTyping` but takes a `*Gateway` directly.

---

### `DeleteMessage`

```go
func DeleteMessage(gateway *Gateway, channelID string, messageID string) bool
```

Same as `client.DeleteMessage` but takes a `*Gateway` directly.

---

### `EditMessage`

```go
func EditMessage(gateway *Gateway, channelID string, messageID string, content string) bool
```

Same as `client.EditMessage` but takes a `*Gateway` directly.

---

### `SendMessageWithReply`

```go
func SendMessageWithReply(gateway *Gateway, channelID string, content string, replyMessageID string) bool
```

Same as `client.SendMessageWithReply` but takes a `*Gateway` directly.

---

### `AddReaction`

```go
func AddReaction(gateway *Gateway, channelID string, messageID string, emoji string) bool
```

Same as `client.AddReaction` but takes a `*Gateway` directly.

---

## discord/interactions.go

---

### `GetSlashCommands`

```go
func GetSlashCommands(gateway *Gateway, guildID string) (types.ApplicationCommandIndex, error)
```

Fetches all slash commands available in a guild. Returns a `types.ApplicationCommandIndex` you can search to find a command by name before calling `SendSlashCommand`.

```go
commands, err := discord.GetSlashCommands(gateway, "9876543210987654321")
if err != nil {
    log.Fatal(err)
}
for _, cmd := range commands {
    fmt.Println(cmd.Name)
}
```

---

### `GetUserSlashCommands`

```go
func GetUserSlashCommands(gateway *Gateway) (types.ApplicationCommandIndex, error)
```

Fetches slash commands available to your user globally, outside of any specific guild. Useful for DM-level or globally registered commands.

```go
commands, err := discord.GetUserSlashCommands(gateway)
```

---

### `SendSlashCommand`

```go
func SendSlashCommand(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand) bool
```

Lower-level version of `client.SendSlashCommand`. Takes a `*Gateway` directly.

---

### `SendSlashCommandWithOptions`

```go
func SendSlashCommandWithOptions(gateway *Gateway, channelID string, guildID string, command types.ApplicationCommand, options []any) bool
```

Fires a slash command with option values. `options` is a slice of values matching the command's parameters in order. Use this when the slash command requires arguments.

```go
discord.SendSlashCommandWithOptions(
    gateway,
    "1234567890123456789",
    "9876543210987654321",
    command,
    []any{"arg1", "arg2"},
)
```

---

### `ClickButton`

```go
func ClickButton(gateway *Gateway, e *types.MessageEventData, interactionID string) bool
```

Clicks a message component button. `e` is the message event containing the button, `interactionID` is the custom ID of the component to trigger. Use this to interact with bots that send button menus.

```go
client.AddHandler("MESSAGE_CREATE", func(raw any) {
    e := raw.(*types.MessageEventData)
    if len(e.Components) > 0 {
        discord.ClickButton(gateway, e, e.Components[0].Components[0].CustomID)
    }
})
```

---

## discord/gateway.go

Internal gateway functions. Most users will not need these directly but they are documented here for completeness.

---

### `CreateGateway`

```go
func CreateGateway(selfbot *Selfbot, config *types.Config) *Gateway
```

Creates a new `*Gateway` from a `*Selfbot` and config. Called internally by `NewClient`. You only need this if you are constructing a gateway manually without using the `Client`.

---

### `Connect`

```go
func (gateway *Gateway) Connect() error
```

Runs the full gateway connection sequence: dials the WebSocket, receives `HELLO`, starts the heartbeat goroutine, sends `IDENTIFY`, and waits for `READY`. Called internally by `client.Connect`.

---

### `Close`

```go
func (gateway *Gateway) Close() error
```

Closes the gateway WebSocket connection. Called internally by `client.Close`.

---

### `GetMembers`

```go
func (gateway *Gateway) GetMembers(id string, ids []string) error
```

Sends a `REQUEST_GUILD_MEMBERS` payload for the given guild and user ID list. Called internally by `client.GetMembers`.

---

## discord/handlers.go

---

### `Add`

```go
func (handlers *Handlers) Add(event string, function any) error
```

Registers a handler function for a named event on a `*Handlers` instance. Called internally by `client.AddHandler`. Returns an error if registration fails.

---

## discord/utils.go

---

### `UtcNow`

```go
func UtcNow() time.Time
```

Returns the current time in UTC.

```go
t := discord.UtcNow()
```

---

### `TimeSnowflake`

```go
func TimeSnowflake(dt time.Time, high bool) int64
```

Converts a `time.Time` to a Discord snowflake ID. Useful for paginating through message history by timestamp instead of message ID. `high` set to `true` returns the snowflake at the top of that millisecond's range, `false` returns the bottom.

```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
snowflake := discord.TimeSnowflake(t, false)
```

---

### `GenerateNonce`

```go
func GenerateNonce() string
```

Generates a nonce string. Discord uses nonces to match outgoing messages to their `MESSAGE_CREATE` echo from the gateway. Called internally when sending messages.

---

### `GenerateSessionID`

```go
func GenerateSessionID() string
```

Generates a random session ID in the format Discord expects. Used during the `IDENTIFY` handshake and `RESUME` reconnect flow. Called internally by the gateway.

---

### `GenerateSuperProperties`

```go
func GenerateSuperProperties(gateway *Gateway) string
```

Returns a base64-encoded JSON string for the `X-Super-Properties` header. This header is sent with API requests to make them look like they come from a real Discord client. Includes OS, browser, build number, and locale. Build number is fetched live from Discord's client assets.

```go
props := discord.GenerateSuperProperties(gateway)
```