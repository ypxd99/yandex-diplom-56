package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
	"github.com/ypxd99/yandex-diplom-56/internal/repository/postgres"
	"github.com/ypxd99/yandex-diplom-56/internal/server"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
	"github.com/ypxd99/yandex-diplom-56/internal/transport/handler"
	"github.com/ypxd99/yandex-diplom-56/util"
)

func main() {
	cfg := util.GetConfig()
	util.InitLogger(cfg.Logger)
	logger := util.GetLogger()
	logger.Info("start gophermart service")

	if cfg.Postgres.MakeMigration {
		go makeMegrations()
	}

	var (
		repo repository.GophermartRepo
		err  error
	)
	postgresRepo, err := postgres.Connect(context.Background())
	if err != nil {
		logger.Fatalf("Failed to initialize Postgres: %v", err)
	}
	repo = postgresRepo
	defer repo.Close()

	service := service.InitService(repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service.StartAccrualWorker(ctx)

	h := handler.InitHandler(service)

	router := gin.Default()
	h.InitRoutes(router)

	srv := server.NewServer(router)
	go func() {
		util.GetLogger().Infof("GOPHERMART server listeing at: %s", cfg.Server.ServerAddress)

		err := srv.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			util.GetLogger().Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctxSD, cancelSD := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelSD()
	if err := srv.Stop(ctxSD); err != nil {
		util.GetLogger().Fatalf("Server forced to shutdown: %s", err.Error())
	}
	util.GetLogger().Log(4, "HTTP GOPHERMART service stopped")
}

func makeMegrations() {
	// migrate UP
	util.GetLogger().Info("start migrations")
	err := postgres.MigrateDBUp(context.Background())
	if err != nil {
		util.GetLogger().Error(err)
		return
	}
	util.GetLogger().Info("migrations up")

	// migrate DOWN
	//err = postgres.MigrateDBDown(context.Background())
	//if err != nil {
	//	util.GetLogger().Error(err)
	//	return
	//}
}
