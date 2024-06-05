package broker_svc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/constants"
	"github.com/FabioSebs/NotiService/internal/domain/controllers"
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

type KafkaService interface {
	ProduceMessages(c echo.Context, message ...kafka.Message) (res constants.DEFAULT_RESPONSE, err error)
	ProduceMessage(c echo.Context) (res constants.DEFAULT_RESPONSE, err error)

	GetConsumer() (reader *kafka.Reader)
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
		conn *kafka.Conn = k.Infra.Broker.Connect()
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

func (k *Kafka) ProduceMessage(c echo.Context) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		conn    *kafka.Conn = k.Infra.Broker.Connect()
		message string      = c.Param("otp")
	)

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
func (k *Kafka) GetConsumer() (reader *kafka.Reader) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", k.Cfg.Kafka.Host, k.Cfg.Kafka.Port)},
		Topic:     k.Cfg.Kafka.Topic,
		Partition: 0,
		MaxBytes:  10e6,
	})
	return
}

func (k *Kafka) ConsumeMessages(c echo.Context) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		conn *kafka.Conn = k.Infra.Broker.Connect()
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
			Topic:     k.Cfg.Kafka.Topic,
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
