// filepath: d:\CODING\CRM-SAAS\server\internal\activity\route.go
package activity

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, controller *Controller, cfg *config.Config) {
	activities := r.Group("/activities")
	activities.Use(middlewares.Authenticate(cfg))
	{
		activities.POST("", controller.Create)
		activities.GET("", controller.GetAll)
		activities.GET("/my", controller.GetMyActivities)
		activities.GET("/entity", controller.GetByEntity)
		activities.GET("/:id", controller.GetByID)
		activities.PUT("/:id", controller.Update)
		activities.DELETE("/:id", controller.Delete)
	}
}
