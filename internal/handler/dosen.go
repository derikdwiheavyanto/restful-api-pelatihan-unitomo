package handler

import (
	"api/internal/domain/dosen"
	"api/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type dosenHandler struct {
	s dosen.Service
}

func NewDosenHandler(s dosen.Service) *dosenHandler {
	return &dosenHandler{s: s}
}

func (h *dosenHandler) Create(c *gin.Context) {
	var dosen dosen.InputDosen
	err := c.ShouldBindJSON(&dosen)
	if err != nil {
		helper.IFErr("Create Dosen Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	newDosen, err := h.s.Create(dosen)
	if err != nil {
		helper.IFErr("Create Dosen Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Create Dosen Success", "success", newDosen, c)
}

func (h *dosenHandler) Update(c *gin.Context) {
	var inputDosen dosen.InputDosen
	err := c.ShouldBindJSON(&inputDosen)
	if err != nil {
		helper.IFErr("Update Dosen Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.IFErr("Update Dosen Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}

	dosen, err := h.s.Update(id, inputDosen)
	if err != nil {
		helper.IFErr("Update Dosen Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Update Dosen Success", "success", dosen, c)
}

func (h *dosenHandler) GetAllData(c *gin.Context) {
	dosens, err := h.s.GetAllData()
	if err != nil {
		helper.IFErr("Get Data Failed", http.StatusInternalServerError, "erorr", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Get Data Success", "success", dosens, c)
}

func (h *dosenHandler) GetDataById(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.IFErr("Update Dosen Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}
	dosen, err := h.s.GetDataById(id)
	if err != nil {
		helper.IFErr("Get Data Failed", http.StatusBadRequest, "erorr", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Get Data Success", "success", dosen, c)
}

func (h *dosenHandler) Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.IFErr("Update Dosen Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}
	err = h.s.Delete(id)
	if err != nil {
		helper.IFErr("Delete Data Failed", http.StatusBadRequest, "erorr", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Delete Data Success", "success", "Delete Data Success", c)
}
