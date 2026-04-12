package discord

import (
	"errors"
	"sync"

	"github.com/krishnassh/discoself/types"
)

type Handlers struct {
	OnReady             []func(data *types.ReadyEventData)
	OnMessageCreate     []func(data *types.MessageEventData)
	OnMessageUpdate     []func(data *types.MessageEventData)
	OnGuildMembersChunk []func(data *types.GuildMembersChunkEventData)
	OnReconnect         []func()
	mutex               sync.Mutex
	OnInvalidated       []func()
}

func (handlers *Handlers) Add(event string, function any) error {
	handlers.mutex.Lock()
	defer handlers.mutex.Unlock()

	failed := false

	switch event {
	case types.GatewayEventReady:
		if function, ok := function.(func(data *types.ReadyEventData)); ok {
			handlers.OnReady = append(handlers.OnReady, function)
		} else {
			failed = true
		}
	case types.GatewayEventMessageCreate:
		if function, ok := function.(func(data *types.MessageEventData)); ok {
			handlers.OnMessageCreate = append(handlers.OnMessageCreate, function)
		} else {
			failed = true
		}
	case types.GatewayEventMessageUpdate:
		if function, ok := function.(func(data *types.MessageEventData)); ok {
			handlers.OnMessageUpdate = append(handlers.OnMessageUpdate, function)
		} else {
			failed = true
		}
	case types.GatewayEventInvalidated:
		if function, ok := function.(func()); ok {
			handlers.OnInvalidated = append(handlers.OnInvalidated, function)
		} else {
			failed = true
		}
	case types.EventNameGuildMembersChunk:
		if function, ok := function.(func(data *types.GuildMembersChunkEventData)); ok {
			handlers.OnGuildMembersChunk = append(handlers.OnGuildMembersChunk, function)
		} else {
			failed = true
		}
	case types.GatewayEventReconnect:
		if function, ok := function.(func()); ok {
			handlers.OnReconnect = append(handlers.OnReconnect, function)
		} else {
			failed = true
		}
	default:
		return errors.New("failed to match event to gateway event")
	}

	if failed {
		return errors.New("function signature was not correct for the specified event")
	}

	return nil
}
