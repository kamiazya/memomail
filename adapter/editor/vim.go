package editor

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/kamiazya/memomail/model/message"
)

func Vim() Editor {
	return new(vim)
}

type vim struct{}

func (e *vim) WriteMessage() (msg *message.Message, err error) {
	msg = new(message.Message)

	file, err := ioutil.TempFile("", "tcho-cli")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("vim", "-c", "setf text", file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	fname := file.Name()

	*msg, err = ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
