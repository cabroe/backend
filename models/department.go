package models

type Department struct {
	BaseModel
	Name        string      `json:"name" gorm:"not null;uniqueIndex"`
	Color       string      `json:"color" gorm:"not null"`
	Description string      `json:"description" gorm:"type:text"`
	Employees   []Employee  `json:"employees,omitempty" swaggerignore:"true"`
	ShiftWeeks  []ShiftWeek `json:"shift_weeks,omitempty" swaggerignore:"true"`
}
