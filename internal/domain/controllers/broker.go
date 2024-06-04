package controllers

import (
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/FabioSebs/NotiService/internal/infrastructure/broker"
	"github.com/jmoiron/sqlx"
)

type KafkaController interface {
}

type kafkaController struct {
	DB     *sqlx.DB
	Broker broker.KafkaInfra
}

func NewKafkaController(infra infrastructure.Infrastructure) KafkaController {
	return kafkaController{
		DB:     infra.PostgresDB,
		Broker: infra.Broker,
	}
}

func (e *kafkaController) ProduceMessage() {

}

func (e *kafkaController) ConsumeMessage() {

}
