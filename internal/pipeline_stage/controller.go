package pipeline_stage

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
	var req PipelineStageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	stage, err := ctrl.service.Create(req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat pipeline stage", stage)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)
	stages, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data pipeline stage", stages)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseStageID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	stage, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data pipeline stage", stage)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseStageID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req PipelineStageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	stage, err := ctrl.service.Update(id, req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate pipeline stage", stage)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseStageID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	if err := ctrl.service.Delete(id, tenantID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus pipeline stage", nil)
}
