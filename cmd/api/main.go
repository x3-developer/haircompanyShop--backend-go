package main

import (
	"context"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"serv_shop_haircompany/internal/config"
	"serv_shop_haircompany/internal/shared/application/container"
	"serv_shop_haircompany/internal/shared/application/transport/rest"
	"serv_shop_haircompany/internal/shared/utils/logging"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	loadEnv()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := newLogger(cfg)
	diContainer := container.NewContainer(cfg, logger)

	srv := newHTTPServer(diContainer)
	runServer(srv, logger)

	<-ctx.Done()
	gracefulShutdown(srv, &wg, logger)

}

func loadEnv() {
	_ = godotenv.Load(".env")
}

func newLogger(cfg *config.Config) *zap.Logger {
	logger, err := logging.InitLogger(
		cfg.AppLogLvl,
		cfg.AppEnv == "production",
		"logs/app.log",
	)
	if err != nil {
		panic(err)
	}

	return logger
}

func newHTTPServer(diContainer *container.Container) *http.Server {
	r := rest.NewHTTPRouter(diContainer)

	return &http.Server{
		Addr:    ":" + diContainer.Config.AppPort,
		Handler: r,
	}
}

func runServer(srv *http.Server, logger *zap.Logger) {
	go func() {
		logger.Info("starting server on:", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			logger.Info("stopped listening server:", zap.Error(err))
		}
	}()
}

func gracefulShutdown(srv *http.Server, wg *sync.WaitGroup, logger *zap.Logger) {
	logger.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("error shutting down server:", zap.Error(err))
	}

	logger.Info("waiting for background goroutines to finish...")
	wg.Wait()
	logger.Info("server gracefully stopped")
}
