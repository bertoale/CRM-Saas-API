package deal

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToDealResponse(d *Deal) *DealResponse {
	return &DealResponse{
		ID:                d.ID,
		TenantID:          d.TenantID,
		CustomerID:        d.CustomerID,
		Title:             d.Title,
		Value:             d.Value,
		StageID:           d.StageID,
		Probability:       d.Probability,
		ExpectedCloseDate: d.ExpectedCloseDate,
		AssignedTo:        d.AssignedTo,
		CreatedAt:         d.CreatedAt,
		UpdatedAt:         d.UpdatedAt,
	}
}

func ParseDealID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid deal ID")
	}
	return uint(id), nil
}
