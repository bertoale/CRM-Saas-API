package tenant

import (
	"crm/pkg/config"
	"crm/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, controller *Controller, cfg *config.Config) {
	tenantGroup := r.Group("/tenants")
	tenantGroup.Use(middlewares.Authenticate(cfg))
	{
		tenantGroup.GET("", controller.GetAll)
		tenantGroup.GET("/:id", controller.GetByID)
		tenantGroup.PUT("/:id", controller.Update)
		tenantGroup.DELETE("/:id", controller.Delete)
	}
}
