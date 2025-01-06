package aplication

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tokenServices "github.com/gody-server/app/token/domain/services"
	tokenDatabase "github.com/gody-server/app/token/infraestructure"

	// Replace for dto
	tokenModel "github.com/gody-server/app/token/domain/models"
)

// Create a tokenDb instance
var tokenDb = tokenDatabase.NewTokenDb()

// Create a Service instance using the tokenDb
var service = tokenServices.NewService(tokenDb)

// Post Token
// @Summary Post Token
// @Schemes
// @Description Post Token
// @Tags Token
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body tokenModel.Token true "Body to create Token"
// @Success 200
// @Router /token [Post]
func CreateToken(c *gin.Context) {
	var token tokenModel.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "token created",
	})
}

// Get Token
// @Summary Get Token
// @Schemes
// @Description Get Token
// @Tags Token
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /token [Get]
func GetToken(c *gin.Context) {
	result, err := service.GetToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Get Token
// @Summary Get Token
// @Schemes
// @Description Get Token
// @Tags Token
// @Security BearerAuth
// @Param tokenId path int true "tokenId"
// @Accept json
// @Produce json
// @Success 200
// @Router /token/{tokenId} [Get]
func GetOneToken(c *gin.Context) {
	tokenId, err := strconv.Atoi(c.Param("tokenId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.GetOneToken(tokenId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// Put Token
// @Summary Put Token
// @Schemes
// @Description Put Token
// @Tags Token
// @Security BearerAuth
// @Param tokenId path int true "tokenId"
// @Accept json
// @Produce json
// @Param Body body tokenModel.Token true "Body to update Token"
// @Success 200
// @Router /token/{tokenId} [Put]
func UpdateToken(c *gin.Context) {
	var token tokenModel.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenId, err := strconv.Atoi(c.Param("tokenId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateToken(tokenId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "token updated",
	})
}

// Delete Token
// @Summary Delete Token
// @Schemes
// @Description Delete Token
// @Tags Token
// @Security BearerAuth
// @Param tokenId path int true "tokenId"
// @Accept json
// @Produce json
// @Success 200
// @Router /token/{tokenId} [Delete]
func DeleteToken(c *gin.Context) {
	tokenId, err := strconv.Atoi(c.Param("tokenId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteToken(tokenId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "token deleted",
	})
}

// Post Token
// @Summary Post Token
// @Schemes
// @Description Post Token
// @Tags Token
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Body body tokenModel.TokenVerify true "Body to verify"
// @Success 200
// @Router /token/verify [Post]
func Verify(c *gin.Context) {
	var token tokenModel.TokenVerify
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Verify(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
}
