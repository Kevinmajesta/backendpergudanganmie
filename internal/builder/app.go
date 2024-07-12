package builder

import (
	"github.com/Kevinmajesta/backendpergudanganmi/internal/http/handler"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/http/router"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/repository"
	"github.com/Kevinmajesta/backendpergudanganmi/internal/service"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/cache"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/encrypt"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/route"
	"github.com/Kevinmajesta/backendpergudanganmi/pkg/token"

	// "github.com/labstack/echo/"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB, redisDB *redis.Client, tokenUseCase token.TokenUseCase, encryptTool encrypt.EncryptTool) []*route.Route {
	userRepository := repository.NewUserRepository(db, nil)
	userService := service.NewUserService(userRepository, tokenUseCase, encryptTool)
	userHandler := handler.NewUserHandler(userService)

	return router.PublicRoutes(userHandler)
}

func BuildPrivateRoutes(db *gorm.DB, redisDB *redis.Client, encryptTool encrypt.EncryptTool, tokenUseCase token.TokenUseCase) []*route.Route {
	cacheable := cache.NewCacheable(redisDB)
	userRepository := repository.NewUserRepository(db, cacheable)
	userService := service.NewUserService(userRepository, nil, encryptTool)
	userHandler := handler.NewUserHandler(userService)

	return router.PrivateRoutes(userHandler)
}
