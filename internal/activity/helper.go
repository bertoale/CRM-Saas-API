package activity

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToActivityResponse(a *Activity) *ActivityResponse {
	var meta map[string]interface{}
	if a.Meta != nil {
		json.Unmarshal(a.Meta, &meta)
	}

	return &ActivityResponse{
		ID:         a.ID,
		TenantID:   a.TenantID,
		UserID:     a.UserID,
		EntityType: string(a.EntityType),
		EntityID:   a.EntityID,
		Action:     a.Action,
		Meta:       meta,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func ParseActivityID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid activity ID")
	}
	return uint(id), nil
}
