package models

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusArchived  = "archived"
)

type ShiftWeek struct {
	BaseModel
	CalendarWeek int        `json:"calendar_week" gorm:"not null;index"`
	Year         int        `json:"year" gorm:"not null;index"`
	DepartmentID *uint      `json:"department_id"`
	Department   Department `json:"department"`
	ShiftDays    []ShiftDay `json:"shift_days,omitempty" swaggerignore:"true"`
	Status       string     `json:"status" gorm:"type:varchar(20);default:'draft'"`
	Notes        string     `json:"notes" gorm:"type:text"`
}

func (sw *ShiftWeek) IsValidStatus() bool {
	return sw.Status == StatusDraft ||
		sw.Status == StatusPublished ||
		sw.Status == StatusArchived
}
