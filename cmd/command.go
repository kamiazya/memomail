package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/kamiazya/memomail/adapter/config"
	"github.com/kamiazya/memomail/adapter/editor"
	"github.com/kamiazya/memomail/adapter/mailer"
	"github.com/kamiazya/memomail/adapter/stdin"
	"github.com/kamiazya/memomail/model/attachment"
	"github.com/kamiazya/memomail/model/message"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func cmd(conf *config.Config) (r *command) {
	r = &command{
		Command: &cobra.Command{
			Use:   "memomail",
			Short: "Memomail is a tool to write a memo to send mail.",
		},
		config: conf,
		mailer: mailer.New(conf.Mailer),
	}

	// add flags
	r.PersistentFlags().
		StringSliceVarP(&r.attachmentPaths, "attach", "a", []string{}, "the path to the attachements.")

	r.Run = r.action()

	return
}

type command struct {
	*cobra.Command
	attachmentPaths []string
	mailer          mailer.Mailer
	config          *config.Config
}

func (c *command) action() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		msg, err := c.message()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := c.send(msg); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func (c *command) message() (msg *message.Message, err error) {
	if terminal.IsTerminal(syscall.Stdin) {
		msg, err = c.writeMessage()
	} else {
		msg, err = c.readMessage()
	}
	return
}

func (c *command) readMessage() (msg *message.Message, err error) {
	// read message from standard input
	msg, err = stdin.ReadMessage()
	return
}

func (c *command) writeMessage() (msg *message.Message, err error) {
	// switch editor
	var edtr editor.Editor
	switch c.config.Editor {
	case "micro":
		edtr = editor.Micro()
	case "emacs":
		edtr = editor.Emacs()
	default:
		edtr = editor.Vim()
	}
	// write message in the editor
	msg, err = edtr.WriteMessage()
	return
}

func (c *command) send(msg *message.Message) error {
	as, err := c.getAttachments()
	if err != nil {
		return err
	}
	return c.mailer.Send(msg, as...)
}

func (c *command) getAttachments() (as []*attachment.Attachment, err error) {
	as = make([]*attachment.Attachment, 0, len(c.attachmentPaths))
	for _, path := range c.attachmentPaths {
		_, err = os.Stat(path)
		if err != nil {
			return nil, err
		}
		as = append(as, &attachment.Attachment{
			Path: path,
		})
	}
	return
}
