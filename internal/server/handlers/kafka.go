package handlers

import (
	"github.com/FabioSebs/NotiService/internal/constants"
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
	var (
		res constants.DEFAULT_RESPONSE
	)

	res, err = k.Service.Broker.ProduceMessage(c)
	if err != nil {
		return
	}

	return c.JSON(constants.STATUS_SUCCESS, res)
}

func (k *kafkaHandler) Consume(c echo.Context) (err error) {
	var (
		res constants.DEFAULT_RESPONSE
	)

	res, err = k.Service.Broker.ConsumeMessage(c)
	if err != nil {
		return
	}

	return c.JSON(constants.STATUS_SUCCESS, res)
}
