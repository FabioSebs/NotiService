package config

type Config struct {
	Database Database
	SMTP     SMTP
	HTTP     HTTP
	Kafka    Kafka
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
	Server     string
	Port       string
	User       string
	Password   string
	Recipients []string
}

type HTTP struct {
	Host string
	Port string
}

type Kafka struct {
	Host   string
	Port   string
	Topics Topics
}

type Topics struct {
	OTP   string
	Email string
	ICCT  string
}

func NewConfig(db Database, smtp SMTP, http HTTP, kafka Kafka) Config {
	return Config{
		db,
		smtp,
		http,
		kafka,
	}
}
