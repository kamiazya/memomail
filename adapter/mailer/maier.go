package mailer

import "github.com/kamiazya/memomail/model/message"

type Mailer interface {
	Send(msg *message.Message) (err error)
}
