## discoself Disboard Bump Example

This example demonstrates how to utilize discoself to automatically send
the Disboard `/bump` slash command every 2 hours, keeping your server
bumped without any manual effort.

**Join [discoself](https://discord.gg/hVGCFanfMC)
discord server for support.**

### Build

This assumes you already have a working Go environment setup and that
discoself is correctly installed on your system.

From within the bump example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

```
Usage of ./bump:
  -t string
        User Token
  -g string
        Guild ID
  -c string
        Channel ID
```

The below example shows how to start the bumper from the bump example folder.

```sh
./bump -t <user-token> -g <guild-id> -c <channel-id> 
```

### How it works

Once connected, the selfbot sends the Disboard `/bump` slash
command and then repeats every 2 hours automatically. It matches the command
by both name **and** Disboard's application ID (`302050872383242240`), so it
will never accidentally trigger another bot's `bump` command if multiple bots
in your server register one.

### Getting Arguments 

You will need to enable **Developer Mode** in Discord to copy IDs.
Go to `Settings → Advanced → Developer Mode` and toggle it on. Then:

- **Guild ID:** Right-click your server icon and select `Copy Server ID`.
- **Channel ID:** Right-click the channel you want to bump in and select `Copy Channel ID`.

- **User Token:** See [this guide](https://gist.github.com/KrishnaSSH/b518ec90cd54f33d70a7d4525e9258a2) for how to obtain your user token
