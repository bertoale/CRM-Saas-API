package pipeline_stage

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	stageGroup := r.Group("/pipeline-stages")
	stageGroup.Use(middlewares.Authenticate(cfg))
	{
		stageGroup.POST("", controller.Create)
		stageGroup.GET("", controller.GetAll)
		stageGroup.GET("/:id", controller.GetByID)
		stageGroup.PUT("/:id", controller.Update)
		stageGroup.DELETE("/:id", controller.Delete)
	}
}
