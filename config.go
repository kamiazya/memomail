package main

import (
	"github.com/kamiazya/memomail/adapter/mailer"
)

var c *config

type config struct {
	Editor string         `yaml:"editor"`
	Mailer *mailer.Config `yaml:"mailer"`
}
