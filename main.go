package main

import (
	"fmt"

	"github.com/kamiazya/memomail/adapter/editor"
	"github.com/kamiazya/memomail/adapter/mailer"
)

func main() {
	c, err := getConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	mlr := mailer.New(c.Mailer)

	edtr := editor.Vim()
	msg, writeMessageErr := edtr.WriteMessage()
	if writeMessageErr != nil {
		fmt.Println(writeMessageErr.Error())
	}

	sendMailErr := mlr.Send(msg)
	if sendMailErr != nil {
		fmt.Println(sendMailErr.Error())
	}

}
