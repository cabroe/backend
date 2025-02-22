package models

type Employee struct {
	BaseModel
	FirstName    string     `json:"first_name" gorm:"not null"`
	LastName     string     `json:"last_name" gorm:"not null"`
	Email        string     `json:"email" gorm:"unique;not null"`
	Password     string     `json:"-" gorm:"not null"`
	Color        string     `json:"color" gorm:"not null"`
	IsAdmin      bool       `json:"is_admin" gorm:"default:false"`
	DepartmentID *uint      `json:"department_id"`
	Department   Department `json:"department"`
	ShiftDays    []ShiftDay `json:"shift_days" swaggerignore:"true"`
}
