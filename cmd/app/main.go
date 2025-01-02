package main

import (
	"log"
	"net/http"
	"proj/internal/common/infrastructure/database"
	commonRouter "proj/internal/common/interfaces/http/router"
	"proj/internal/user/infrastructure/repository/mongodb"
	"proj/pkg/logger"
	"time"

	"proj/internal/config"
	"proj/pkg/tracer"

	userRouter "proj/internal/user/interfaces/http/router"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger.Init(cfg.Environment)
	defer logger.Sync()

	// Даем время контейнерам подняться
	time.Sleep(5 * time.Second)

	if err := tracer.Init(cfg.Telemetry.Host, cfg.Telemetry.Port); err != nil {
		logger.Fatal("Failed to initialize tracer", zap.Error(err))
	}

	logger.Info("Starting application",
		zap.String("environment", cfg.Environment),
		zap.String("port", cfg.Server.Port),
	)

	db, err := database.NewMongoConnection(cfg.MongoDB)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB",
			zap.Error(err),
		)
	}

	userRepo := mongodb.NewUserRepository(db)
	uRouter := userRouter.NewUserRouter(userRepo)
	r := commonRouter.NewRouter(uRouter)

	addr := ":" + cfg.Server.Port
	logger.Info("Server starting", zap.String("address", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
