package reminder

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToReminderResponse(r *Reminder) *ReminderResponse {
	return &ReminderResponse{
		ID:         r.ID,
		TenantID:   r.TenantID,
		UserID:     r.UserID,
		EntityType: string(r.EntityType),
		EntityID:   r.EntityID,
		RemindAt:   r.RemindAt,
		Status:     string(r.Status),
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func ParseReminderID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid reminder ID")
	}
	return uint(id), nil
}
