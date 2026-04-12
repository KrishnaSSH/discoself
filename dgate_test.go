package discoself

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/krishnassh/discoself/types"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_Close(t *testing.T) {
	token, err := os.ReadFile("test_token")
	handleErr(err)

	client := NewClient(string(token), &types.DefaultConfig)
	client.AddHandler(types.GatewayEventReady, func(e *types.ReadyEventData) {
		fmt.Printf("%#v", e)
	})

	go func() {
		handleErr(client.Connect())
	}()

	time.Sleep(10 * time.Second)
	client.Close()
}
func Test_Main(t *testing.T) {
	token, err := os.ReadFile("test_token")
	handleErr(err)

	client := NewClient(string(token), &types.DefaultConfig)

	client.AddHandler(types.GatewayEventReady, func(e *types.ReadyEventData) {
		fmt.Printf("%#v\n", e)
	})
	client.AddHandler(types.GatewayEventMessageCreate, func(e *types.MessageEventData) {
		fmt.Printf("%#v\n", e)
	})
	client.AddHandler(types.GatewayEventMessageUpdate, func(e *types.MessageEventData) {
		fmt.Printf("%#v\n", e)
	})

	handleErr(client.Connect())
}
