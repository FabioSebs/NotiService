package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FabioSebs/NotiService/internal/infrastructure/environment"
	"github.com/labstack/gommon/color"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())

	env := environment.NewEnvironment()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	go func() {
		<-ctx.Done()
	}()

	color.Println(color.Green(fmt.Sprintf("⇨ server up on port %s", env.Cfg.HTTP.Port)))

	env.Server.StartServer()
}
