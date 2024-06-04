package handlers

import (
	"github.com/FabioSebs/NotiService/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type KafkaHandler interface {
	Produce(c echo.Context) (err error)
	Consume(c echo.Context) (err error)
}

type kafkaHandler struct {
	Service services.Services
}

func NewKafkaHandler(svc services.Services) KafkaHandler {
	return &kafkaHandler{
		Service: svc,
	}
}

func (k *kafkaHandler) Produce(c echo.Context) (err error) {
	return
}

func (k *kafkaHandler) Consume(c echo.Context) (err error) {
	return
}
