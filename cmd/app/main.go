package main

import (
	"github.com/Kevinmajesta/backendpergudanganmi/configs"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/builder"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/entity"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/cache"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/encrypt"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/postgres"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/server"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/token"
)

func main() {
	// Load configurations from .env file
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	// Initialize PostgreSQL database connection
	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	// Initialize Redis cache connection
	redisDB := cache.InitCache(&cfg.Redis)

	// Initialize encryption tool
	encryptTool := encrypt.NewEncryptTool(cfg.Encrypt.SecretKey, cfg.Encrypt.IV)

	// Initialize JWT token use case
	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)

	// Convert configs.Config to *entity.Config
	entityCfg := convertToEntityConfig(cfg)

	// Build public and private routes
	publicRoutes := builder.BuildPublicRoutes(db, redisDB, entityCfg, tokenUseCase, encryptTool)
	privateRoutes := builder.BuildPrivateRoutes(db, redisDB, encryptTool, tokenUseCase)

	// Initialize and run the server
	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Example function to convert configs.Config to *entity.Config
func convertToEntityConfig(cfg *configs.Config) *entity.Config {
	return &entity.Config{
		SMTP: entity.SMTPConfig{
			Host:     cfg.SMTP.Host,
			Port:     cfg.SMTP.Port,
			Password: cfg.SMTP.Password,
		},
		// Add other fields as needed
	}
}
