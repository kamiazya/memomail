package editor

import "github.com/kamiazya/memomail/model/message"

type Editor interface {
	WriteMessage() (msg *message.Message, err error)
}
