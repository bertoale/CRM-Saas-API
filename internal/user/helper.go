package user

import (
	"crm/pkg/config"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID       uint     `json:"id"`
	TenantID uint     `json:"tenant_id"`
	Role     RoleENUM `json:"role"`
	jwt.RegisteredClaims
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		TenantID:  u.TenantID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      string(u.Role),
		Status:    string(u.Status),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// GenerateToken generates JWT token for user
func GenerateToken(user *User, cfg *config.Config) (string, error) {
	duration, err := time.ParseDuration(cfg.JWTExpires)
	if err != nil {
		duration = 168 * time.Hour // default 7 days
	}
	claims := Claims{
		ID:       user.ID,
		TenantID: user.TenantID,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// ParseUserID parses user ID from URL parameter
func ParseUserID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// GetUserIDFromContext retrieves user ID from gin context
func GetUserIDFromContext(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}
	uid, ok := userID.(uint)
	return uid, ok
}

// GetTenantIDFromContext retrieves tenant ID from gin context
func GetTenantIDFromContext(c *gin.Context) (uint, bool) {
	tenantID, exists := c.Get("tenantID")
	if !exists {
		return 0, false
	}
	tid, ok := tenantID.(uint)
	return tid, ok
}

// GetUserRoleFromContext retrieves user role from gin context
func GetUserRoleFromContext(c *gin.Context) (string, bool) {
	userRole, exists := c.Get("userRole")
	if !exists {
		return "", false
	}
	role, ok := userRole.(string)
	return role, ok
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
