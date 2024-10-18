package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          int       `gorm:"primaryKey" json:"-"`
	UUID        uuid.UUID `gorm:"type:char(36);not null;uniqueIndex" json:"uuid"`
	VariantName string    `gorm:"size:255;not null" json:"variant_name" validate:"required,min=2,max=100"`
	Quantity    int       `gorm:"not null" json:"quantity" validate:"required,min=1"`
	ProductID   uint      `gorm:"not null" json:"product_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (variant *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	variant.UUID = uuid.New()
	variant.CreatedAt = time.Now()
	variant.UpdatedAt = time.Now()
	return
}

func (variant *Variant) Validate() error {
	validate := validator.New()
	return validate.Struct(variant)
}
