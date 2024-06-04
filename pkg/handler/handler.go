package handler

import (
	"NotificationOfBirthdays/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	api := router.Group("/api")
	{
		api.GET("/users", h.getUsers)

		subscriptions := api.Group("/subscriptions")
		{
			subscriptions.POST("/", h.createSubscription)
			subscriptions.DELETE("/", h.deleteSubscription)
		}
	}

	return router
}
