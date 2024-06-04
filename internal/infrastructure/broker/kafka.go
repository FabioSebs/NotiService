package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/segmentio/kafka-go"
)

type KafkaInfra interface {
	Connect() *kafka.Conn
}

type kafkaInfra struct {
	Conn *kafka.Conn
	Cfg  config.Kafka
}

func NewKafkaInfra(cfg config.Kafka) KafkaInfra {
	var (
		ctx       context.Context = context.Background()
		network   string          = "tcp"
		address   string          = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
		topic     string          = cfg.Topic
		partition int             = 0
	)

	conn, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return &kafkaInfra{
		Conn: conn,
		Cfg:  cfg,
	}
}

func (k *kafkaInfra) Connect() *kafka.Conn {
	return k.Conn
}
