package handlers

import (
	"github.com/FabioSebs/NotiService/internal/constants"
	"github.com/FabioSebs/NotiService/internal/domain/entity"
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
		req entity.Message
	)

	// TODO: write middleware

	// binding
	if err := c.Bind(&req); err != nil {
		res.Message = constants.STATUS_ERROR_MSG
		res.Data = err.Error()
		return c.JSON(constants.STATUS_ERROR, res)
	}

	// service
	res, err = k.Service.Broker.ProduceMessage(c, req)
	if err != nil {
		res.Message = constants.STATUS_ERROR_MSG
		res.Data = err.Error()
		return c.JSON(constants.STATUS_ERROR, res)
	}

	// response
	return c.JSON(constants.STATUS_SUCCESS, res)
}

func (k *kafkaHandler) Consume(c echo.Context) (err error) {
	var (
		res constants.DEFAULT_RESPONSE
		req entity.Message
	)

	// TODO: write middleware

	// binding
	if err := c.Bind(&req); err != nil {
		res.Message = constants.STATUS_ERROR_MSG
		res.Data = err.Error()
		return c.JSON(constants.STATUS_ERROR, res)
	}

	// service
	res, err = k.Service.Broker.ConsumeMessage(c)
	if err != nil {
		res.Message = constants.STATUS_ERROR_MSG
		res.Data = err.Error()
		return c.JSON(constants.STATUS_ERROR, res)
	}

	return c.JSON(constants.STATUS_SUCCESS, res)
}
