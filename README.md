<div align="left">
  <img src="assets/mascot.png" alt="discoself mascot" width="200" />
</div>

# discoself

[![Go Reference](https://img.shields.io/badge/Go%20Reference-pkg.go.dev-00ADD8?logo=go&logoColor=white)](https://pkg.go.dev/github.com/krishnassh/discoself) [![Go Report Card](https://img.shields.io/badge/Go%20Report-A+-brightgreen?logo=go&logoColor=white)](https://goreportcard.com/report/github.com/krishnassh/discoself) [![Code Size](https://img.shields.io/github/languages/code-size/krishnassh/discoself?label=code%20size&logo=github&color=blue)](https://github.com/krishnassh/discoself)


discoself is a [Go](https://golang.org/) package that provides low level bindings to the [Discord](https://discord.com/) client API for selfbots. it is a hard fork of an unmaintained project [discordgo-self](https://github.com/QuartzWarrior/discordgo-self).
## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` will always pull the latest tagged release from the main branch.

```sh
go get github.com/krishnassh/discoself
```

To update to the latest version:

```sh
go get -u github.com/krishnassh/discoself
```

### Usage

Import the package into your project.

```go
import (
	"fmt"
	"log"

	"github.com/krishnassh/discoself/discord"
	"github.com/krishnassh/discoself/types"
)
```

Create a new client and connect to the Discord gateway.

```go
func main() {
client := discord.NewClient("user-token", &types.DefaultConfig)

client.AddHandler(types.GatewayEventReady, func(e *types.ReadyEventData) {
fmt.Println("Logged in as:", e.User.Username)
  }
}
```

See Examples and API Reference below for more detailed information.

## Examples
a list of examples that demonstrate how to use this library can be found [here](https://github.com/KrishnaSSH/discoself/tree/main/examples)

## API Reference

exported functions and types are documented check out [docs/api.md](docs/api.md).

## Contributing

Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be discussed.
- Try to match current naming conventions as closely as possible.
- This package is intended to be a low level direct mapping of the Discord client API, so please avoid adding enhancements outside of that scope without first discussing it.
- Create a Pull Request with your changes against the main branch.

## Disclaimer

discoself interacts with the Discord client API in ways that are outside Discord's official bot platform. Use of selfbots violates [Discord's Terms of Service](https://discord.com/terms). I am not responsible for any misuse of this project or any consequences that may arise from its use.