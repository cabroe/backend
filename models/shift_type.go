package models

type ShiftType struct {
	BaseModel
	Name        string     `json:"name" gorm:"not null;unique"`
	Description string     `json:"description" gorm:"type:text"`
	Color       string     `json:"color" gorm:"not null"`
	StartTime   string     `json:"start_time" gorm:"not null"` // Format: "HH:MM"
	EndTime     string     `json:"end_time" gorm:"not null"`   // Format: "HH:MM"
	ShiftDays   []ShiftDay `json:"shift_days,omitempty" swaggerignore:"true"`
}
