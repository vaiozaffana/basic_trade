package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	UUID      uuid.UUID `gorm:"type:char(36);not null;uniqueIndex" json:"-"`
	Name      string    `gorm:"size:255;not null" json:"name" validate:"required,min=2,max=100"`
	Email     string    `gorm:"size:255;not null;unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"size:255;not null" json:"password" validate:"required,min=6"`
	Products  []Product `gorm:"foreignKey:AdminID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	admin.UUID = uuid.New()
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()
	return
}

func (admin *Admin) Validate() error {
	validate := validator.New()
	return validate.Struct(admin)
}
