package mailer

type Config struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	EmailAddress string `yaml:"email-address" mapstructure:"email-address"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}
