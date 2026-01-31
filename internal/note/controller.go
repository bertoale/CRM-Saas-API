package note

import (
	"crm/internal/user"
	"crm/pkg/response"
	"fmt"
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
	var req NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	userID, _ := user.GetUserIDFromContext(c)
	
	note, err := ctrl.service.Create(req, tenantID, userID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Berhasil membuat note", note)
}

func (ctrl *Controller) GetAll(c *gin.Context) {
	tenantID, _ := user.GetTenantIDFromContext(c)
	notes, err := ctrl.service.GetAll(tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data note", notes)
}

func (ctrl *Controller) GetByID(c *gin.Context) {
	id, err := ParseNoteID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	note, err := ctrl.service.GetByID(id, tenantID)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data note", note)
}

func (ctrl *Controller) GetByEntity(c *gin.Context) {
	entityType := c.Query("entity_type")
	entityIDStr := c.Query("entity_id")
	
	if entityType == "" || entityIDStr == "" {
		response.Error(c, http.StatusBadRequest, "entity_type dan entity_id wajib diisi")
		return
	}

	var entityID uint
	if _, err := fmt.Sscanf(entityIDStr, "%d", &entityID); err != nil {
		response.Error(c, http.StatusBadRequest, "entity_id tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	notes, err := ctrl.service.GetByEntity(entityType, entityID, tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengambil data note", notes)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := ParseNoteID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	note, err := ctrl.service.Update(id, req, tenantID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil mengupdate note", note)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := ParseNoteID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	tenantID, _ := user.GetTenantIDFromContext(c)
	if err := ctrl.service.Delete(id, tenantID); err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Berhasil menghapus note", nil)
}
