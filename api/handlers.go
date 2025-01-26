package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if err := h.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) EditUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if err := h.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
