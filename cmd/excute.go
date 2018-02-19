package cmd

import (
	"fmt"
	"os"

	"github.com/kamiazya/memomail/adapter/config"
)

func Execute() {
	conf, err := config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd(conf).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
