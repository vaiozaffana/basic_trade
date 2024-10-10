package services

import (
	"BasicTrade/database"
	"BasicTrade/models"
	"strings"
)

func CreateProduct(product *models.Product) error {
	if err := database.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(page, pageSize int, search string) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Product{})

	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchTerm)
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
func GetProductByUUID(uuid string) (models.Product, error) {
	var product models.Product
	if err := database.DB.Where("uuid = ?", uuid).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(uuid string, updatedProduct *models.Product) error {
	var product models.Product
	if err := database.DB.Where("uuid = ?", uuid).First(&product).Error; err != nil {
		return err
	}
	product.Name = updatedProduct.Name
	product.ImageURL = updatedProduct.ImageURL
	if err := database.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(uuid string) error {
	if err := database.DB.Where("uuid = ?", uuid).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
