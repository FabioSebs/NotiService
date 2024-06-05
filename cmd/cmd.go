package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FabioSebs/NotiService/internal/infrastructure/environment"
	"github.com/labstack/gommon/color"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())

	env := environment.NewEnvironment()

	// go routine for listening for os signal (os.Interrupt)
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals // thread hangs until signal is found and written
		cancel()  // sends message to ctx
	}()

	// start server in seperate thread
	go func() {
		color.Println(color.Green(fmt.Sprintf("⇨ server up on port %s", env.Cfg.HTTP.Port)))
		env.Server.StartServer()
	}()

	// start broker in seperate thread
	go env.Broker.HandleOTPEvent(ctx, cancel)

	// main thread is waiting for os interrupt aka context cancel
	<-ctx.Done()

	// Perform graceful shutdown
	color.Println(color.Red(fmt.Sprintf("⇨ shutting down server on port : %s", env.Cfg.HTTP.Port)))
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	env.Server.ShutDownServer(shutdownCtx)
	color.Println(color.Green("⇨ server successfully shutdown"))
}
