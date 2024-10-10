// controllers/variant_controller.go
package controllers

import (
	"BasicTrade/models"
	"BasicTrade/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Variant
func CreateVariant(c *gin.Context) {
	var variant models.Variant

	if err := c.ShouldBindJSON(&variant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateVariant(&variant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": variant})
}

func GetAllVariants(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.DefaultQuery("search", "")

	variants, err := services.GetAllVariants(page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"variants": variants})
}

func GetVariantByUUID(c *gin.Context) {
	variantUUID := c.Param("variantUUID")

	variant, err := services.GetVariantByUUID(variantUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": variant})
}

func UpdateVariant(c *gin.Context) {
	variantUUID := c.Param("variantUUID")

	var updatedVariant models.Variant
	if err := c.ShouldBindJSON(&updatedVariant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateVariant(variantUUID, &updatedVariant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedVariant})
}

func DeleteVariant(c *gin.Context) {
	variantUUID := c.Param("variantUUID")

	if err := services.DeleteVariant(variantUUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Variant deleted"})
}
