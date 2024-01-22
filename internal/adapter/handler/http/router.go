package http

import (
	"EWallet/internal/adapter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"strings"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	walletHandler WalletHandler,
	transactionHandler TransactionHandler,
) (*Router, error) {
	// Disable debug mode in production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/api/v1")
	{
		wallet := v1.Group("/wallet")
		{
			wallet.POST("/", walletHandler.Create)
			wallet.GET("/:walletId", walletHandler.GetWallet)
			wallet.POST("/:walletId/send", transactionHandler.CreateTransaction)
			wallet.GET("/:walletId/history", transactionHandler.ListTransactions)
		}
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
