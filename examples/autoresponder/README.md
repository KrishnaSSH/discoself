## discoself Auto Responder Example

This example demonstrates how to utilize discoself to automatically respond
to messages containing a specific trigger phrase. Useful for quick replies,
FAQ automation, or simple chatbot behavior.

**Join [discoself](https://discord.gg/hVGCFanfMC)
discord server for support.**

### Build

This assumes you already have a working Go environment setup and that
discoself is correctly installed on your system.

From within the autoresponder example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

```
Usage of ./autoresponder:
  -t string
        User Token
  -c string
        Channel ID (optional, if not set responds in all channels)
  -r string
        Trigger phrase to respond to (default "!api")
  -R string
        Response message (default "read the api reference here https://github.com/KrishnaSSH/discoself/blob/main/docs/api.md")
```

The below example shows how to start the autoresponder from the autoresponder example folder.

```sh
./autoresponder -t <user-token>
```

To respond to a custom trigger with a custom response:

```sh
./autoresponder -t <user-token> -r "!help" -R "Check the docs at https://github.com/krishnassh/discoself/blob/main/docs/api.md"
```

To only respond in a specific channel:

```sh
./autoresponder -t <user-token> -c <channel-id> -r "ping" -R "pong"
```

### How it works

Once connected, the selfbot listens for all incoming messages. When a message
contains the trigger phrase (case-insensitive), it automatically replies with
the configured response message. The bot will not respond to its own messages
to prevent infinite loops.

Example interaction:

```
when someone types !api in the monitoring channel
Bot responds with read the api reference here https://github.com/krishnassh/discoself/blob/main/docs/api.md
```

### Getting Arguments

You will need to enable **Developer Mode** in Discord to copy IDs.
Go to `Settings -> Advanced -> Developer Mode` and toggle it on.

- **Channel ID:** Right-click the channel and select `Copy Channel ID`.
- **User Token:** See [this guide](https://gist.github.com/KrishnaSSH/b518ec90cd54f33d70a7d4525e9258a2) for how to obtain your user token.
