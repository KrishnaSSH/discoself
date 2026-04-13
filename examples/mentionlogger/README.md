## discoself Mention Logger Example

This example demonstrates how to utilize discoself to log every message
that mentions you across all servers, saving them to a file with timestamps,
guild, channel, and author information.

**Join [discoself](https://discord.gg/hVGCFanfMC)
discord server for support.**

### Build

This assumes you already have a working Go environment setup and that
discoself is correctly installed on your system.

From within the mentionlogger example folder, run the below command to compile the
example.

```sh
go build
```

### Usage

```
Usage of ./mentionlogger:
  -t string
        User Token
  -f string
        Log file path (default "mentions.log")
```

The below example shows how to start the mention logger from the mentionlogger example folder.

```sh
./mentionlogger -t <user-token>
```

To write logs to a custom file:

```sh
./mentionlogger -t <user-token> -f /path/to/output.log
```

### How it works

Once connected, the selfbot listens for all incoming messages across every
server you are in. When a message contains a mention matching your user ID,
it appends a line to the log file and prints it to stdout. Each log entry
includes the timestamp, guild ID, channel ID, and the author's username and
discriminator.

Example log output:

```
[2024-03-01 14:22:10] guild:123456789 channel:987654321 author:someone#1234 said: hey <@yourid> check this out
```

### Getting Arguments

You will need to enable **Developer Mode** in Discord to copy IDs.
Go to `Settings -> Advanced -> Developer Mode` and toggle it on.

- **User Token:** See [this guide](https://gist.github.com/KrishnaSSH/b518ec90cd54f33d70a7d4525e9258a2) for how to obtain your user token.