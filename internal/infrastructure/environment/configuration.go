package environment

import (
	"os"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/domain/services/email"
	"github.com/FabioSebs/NotiService/internal/utils"
)

type Environment struct {
	Cfg config.Config
	Svc Services
}

type Services struct {
	Email email.Email
}

func NewConfiguration() (env Environment) {
	// infrastructure
	database := config.Database{
		ConnString: os.Getenv("database.connection"),
		Port:       utils.GetInt("database.port"),
		User:       os.Getenv("database.user"),
		Password:   os.Getenv("database.password"),
	}

	smtp := config.SMTP{
		Server:   os.Getenv("smtp.server"),
		Port:     os.Getenv("smtp.port"),
		User:     os.Getenv("smtp.user"),
		Password: os.Getenv("smtp.password"),
	}

	config := config.NewConfig(database, smtp)

	// services
	svcs := Services{
		Email: email.NewEmailer(config),
	}

	// return
	env = Environment{
		Cfg: config,
		Svc: svcs,
	}

	return
}
