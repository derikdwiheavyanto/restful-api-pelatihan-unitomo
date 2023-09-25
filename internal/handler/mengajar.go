package handler

import (
	"api/internal/domain/mengajar"
	"api/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type mengajarHandler struct {
	s mengajar.Service
}

func NewMengajarHandler(s mengajar.Service) *mengajarHandler {
	return &mengajarHandler{s: s}
}

func (h *mengajarHandler) Create(c *gin.Context) {
	var inputMengajar mengajar.InputMengajar
	err := c.ShouldBindJSON(&inputMengajar)
	if err != nil {
		helper.IFErr("Create mengajar Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	newmengajar, err := h.s.Create(inputMengajar)
	if err != nil {
		helper.IFErr("Create mengajar Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Create Matakuliah Success", "success", newmengajar, c)
}

func (h *mengajarHandler) Update(c *gin.Context) {
	var inputMengajar mengajar.InputMengajar

	err := c.ShouldBindJSON(&inputMengajar)
	if err != nil {
		helper.IFErr("Update Mengajar Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	mengajar, err := h.s.Update(inputMengajar)
	if err != nil {
		helper.IFErr("Update Mengajar Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Update Mengajar Success", "success", mengajar, c)
}

func (h *mengajarHandler) GetAllData(c *gin.Context) {
	matakuliahs, err := h.s.GetAllData()
	if err != nil {
		helper.IFErr("Get Data Failed", http.StatusInternalServerError, "erorr", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Get Data Success", "success", matakuliahs, c)
}

// func (h *mengajarHandler) GetDataById(c *gin.Context) {
// 	kodeMK := c.Param("kode_mk")
// 	matakuliah, err := h.s.GetDataById(kodeMK)
// 	if err != nil{
// 		helper.IFErr("Get Data Failed", http.StatusBadRequest, "erorr", err.Error(), c)
// 		return
// 	}
// 	helper.BindJsonWhenSuccess("Get Data Success", "success", matakuliah, c)
// }

func (h *mengajarHandler) Delete(c *gin.Context) {
	idDosen, err := uuid.FromString(c.Param("id_dosen"))
	if err != nil {
		helper.IFErr("Delete Mengajar Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}
	kodeMK := c.Param("kode_mk")
	err = h.s.Delete(idDosen, kodeMK)
	if err != nil {
		helper.IFErr("Delete Mengajar Failed", http.StatusBadRequest, "erorr", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Delete Mengajar Success", "success", "Delete Data Success", c)
}
