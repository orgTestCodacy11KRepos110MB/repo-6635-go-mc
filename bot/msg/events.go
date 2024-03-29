package msg

import "github.com/Tnze/go-mc/chat"

type EventsHandler struct {
	SystemChat        func(msg chat.Message, overlay bool) error
	PlayerChatMessage func(msg chat.Message, validated bool) error
	DisguisedChat     func(msg chat.Message) error
}
