package note

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	noteGroup := r.Group("/notes")
	noteGroup.Use(middlewares.Authenticate(cfg))
	{
		noteGroup.POST("", controller.Create)
		noteGroup.GET("", controller.GetAll)
		noteGroup.GET("/entity", controller.GetByEntity)
		noteGroup.GET("/:id", controller.GetByID)
		noteGroup.PUT("/:id", controller.Update)
		noteGroup.DELETE("/:id", controller.Delete)
	}
}
