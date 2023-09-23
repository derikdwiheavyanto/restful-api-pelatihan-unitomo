package handler

import (
	"api/internal/domain/user"
	"api/internal/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidattionError(err)
		errorMessages := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account Failed !", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed !", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetAllData(c *gin.Context) {
	getUser, err := h.userService.GetAllData()
	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Register Account Failed !", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get Data Success !", http.StatusOK, "success", getUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUsersFakultas(c *gin.Context) {
	newUserFakultas, err := h.userService.GetUsersFakultas()
	if err != nil {
		helper.IFErr("Get Users Fakultas Failed !", http.StatusBadRequest, "error", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Get Users Fakultas Success", "success", newUserFakultas, c)
}
