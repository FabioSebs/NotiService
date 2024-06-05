package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/FabioSebs/NotiService/internal/config"
	broker_svc "github.com/FabioSebs/NotiService/internal/domain/services/broker"
	email_svc "github.com/FabioSebs/NotiService/internal/domain/services/email"
	"github.com/labstack/gommon/color"
)

type Broker struct {
	Cfg      config.Kafka
	Svc      broker_svc.KafkaService
	EmailSvc email_svc.EmailService
}

func NewBroker(cfg config.Config, svc broker_svc.KafkaService, e_svc email_svc.EmailService) Broker {
	return Broker{
		Cfg:      cfg.Kafka,
		Svc:      svc,
		EmailSvc: e_svc,
	}
}

func (b *Broker) RunConsumer(pipe chan string, cancel context.CancelFunc) {
	var (
		ctx    context.Context = context.Background()
		reader                 = b.Svc.GetConsumer()
	)
	defer reader.Close()

	for {
		select {
		case <-ctx.Done():
			color.Println(color.Red("broker shutting down ..."))
			return
		default:
			message, err := reader.ReadMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					cancel()
					return
				}
				fmt.Println("error encountered: " + err.Error())
				time.Sleep(time.Second) // wait before retrying
				continue
			}
			fmt.Println("received message: " + string(message.Value))
			pipe <- string(message.Value)
			continue
		}
	}
}

func (b *Broker) HandleOTPEvent(ctx context.Context, cancel context.CancelFunc) {
	pipe := make(chan string) // making buffer to get otp value from

	defer close(pipe) // must close channel so new values can be inserted and read !
	// if not close channel then will block forever!

	go b.RunConsumer(pipe, cancel)

	// reading from channel
	for {
		select {

		case <-ctx.Done():
			color.Println(color.Red("handlers shutting down ..."))
			cancel()
			return

		case otp := <-pipe:
			fmt.Println("Processing OTP: " + otp)
			if _, err := b.EmailSvc.SendHTML(otp); err != nil {
				color.Println(color.Red("problem encountered sending email"))
			}
		}
	}
}