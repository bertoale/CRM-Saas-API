// filepath: d:\CODING\CRM-SAAS\server\internal\reminder\controller.go
package reminder

import (
	"crm/internal/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req ReminderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	reminder, err := ctrl.service.Create(&req, userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reminder)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseReminderID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	reminder, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)

	reminders, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func (ctrl *Controller) GetMyReminders(c *gin.Context) {
	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	reminders, err := ctrl.service.GetMyReminders(userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func (ctrl *Controller) GetMyPendingReminders(c *gin.Context) {
	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	reminders, err := ctrl.service.GetMyPendingReminders(userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func (ctrl *Controller) GetUpcoming(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from and to parameters are required"})
		return
	}

	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date format, use RFC3339"})
		return
	}

	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date format, use RFC3339"})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	reminders, err := ctrl.service.GetUpcoming(tenantID, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseReminderID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req ReminderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	reminder, err := ctrl.service.Update(id, &req, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func (ctrl *Controller) UpdateStatus(c *gin.Context) {
	id, err := ParseReminderID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req UpdateReminderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	reminder, err := ctrl.service.UpdateStatus(id, req.Status, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseReminderID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)

	if err := ctrl.service.Delete(id, tenantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted successfully"})
}
