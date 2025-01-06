package aplication

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	usersServices "github.com/gody-server/app/users/domain/services"
	usersDatabase "github.com/gody-server/app/users/infraestructure"

	// Replace for dto
	usersModel "github.com/gody-server/app/users/domain/models"
)

// Create a usersDb instance
var usersDb = usersDatabase.NewUsersDb()

// Create a Service instance using the usersDb
var service = usersServices.NewService(usersDb)

// Post Users
// @Summary Post Users
// @Schemes
// @Description Post Users
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body usersModel.UsersCreate true "Body to create Users"
// @Success 200
// @Router /users [Post]
func CreateUsers(c *gin.Context) {
	var users usersModel.UsersCreate
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateUsers(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "users created",
	})
}

// Get Users
// @Summary Get Users
// @Schemes
// @Description Get Users
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /users [Get]
func GetUsers(c *gin.Context) {
	result, err := service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Get Users
// @Summary Get Users
// @Schemes
// @Description Get Users
// @Tags Users
// @Security BearerAuth
// @Param usersId path int true "usersId"
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{usersId} [Get]
func GetOneUsers(c *gin.Context) {
	usersId, err := strconv.Atoi(c.Param("usersId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.GetOneUsers(usersId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Put Users
// @Summary Put Users
// @Schemes
// @Description Put Users
// @Tags Users
// @Security BearerAuth
// @Param usersId path int true "usersId"
// @Accept json
// @Produce json
// @Param Body body usersModel.UsersPut true "Body to update Users"
// @Success 200
// @Router /users/{usersId} [Put]
func UpdateUsers(c *gin.Context) {
	var user usersModel.UsersPut
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersId, err := strconv.Atoi(c.Param("usersId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateUsers(usersId, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "users updated",
	})
}

// Delete Users
// @Summary Delete Users
// @Schemes
// @Description Delete Users
// @Tags Users
// @Security BearerAuth
// @Param usersId path int true "usersId"
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{usersId} [Delete]
func DeleteUsers(c *gin.Context) {
	usersId, err := strconv.Atoi(c.Param("usersId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteUsers(usersId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "users deleted",
	})
}

// Post Users
// @Summary Post Users
// @Schemes
// @Description Post Users
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body usersModel.LoginRequest true "Body to login"
// @Success 200
// @Router /users/login [Post]
func Login(c *gin.Context) {
	var users usersModel.LoginRequest
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.Login(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}
