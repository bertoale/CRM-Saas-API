package customer

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
	var req CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	customer, err := ctrl.service.Create(req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat customer", customer)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)
	customers, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data customer", customers)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseCustomerID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	customer, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data customer", customer)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseCustomerID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	customer, err := ctrl.service.Update(id, req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate customer", customer)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseCustomerID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	if err := ctrl.service.Delete(id, tenantID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus customer", nil)
}
