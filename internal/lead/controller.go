package lead

import (
	"crm/internal/user"
	"crm/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req LeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	lead, err := ctrl.service.Create(req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat lead", lead)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)
	leads, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data lead", leads)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseLeadID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	lead, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data lead", lead)
}

func (ctrl *Controller) GetMyLeads(c *gin.Context) {
	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	leads, err := ctrl.service.GetByAssignedTo(userID, tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data lead saya", leads)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseLeadID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req LeadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	lead, err := ctrl.service.Update(id, req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate lead", lead)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseLeadID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	if err := ctrl.service.Delete(id, tenantID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus lead", nil)
}
