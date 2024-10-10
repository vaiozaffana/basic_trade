package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" json:"-"`
	UUID      uuid.UUID `gorm:"type:char(36);not null;uniqueIndex" json:"uuid"`
	Name      string    `gorm:"size:255;not null" json:"name" validate:"required,min=2,max=100"`
	ImageURL  string    `gorm:"size:255;not null" json:"image_url" validate:"required,url"`
	AdminID   uint      `json:"admin_id"`
	Admin     Admin     `gorm:"foreignKey:AdminID" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.UUID = uuid.New()
	return
}

func (product *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(product)
}
