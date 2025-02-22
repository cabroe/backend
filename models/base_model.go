package models

import (
	"time"

	"gorm.io/gorm"
)

type DeletedAt = gorm.DeletedAt

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement" swaggertype:"integer"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP" swaggertype:"string" format:"date-time"`
	DeletedAt DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" format:"date-time"`
	CreatedBy uint      `json:"created_by,omitempty" gorm:"default:null" swaggertype:"integer"`
	UpdatedBy uint      `json:"updated_by,omitempty" gorm:"default:null" swaggertype:"integer"`
	Version   uint      `json:"version" gorm:"default:1" swaggertype:"integer"`
}
