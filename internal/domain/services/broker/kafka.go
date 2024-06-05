package broker_svc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/constants"
	"github.com/FabioSebs/NotiService/internal/domain/controllers"
	"github.com/FabioSebs/NotiService/internal/domain/entity"
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type KafkaService interface {
	ProduceMessages(c echo.Context, message ...kafka.Message) (res constants.DEFAULT_RESPONSE, err error)
	ProduceMessage(c echo.Context, req entity.Message) (res constants.DEFAULT_RESPONSE, err error)

	GetConsumer(topic string) (reader *kafka.Reader)
	ConsumeMessages(c echo.Context) (res constants.DEFAULT_RESPONSE, err error)
	ConsumeMessage(c echo.Context) (res constants.DEFAULT_RESPONSE, err error)
}

type Kafka struct {
	Infra infrastructure.Infrastructure
	Cfg   config.Config
	Ctrls controllers.Controllers
}

func NewKafkaService(cfg config.Config, ctrl controllers.Controllers, infra infrastructure.Infrastructure) KafkaService {
	return &Kafka{
		Infra: infra,
		Cfg:   cfg,
		Ctrls: ctrl,
	}
}

func (k *Kafka) ProduceMessages(c echo.Context, message ...kafka.Message) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		conn *kafka.Conn = k.Infra.Broker.Connect(k.Cfg.Kafka.Topics.OTP)
	)
	// have to open connection everytime

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.WriteMessages(message...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    nil,
	}
	return
}

func (k *Kafka) ProduceMessage(c echo.Context, req entity.Message) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		conn    *kafka.Conn
		message string
	)

	switch c.Param("topic") {
	case k.Cfg.Kafka.Topics.OTP:
		conn = k.Infra.Broker.Connect(k.Cfg.Kafka.Topics.OTP)
		message = req.OTP

	case k.Cfg.Kafka.Topics.Email:
		conn = k.Infra.Broker.Connect(k.Cfg.Kafka.Topics.Email)
		message = req.Email

	case k.Cfg.Kafka.Topics.ICCT:
		conn = k.Infra.Broker.Connect(k.Cfg.Kafka.Topics.ICCT)
		message = req.ICCT
	default:
		res = constants.DEFAULT_RESPONSE{
			Message: constants.STATUS_ERROR_MSG,
			Data:    "invalid topic name",
		}
		return
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    nil,
	}

	return
}
func (k *Kafka) GetConsumer(topic string) (reader *kafka.Reader) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", k.Cfg.Kafka.Host, k.Cfg.Kafka.Port)},
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6,
	})
	return
}

func (k *Kafka) ConsumeMessages(c echo.Context) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		conn *kafka.Conn = k.Infra.Broker.Connect("")
	)

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message

	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    string(b[:]),
	}

	return
}

func (k *Kafka) ConsumeMessage(c echo.Context) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		ctx context.Context = context.Background()

		reader *kafka.Reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{fmt.Sprintf("%s:%s", k.Cfg.Kafka.Host, k.Cfg.Kafka.Port)},
			Topic:     c.Param("topic"),
			Partition: 0,
			MaxBytes:  10e6,
		})

		message kafka.Message
	)

	for {
		message, err = reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    string(message.Value),
	}

	return
}
