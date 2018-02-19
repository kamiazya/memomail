package mailer

import (
	"github.com/kamiazya/memomail/model/attachment"
	"github.com/kamiazya/memomail/model/message"
)

type Mailer interface {
	Send(msg *message.Message, ats ...*attachment.Attachment) (err error)
}
