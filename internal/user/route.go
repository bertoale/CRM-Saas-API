package user

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	// Auth routes (public)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register-tenant", controller.CreateTenantWithOwner)
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", middlewares.Authenticate(cfg), controller.Logout)
	}

	// User routes (protected)
	userGroup := r.Group("/users")
	userGroup.Use(middlewares.Authenticate(cfg))
	{
		userGroup.POST("", controller.Create)
		userGroup.GET("", controller.GetAll)
		userGroup.GET("/:id", controller.GetByID)
		userGroup.PUT("/:id", controller.Update)
		userGroup.DELETE("/:id", controller.Delete)
	}
}
