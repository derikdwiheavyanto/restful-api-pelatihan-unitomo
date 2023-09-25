package handler

import (
	"api/internal/domain/matakuliah"
	"api/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type matakuliahHandler struct {
	s matakuliah.Service
}

func NewMatakuliahHandler(s matakuliah.Service) *matakuliahHandler {
	return &matakuliahHandler{s: s}
}

func (h *matakuliahHandler) Create(c *gin.Context) {
	var matakuliah matakuliah.InputMatakuliah
	err := c.ShouldBindJSON(&matakuliah)
	if err != nil {
		helper.IFErr("Create Matakuliah Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	newMatakuliah, err := h.s.Create(matakuliah)
	if err != nil {
		helper.IFErr("Create Matakuliah Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Create Matakuliah Success", "success", newMatakuliah, c)
}

func (h *matakuliahHandler) Update(c *gin.Context) {
	var inputMatakuliah matakuliah.InputMatakuliah
	err := c.ShouldBindJSON(&inputMatakuliah)
	if err != nil {
		helper.IFErr("Update Matakuliah Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	matakuliah, err := h.s.Update(inputMatakuliah)
	if err != nil {
		helper.IFErr("Update Matakuliah Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Update Matakuliah Success", "success", matakuliah, c)
}

func (h *matakuliahHandler) GetAllData(c *gin.Context) {
	matakuliahs, err := h.s.GetAllData()
	if err != nil {
		helper.IFErr("Get Data Failed", http.StatusInternalServerError, "erorr", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Get Data Success", "success", matakuliahs, c)
}

// func (h *matakuliahHandler) GetDataById(c *gin.Context) {
// 	kodeMK := c.Param("kode_mk")
// 	matakuliah, err := h.s.GetDataById(kodeMK)
// 	if err != nil {
// 		helper.IFErr("Get Data Failed", http.StatusBadRequest, "erorr", err.Error(), c)
// 		return
// 	}
// 	helper.BindJsonWhenSuccess("Get Data Success", "success", matakuliah, c)
// }

func (h *matakuliahHandler) Delete(c *gin.Context) {
	kodeMK := c.Param("kode_mk")
	err := h.s.Delete(kodeMK)
	if err != nil {
		helper.IFErr("Delete Data Failed", http.StatusBadRequest, "erorr", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Delete Data Success", "success", "Delete Data Success", c)
}
