package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/kamiazya/memomail/adapter/mailer"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func New() (config *Config, err error) {
	config = new(Config)
	config.setDefault()
	if err = config.readConfig(); err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	Editor string         `yaml:"editor"`
	Mailer *mailer.Config `yaml:"mailer"`
}

func (c *Config) setDefault() {
	// editor
	viper.SetDefault("editor", "vim")

	// mailer
	viper.SetDefault("mailer.host", "localhost")
	viper.SetDefault("mailer.port", 25)
	viper.SetDefault("mailer.email-address", "test@example.com")
	viper.SetDefault("mailer.username", "")
	viper.SetDefault("mailer.password", "")
}

func (c *Config) readConfig() error {
	// file
	viper.SetConfigName("memomail")
	viper.SetConfigType("yaml")

	// config dir
	viper.AddConfigPath("/etc/memomail/")         // global
	viper.AddConfigPath("$HOME/.config/memomail") // user
	viper.AddConfigPath(".")                      // on directory

	// read config
	viper.ReadInConfig()
	err := viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SaveConfig() {
	usr, err := user.Current()
	if err != nil {
		return
	}

	dirPath := filepath.Join(usr.HomeDir, ".config/memomail")
	// create config file if file not exist.
	if _, err := os.Stat(dirPath); err != nil {
		e := os.MkdirAll(dirPath, 0755)
		if e != nil {
			fmt.Println(e.Error())
		}
	}

	filePath := filepath.Join(dirPath, "memomail.yml")
	if _, fileFindErr := os.Stat(filePath); fileFindErr != nil {
		// create config file if file not exist.
		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		if d, err := yaml.Marshal(c); err == nil {
			file.Write(d)
		}
	}
}
