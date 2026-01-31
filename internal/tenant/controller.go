package tenant

import (
	"crm/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenants, err := ctrl.service.GetAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Berhasil mengambil data tenant", tenants)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseTenantID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenant, err := ctrl.service.GetByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data tenant", tenant)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseTenantID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req TenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenant, err := ctrl.service.Update(id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate tenant", tenant)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseTenantID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	if err := ctrl.service.Delete(id); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus tenant", nil)
}
