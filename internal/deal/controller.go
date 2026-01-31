package deal

import (
	"crm/internal/user"
	"crm/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req DealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	deal, err := ctrl.service.Create(req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat deal", deal)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)
	deals, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data deal", deals)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseDealID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	deal, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data deal", deal)
}

func (ctrl *Controller) GetMyDeals(c *gin.Context) {
	userID, _ := user.GetUserIDFromContext(c)
	tenantID, _ := user.GetTenantIDFromContext(c)

	deals, err := ctrl.service.GetByAssignedTo(userID, tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data deal saya", deals)
}

func (ctrl *Controller) GetByStage(c *gin.Context) {
	stageID, err := strconv.ParseUint(c.Param("stage_id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Stage ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	deals, err := ctrl.service.GetByStageID(uint(stageID), tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data deal berdasarkan stage", deals)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseDealID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req DealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	deal, err := ctrl.service.Update(id, req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate deal", deal)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseDealID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	if err := ctrl.service.Delete(id, tenantID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus deal", nil)
}
