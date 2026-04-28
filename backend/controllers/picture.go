package controllers

import (
	"net/http"

	"pullpicture/backend/models"

	"github.com/gin-gonic/gin"
)

func GetPictures(c *gin.Context) {
	var pictures []models.Picture
	models.DB.Find(&pictures)
	c.JSON(http.StatusOK, gin.H{"data": pictures})
}

func GetPicture(c *gin.Context) {
	var picture models.Picture
	if err := models.DB.Where("id = ?", c.Param("id")).First(&picture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": picture})
}

func CreatePicture(c *gin.Context) {
	var input struct {
		URL      string `json:"url" binding:"required"`
		Title    string `json:"title"`
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	picture := models.Picture{URL: input.URL, Title: input.Title, Category: input.Category}
	models.DB.Create(&picture)

	c.JSON(http.StatusOK, gin.H{"data": picture})
}

func UpdatePicture(c *gin.Context) {
	var picture models.Picture
	if err := models.DB.Where("id = ?", c.Param("id")).First(&picture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input struct {
		URL      string `json:"url"`
		Title    string `json:"title"`
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&picture).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": picture})
}

func DeletePicture(c *gin.Context) {
	var picture models.Picture
	if err := models.DB.Where("id = ?", c.Param("id")).First(&picture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&picture)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
