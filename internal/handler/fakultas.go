package handler

import (
	"api/internal/domain/fakultas"
	"api/internal/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type fakultasHandler struct {
	fakultasService fakultas.Service
}

func NewFakultasHandler(fakultasService fakultas.Service) *fakultasHandler {
	return &fakultasHandler{fakultasService: fakultasService}
}

func (h *fakultasHandler) CreateFakultas(c *gin.Context) {
	var input fakultas.InputFakultas

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidattionError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Fakultas Failed!", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newFakultas, err := h.fakultasService.InputFakultas(input)

	if err != nil {
		response := helper.APIResponse("Create Fakultas Failed!", http.StatusBadRequest, "errors", newFakultas)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create Fakultas success", http.StatusOK, "success", newFakultas)
	c.JSON(http.StatusOK, response)
}

func (h *fakultasHandler) GetAllData(c *gin.Context) {
	data, err := h.fakultasService.GetAllData()

	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Get Data Failed", http.StatusInternalServerError, "errors", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(data) == 0 {
		response := helper.APIResponse("Get Data Success", http.StatusOK, "success", []any{})
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.APIResponse("Get Data Success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *fakultasHandler) GetDataById(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))

	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Get Data Failed", http.StatusInternalServerError, "errors", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	newFakultas, err := h.fakultasService.GetDataById(id)

	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Get Data Failed", http.StatusBadRequest, "errors", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if newFakultas.ID.IsNil() {
		response := helper.APIResponse("Get Data Success", http.StatusOK, "success", map[string]any{})
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.APIResponse("Get Data Success", http.StatusOK, "success", newFakultas)
	c.JSON(http.StatusOK, response)
}

func (h *fakultasHandler) Update(c *gin.Context) {
	var input fakultas.UpdateFakultas

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Id Not Found", http.StatusBadGateway, "erorrs", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidattionError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Fakultas Failed!", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fakultas, err := h.fakultasService.Update(id, input)

	if err != nil {
		response := helper.APIResponse("Update Fakultas Failed!", http.StatusUnprocessableEntity, "errors", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Update Fakultas Success!", http.StatusOK, "Success", fakultas)
	c.JSON(http.StatusOK, response)
}

func (h *fakultasHandler) Delete(c *gin.Context) {

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Delete Fakultas Failed!", http.StatusBadRequest, "errors", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.fakultasService.Delete(id)
	if err != nil {
		response := helper.APIResponse("Delete Fakultas Failed!", http.StatusBadRequest, "errors", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete Fakultas success", http.StatusOK, "success", "Delete Fakultas Success")
	c.JSON(http.StatusOK, response)
}

func (h *fakultasHandler) GetTotal(c *gin.Context) {
	total, err := h.fakultasService.GetTotal()
	if err != nil {
		response := helper.APIResponse("Get Total Failed!", http.StatusInternalServerError, "errors", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Get Total Success", http.StatusOK, "success", total)
	c.JSON(http.StatusOK, response)

}
