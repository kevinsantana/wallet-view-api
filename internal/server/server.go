package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/wallet-view/internal/config"
)

func Run() {

	r := Router()

	ListenAndServe(r)
}

// ListenAndServe starts http server and handles graceful shutdown
func ListenAndServe(srv *fiber.App) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		log.Info("Gracefully shutting down...")

		_ = srv.Shutdown()
	}()

	log.WithField("host", config.Configuration.ApiHost).
		WithField("port", config.Configuration.ApiPort).
		Info("Wallet view api server started")

	srvHost := net.JoinHostPort(config.Configuration.ApiHost, config.Configuration.ApiPort)

	if err := srv.Listen(srvHost); err != nil && err != http.ErrServerClosed {
		log.Panicf("server error: %v", err)
	}
}