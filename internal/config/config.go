package config

import _ "github.com/joho/godotenv/autoload"

type Config struct {
	Database Database
	SMTP     SMTP
}

type Database struct {
	ConnString string
	Port       int
	User       string
	Password   string
}

type SMTP struct {
	Server   string
	Port     string
	User     string
	Password string
}

func NewConfig(db Database, smtp SMTP) Config {
	return Config{
		db,
		smtp,
	}
}
