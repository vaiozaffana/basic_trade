package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"type:char(36);not null;uniqueIndex" json:"uuid"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	ImageURL  string    `gorm:"size:255;not null" json:"image_url"`
	AdminID   uint      `json:"admin_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.UUID = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return
}

func (product *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(product)
}
