package lead

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	leadGroup := r.Group("/leads")
	leadGroup.Use(middlewares.Authenticate(cfg))
	{
		leadGroup.POST("", controller.Create)
		leadGroup.GET("", controller.GetAll)
		leadGroup.GET("/my-leads", controller.GetMyLeads)
		leadGroup.GET("/:id", controller.GetByID)
		leadGroup.PUT("/:id", controller.Update)
		leadGroup.DELETE("/:id", controller.Delete)
	}
}
