package customer

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToCustomerResponse(c *Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:          c.ID,
		TenantID:    c.TenantID,
		Name:        c.Name,
		Email:       c.Email,
		Phone:       c.Phone,
		CompanyName: c.CompanyName,
		LeadID:      c.LeadID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func ParseCustomerID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid customer ID")
	}
	return uint(id), nil
}
