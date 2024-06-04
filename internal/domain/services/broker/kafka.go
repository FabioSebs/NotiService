package broker_svc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/domain/controllers"
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/segmentio/kafka-go"
)

type KafkaService interface {
	ProduceMessages(message ...kafka.Message)
	ProduceMessage(message string)
	ConsumeMessages()
	ConsumeMessage()
}

type Kafka struct {
	Conn  *kafka.Conn
	Cfg   config.Config
	Ctrls controllers.Controllers
}

func NewKafkaService(cfg config.Config, ctrl controllers.Controllers, infra infrastructure.Infrastructure) KafkaService {
	return &Kafka{
		Conn:  infra.Broker.Connect(),
		Cfg:   cfg,
		Ctrls: ctrl,
	}
}

func (k *Kafka) ProduceMessages(message ...kafka.Message) {
	k.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err := k.Conn.WriteMessages(message...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := k.Conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func (k *Kafka) ProduceMessage(message string) {
	k.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err := k.Conn.Write([]byte(message))
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	if err := k.Conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func (k *Kafka) ConsumeMessages() {
	k.Conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := k.Conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

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

	if err := k.Conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func (k *Kafka) ConsumeMessage() {
	var (
		ctx context.Context = context.Background()

		reader *kafka.Reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{fmt.Sprintf("%s:%s", k.Cfg.Kafka.Host, k.Cfg.Kafka.Port)},
			Topic:     k.Cfg.Kafka.Topic,
			Partition: 0,
			MaxBytes:  10e6,
		})
	)

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
