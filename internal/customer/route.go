package customer

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	customerGroup := r.Group("/customers")
	customerGroup.Use(middlewares.Authenticate(cfg))
	{
		customerGroup.POST("", controller.Create)
		customerGroup.GET("", controller.GetAll)
		customerGroup.GET("/:id", controller.GetByID)
		customerGroup.PUT("/:id", controller.Update)
		customerGroup.DELETE("/:id", controller.Delete)
	}
}
