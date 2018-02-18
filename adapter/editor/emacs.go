package editor

import "os"

func Emacs() Editor {
	return &editor{emacs{}}
}

type emacs struct{}

func (emacs) command() string {
	return "emacs"
}

func (emacs) options(file *os.File) []string {
	return []string{file.Name()}
}
