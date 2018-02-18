package editor

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/kamiazya/memomail/model/message"
)

type editor struct {
	app
}

func (e *editor) WriteMessage() (msg *message.Message, err error) {
	msg = new(message.Message)

	file, err := ioutil.TempFile("", "memomail")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cmd := exec.Command(e.command(), e.options(file)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	*msg, err = ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	return msg, nil
}
