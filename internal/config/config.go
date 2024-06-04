package config

type Config struct {
	Database Database
	SMTP     SMTP
	HTTP     HTTP
}

type Database struct {
	ConnString string
	Port       int
	Name       string
	Password   string
	User       string
	Host       string
}

type SMTP struct {
	Server   string
	Port     string
	User     string
	Password string
}

type HTTP struct {
	Host string
	Port string
}

func NewConfig(db Database, smtp SMTP, http HTTP) Config {
	return Config{
		db,
		smtp,
		http,
	}
}
