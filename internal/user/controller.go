package user

import (
	"crm/pkg/config"
	"crm/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
	cfg     *config.Config
}

func NewController(service Service, cfg *config.Config) *Controller {
	return &Controller{
		service: service,
		cfg:     cfg,
	}
}

func (ctrl *Controller) CreateTenantWithOwner(c *gin.Context) {
	var req CreateTenantWithOwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.service.CreateTenantWithOwner(req.Tenant, req.User)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat tenant dan owner", result)
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ctrl.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Login berhasil", result)
}

func (ctrl *Controller) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" && len(token) > 7 {
		token = token[7:] // Remove "Bearer " prefix
	}

	if err := ctrl.service.Logout(token); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Logout berhasil", nil)
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	currentUser := ctrl.getCurrentUser(c)
	result, err := ctrl.service.Create(req, currentUser)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat user", result)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	currentUser := ctrl.getCurrentUser(c)
	users, err := ctrl.service.GetAll(currentUser)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data user", users)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	currentUser := ctrl.getCurrentUser(c)
	user, err := ctrl.service.GetByID(id, currentUser)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Berhasil mengambil data user", user)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	currentUser := ctrl.getCurrentUser(c)
	user, err := ctrl.service.Update(id, req, currentUser)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate user", user)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	currentUser := ctrl.getCurrentUser(c)
	if err := ctrl.service.Delete(id, currentUser); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus user", nil)
}

func (ctrl *Controller) getCurrentUser(c *gin.Context) *User {
	userID, _ := GetUserIDFromContext(c)
	tenantID, _ := GetTenantIDFromContext(c)
	userRole, _ := GetUserRoleFromContext(c)

	return &User{
		ID:       userID,
		TenantID: tenantID,
		Role:     RoleENUM(userRole),
	}
}
