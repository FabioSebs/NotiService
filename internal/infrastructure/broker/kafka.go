package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/segmentio/kafka-go"
)

type KafkaInfra interface {
	Connect(topic string) *kafka.Conn
}

type kafkaInfra struct {
	Cfg config.Kafka
}

func NewKafkaInfra(cfg config.Kafka) KafkaInfra {
	return &kafkaInfra{
		Cfg: cfg,
	}
}

func (k *kafkaInfra) Connect(topic string) (conn *kafka.Conn) {
	var (
		ctx       context.Context = context.Background()
		network   string          = "tcp"
		address   string          = fmt.Sprintf("%s:%s", k.Cfg.Host, k.Cfg.Port)
		partition int             = 0
	)

	conn, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader: ", err)
	}

	return
}
