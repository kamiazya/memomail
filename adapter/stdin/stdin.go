package stdin

import (
	"io/ioutil"
	"os"

	"github.com/kamiazya/memomail/model/message"
)

func ReadMessage() (msg *message.Message, err error) {
	msg = new(message.Message)
	body, err := ioutil.ReadAll(os.Stdin)
	*msg = body
	return
}
