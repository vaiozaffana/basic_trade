package controllers

import (
	"BasicTrade/models"
	"BasicTrade/services"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const MaxFileSize = 5 * 1024 * 1024

func CreateProduct(c *gin.Context) {
	var product models.Product

	adminID, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin ID not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image file"})
		return
	}

	if file.Size > MaxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file size is too large. Maximum file size is 5 MB."})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
		return
	}
	defer src.Close()

	filePath := filepath.Join("temp", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	imageURL, err := services.UploadToCloudinary(filePath, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	product.Name = c.PostForm("name")
	product.ImageURL = c.PostForm("image_url")
	product.ImageURL = imageURL
	product.AdminID = adminID.(uint)

	if err := product.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.DefaultQuery("search", "")

	products, err := services.GetAllProducts(page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductByUUID(c *gin.Context) {
	uuid := c.Param("productUUID")
	product, err := services.GetProductByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
	uuid := c.Param("productUUID")
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateProduct(uuid, &updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	uuid := c.Param("productUUID")
	if err := services.DeleteProduct(uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
