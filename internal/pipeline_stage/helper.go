package pipeline_stage

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToPipelineStageResponse(p *PipelineStage) *PipelineStageResponse {
	return &PipelineStageResponse{
		ID:         p.ID,
		TenantID:   p.TenantID,
		Name:       p.Name,
		OrderIndex: p.OrderIndex,
		IsWon:      p.IsWon,
		IsLost:     p.IsLost,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func ParseStageID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid stage ID")
	}
	return uint(id), nil
}
