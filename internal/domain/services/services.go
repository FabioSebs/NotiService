package services

import (
	broker_svc "github.com/FabioSebs/NotiService/internal/domain/services/broker"
	email_svc "github.com/FabioSebs/NotiService/internal/domain/services/email"
)

type Services struct {
	Email  email_svc.EmailService
	Broker broker_svc.KafkaService
}
