package editor

import (
	"os"
)

func Vim() Editor {
	return &editor{vim{}}
}

type vim struct{}

func (vim) command() string {
	return "vim"
}

func (vim) options(file *os.File) []string {
	return []string{"-c", "setf text", file.Name()}
}
