// filepath: d:\CODING\CRM-SAAS\server\internal\reminder\route.go
package reminder

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, controller *Controller, cfg *config.Config) {
	reminders := r.Group("/reminders")
	reminders.Use(middlewares.Authenticate(cfg))
	{
		reminders.POST("", controller.Create)
		reminders.GET("", controller.GetAll)
		reminders.GET("/my", controller.GetMyReminders)
		reminders.GET("/my/pending", controller.GetMyPendingReminders)
		reminders.GET("/upcoming", controller.GetUpcoming)
		reminders.GET("/:id", controller.GetByID)
		reminders.PUT("/:id", controller.Update)
		reminders.PATCH("/:id/status", controller.UpdateStatus)
		reminders.DELETE("/:id", controller.Delete)
	}
}
