package services

import (
	"BasicTrade/database"
	"BasicTrade/models"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

func CreateVariant(variant *models.Variant) error {
	// Validate the variant data before creating it
	if err := variant.Validate(); err != nil {
		return err
	}

	if err := database.DB.Create(variant).Error; err != nil {
		return err
	}
	return nil
}

func GetAllVariants(page, pageSize int, search string) ([]models.Variant, error) {
	var variants []models.Variant
	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Variant{})

	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(variant_name) LIKE ?", searchTerm)
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&variants).Error; err != nil {
		return nil, err
	}

	return variants, nil
}

func GetVariantByUUID(uuid string) (*models.Variant, error) {
	var variant models.Variant
	if err := database.DB.Where("uuid = ?", uuid).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("variant not found")
		}
		return nil, err
	}
	return &variant, nil
}

func UpdateVariant(uuid string, updatedVariant *models.Variant) error {
	var variant models.Variant

	if err := database.DB.Where("uuid = ?", uuid).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("variant not found")
		}
		return err
	}

	if err := updatedVariant.Validate(); err != nil {
		return err
	}

	updatedVariant.ID = variant.ID
	updatedVariant.UUID = variant.UUID
	updatedVariant.ProductID = variant.ProductID
	updatedVariant.CreatedAt = variant.CreatedAt
	updatedVariant.UpdatedAt = time.Now()

	// Lakukan update hanya pada field yang diizinkan
	if err := database.DB.Model(&variant).Updates(map[string]interface{}{
		"variant_name": updatedVariant.VariantName,
		"quantity":     updatedVariant.Quantity,
		"updated_at":   updatedVariant.UpdatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteVariant(uuid string) error {
	var variant models.Variant
	if err := database.DB.Where("uuid = ?", uuid).First(&variant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("variant not found")
		}
		return err
	}

	if err := database.DB.Delete(&variant).Error; err != nil {
		return err
	}
	return nil
}
