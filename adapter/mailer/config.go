package mailer

type Config struct {
	Host         string
	Port         int
	EmailAddress string `mapstructure:"email-address"`
	Username     string
	Password     string
}
