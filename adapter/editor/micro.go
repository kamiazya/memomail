package editor

import "os"

func Micro() Editor {
	return &editor{micro{}}
}

type micro struct{}

func (micro) command() string {
	return "micro"
}

func (micro) options(file *os.File) []string {
	return []string{file.Name()}
}
