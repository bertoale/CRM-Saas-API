package note

import (
"fmt"
"strconv"

"github.com/gin-gonic/gin"
)

func ToNoteResponse(n *Note) *NoteResponse {
return &NoteResponse{
ID:         n.ID,
TenantID:   n.TenantID,
EntityType: string(n.EntityType),
EntityID:   n.EntityID,
Content:    n.Content,
CreatedBy:  n.CreatedBy,
CreatedAt:  n.CreatedAt,
UpdatedAt:  n.UpdatedAt,
}
}

func ParseNoteID(c *gin.Context) (uint, error) {
idParam := c.Param("id")
id, err := strconv.ParseUint(idParam, 10, 32)
if err != nil {
return 0, fmt.Errorf("invalid note ID")
}
return uint(id), nil
}
