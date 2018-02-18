package editor

import "os"

type app interface {
	command() string
	options(file *os.File) []string
}
