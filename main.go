package main

import (
	"fmt"
	"syscall"

	"github.com/kamiazya/memomail/adapter/editor"
	"github.com/kamiazya/memomail/adapter/mailer"
	"github.com/kamiazya/memomail/adapter/stdin"
	"github.com/kamiazya/memomail/model/message"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	c, getConfigErr := getConfig()
	if getConfigErr != nil {
		fmt.Println(getConfigErr.Error())
		return
	}
	mlr := mailer.New(c.Mailer)

	var (
		msg *message.Message
		err error
	)
	if terminal.IsTerminal(syscall.Stdin) {
		edtr := editor.Vim()
		msg, err = edtr.WriteMessage()
	} else {
		msg, err = stdin.ReadMessage()
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sendMailErr := mlr.Send(msg)
	if sendMailErr != nil {
		fmt.Println(sendMailErr.Error())
	}
}
