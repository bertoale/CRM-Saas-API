package tenant

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToTenantResponse(t *Tenant) *TenantResponse {
	return &TenantResponse{
		ID:        t.ID,
		Name:      t.Name,
		Domain:    t.Domain,
		Status:    string(t.Status),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// ParseTenantID parses tenant ID from URL parameter
func ParseTenantID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid tenant ID")
	}
	return uint(id), nil
}
