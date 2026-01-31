// filepath: d:\CODING\CRM-SAAS\server\internal\activity\controller.go
package activity

import (
	"crm/internal/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req ActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	activity, err := ctrl.service.Create(&req, userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseActivityID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	activity, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)

	activities, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (ctrl *Controller) GetByEntity(c *gin.Context) {
	entityType := c.Query("entity_type")
	if entityType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "entity_type is required"})
		return
	}

	entityID := c.Query("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "entity_id is required"})
		return
	}

	var entityIDUint uint
	if _, err := fmt.Sscanf(entityID, "%d", &entityIDUint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entity_id"})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	activities, err := ctrl.service.GetByEntity(entityType, entityIDUint, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (ctrl *Controller) GetMyActivities(c *gin.Context) {
	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	activities, err := ctrl.service.GetByUserID(userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseActivityID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req ActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	activity, err := ctrl.service.Update(id, &req, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseActivityID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	if err := ctrl.service.Delete(id, tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
