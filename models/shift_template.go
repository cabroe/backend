package models

import (
	"time"
)

// ShiftTemplate repräsentiert eine Vorlage für Schichtpläne
type ShiftTemplate struct {
	BaseModel
	Name         string             `json:"name" gorm:"size:100;not null"`
	Description  string             `json:"description" gorm:"type:text"`
	DepartmentID uint               `json:"department_id" gorm:"not null"`
	Department   Department         `json:"department"`
	ShiftDays    []ShiftTemplateDay `json:"shift_days"`
	Status       string             `json:"status" gorm:"type:enum('draft','active','inactive');default:'draft'"`
	ValidFrom    time.Time          `json:"valid_from" gorm:"not null"`
	ValidUntil   time.Time          `json:"valid_until" gorm:"not null"`
}

// ShiftTemplateDay repräsentiert einen Tag in der Schichtvorlage
type ShiftTemplateDay struct {
	BaseModel
	ShiftTemplateID uint      `json:"shift_template_id" gorm:"not null"`
	ShiftTypeID     uint      `json:"shift_type_id" gorm:"not null"`
	ShiftType       ShiftType `json:"shift_type"`
	WeekDay         int       `json:"week_day" gorm:"type:int;check:week_day >= 0 AND week_day <= 6;not null"`
	Notes           string    `json:"notes" gorm:"type:text"`
}

func (st *ShiftTemplate) IsActive() bool {
	return st.Status == "active"
}

func (st *ShiftTemplate) CanBeModified() bool {
	return st.Status == "draft"
}

func (st *ShiftTemplate) Validate() bool {
	now := time.Now()
	return st.ValidFrom.Before(st.ValidUntil) && st.ValidFrom.After(now)
}
