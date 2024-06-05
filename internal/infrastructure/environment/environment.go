package environment

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	broker_process "github.com/FabioSebs/NotiService/internal/broker"
	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/domain/controllers"
	"github.com/FabioSebs/NotiService/internal/domain/services"
	broker_svc "github.com/FabioSebs/NotiService/internal/domain/services/broker"
	email_svc "github.com/FabioSebs/NotiService/internal/domain/services/email"
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/FabioSebs/NotiService/internal/infrastructure/broker"
	"github.com/FabioSebs/NotiService/internal/infrastructure/database"
	"github.com/FabioSebs/NotiService/internal/server"
	"github.com/FabioSebs/NotiService/internal/server/handlers"
	"github.com/FabioSebs/NotiService/internal/utils"
)

type Environment struct {
	Cfg      config.Config
	Svc      services.Services
	Handlers handlers.Handlers
	Server   server.Server
	Broker   broker_process.Broker
}

func NewEnvironment() (env Environment) {
	//////////////////////////////////////////////////////////////////////////////////////////
	///////// configuration ////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////
	db := config.Database{
		ConnString: os.Getenv("database.connection.string"),
		Port:       utils.GetInt(os.Getenv("database.port")),
		Name:       os.Getenv("database.name"),
		User:       os.Getenv("database.user"),
		Password:   os.Getenv("database.password"),
		Host:       os.Getenv("database.host"),
	}

	smtp := config.SMTP{
		Server:   os.Getenv("smtp.server"),
		Port:     os.Getenv("smtp.port"),
		User:     os.Getenv("smtp.user"),
		Password: os.Getenv("smtp.password"),
	}

	http := config.HTTP{
		Host: os.Getenv("api.host"),
		Port: os.Getenv("api.port"),
	}

	kafka := config.Kafka{
		Host:  os.Getenv("kafka.host"),
		Port:  os.Getenv("kafka.port"),
		Topic: os.Getenv("kafka.topic"),
	}

	config := config.NewConfig(
		db,
		smtp,
		http,
		kafka,
	)
	//////////////////////////////////////////////////////////////////////////////////////////
	///////// infrastructure ////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////
	infra := infrastructure.Infrastructure{
		PostgresDB: database.ConnectPostgresDB(
			config.Database.Host,
			config.Database.User,
			config.Database.Password,
			config.Database.Name,
			config.Database.Port,
		),
		Broker: broker.NewKafkaInfra(config.Kafka),
	}

	//////////////////////////////////////////////////////////////////////////////////////////
	///////// controllers ////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////
	ctrls := controllers.Controllers{
		Broker: controllers.NewKafkaController(infra),
		Email:  controllers.NewEmailController(infra),
	}

	//////////////////////////////////////////////////////////////////////////////////////////
	///////// services ////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////
	svcs := services.Services{
		Email:  email_svc.NewEmailService(config, ctrls),
		Broker: broker_svc.NewKafkaService(config, ctrls, infra),
		// add more
	}

	//////////////////////////////////////////////////////////////////////////////////////////
	///////// handlers ////////////////////////////////////////////////////////////////////////
	/////////////////////////////////////////s///////////////////////////////////////////////
	handles := handlers.Handlers{
		EmailHandler: handlers.NewEmailHandler(svcs),
		KafkaHandler: handlers.NewKafkaHandler(svcs),
	}

	//////////////////////////////////////////////////////////////////////////////////////////
	///////// processes + environment ////////////////////////////////////////////////////////////////////////
	////////////////////////// //////////////////////////////////////////////////////////////
	server := server.NewServer(config, handles)
	broker := broker_process.NewBroker(config, svcs.Broker, svcs.Email)

	env = Environment{
		Cfg:      config,
		Svc:      svcs,
		Handlers: handles,
		Server:   server,
		Broker:   broker,
	}
	return
}
