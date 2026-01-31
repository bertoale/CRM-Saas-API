package lead

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToLeadResponse(l *Lead) *LeadResponse {
	return &LeadResponse{
		ID:         l.ID,
		TenantID:   l.TenantID,
		Name:       l.Name,
		Email:      l.Email,
		Phone:      l.Phone,
		Source:     l.Source,
		Status:     string(l.Status),
		AssignedTo: l.AssignedTo,
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
	}
}

func ParseLeadID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid lead ID")
	}
	return uint(id), nil
}
