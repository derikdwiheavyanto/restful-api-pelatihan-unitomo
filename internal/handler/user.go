package handler

import (
	"api/internal/domain/user"
	"api/internal/helper"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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

func (h *userHandler) Update(c *gin.Context) {

	var input user.UpdateInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.IFErr("Update Failed !", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	newUser, err := h.userService.Update(input)
	if err != nil {
		helper.IFErr("Update Failed !", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	helper.BindJsonWhenSuccess("Update Users Success", "Success", newUser, c)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	form, err := c.MultipartForm()
	files := form.File["avatar"]
	id := form.Value["id"]
	fmt.Println("id: ", id[0])
	userID := uuid.FromStringOrNil(id[0])
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		helper.IFErr("Upload Avatar Failed", http.StatusUnprocessableEntity, "error", err.Error(), c)
		return
	}

	file := files[0]

	t := time.Now()
	timeFormat := t.Format("20060102150405")
	path := "images/" + timeFormat + "-" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		helper.IFErr("Upload Avatar Failed", http.StatusUnprocessableEntity, "error", data, c)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		helper.IFErr("Upload Avatar Failed", http.StatusUnprocessableEntity, "error", data, c)
		return
	}
	data := gin.H{"is_uploaded": true}
	helper.BindJsonWhenSuccess("Upload Avatar Success", "success", data, c)
}

func (h *userHandler) Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		helper.IFErr("Delete User Failed", http.StatusBadRequest, "error", err.Error(), c)
		return
	}
	err = h.userService.Delete(id)
	if err != nil {
		helper.IFErr("Delete User Failed", http.StatusBadRequest, "erorr", err.Error(), c)
		return
	}
	helper.BindJsonWhenSuccess("Delete User Success", "success", "Delete Data Success", c)
}
