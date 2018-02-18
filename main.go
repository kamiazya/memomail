package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/kamiazya/memomail/adapter/editor"
	"github.com/kamiazya/memomail/adapter/mailer"
	"github.com/kamiazya/memomail/adapter/stdin"
	"github.com/kamiazya/memomail/model/message"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// save config.
	defer saveConfig()

	mlr := mailer.New(c.Mailer)

	var (
		msg *message.Message
		err error
	)
	if terminal.IsTerminal(syscall.Stdin) {

		// switch editor
		var edtr editor.Editor
		switch c.Editor {
		case "micro":
			edtr = editor.Micro()
		case "emacs":
			edtr = editor.Emacs()
		default:
			edtr = editor.Vim()
		}

		// write message in the editor
		msg, err = edtr.WriteMessage()
	} else {
		// read message from standard input
		msg, err = stdin.ReadMessage()
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sendMailErr := mlr.Send(msg)
	if sendMailErr != nil {
		fmt.Println(sendMailErr.Error())
		os.Exit(1)
	}
}
