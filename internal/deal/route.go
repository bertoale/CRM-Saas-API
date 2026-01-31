package deal

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	dealGroup := r.Group("/deals")
	dealGroup.Use(middlewares.Authenticate(cfg))
	{
		dealGroup.POST("", controller.Create)
		dealGroup.GET("", controller.GetAll)
		dealGroup.GET("/my-deals", controller.GetMyDeals)
		dealGroup.GET("/stage/:stage_id", controller.GetByStage)
		dealGroup.GET("/:id", controller.GetByID)
		dealGroup.PUT("/:id", controller.Update)
		dealGroup.DELETE("/:id", controller.Delete)
	}
}
