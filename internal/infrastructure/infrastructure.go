package infrastructure

import (
	"github.com/FabioSebs/NotiService/internal/infrastructure/broker"
	"github.com/jmoiron/sqlx"
)

type Infrastructure struct {
	PostgresDB *sqlx.DB
	Broker     broker.KafkaInfra
}
