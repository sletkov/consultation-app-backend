package models

type Consultation struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description,omitempty""`
	Format        string `json:"format"`
	Type          string `json:"type"`
	TeacherName   string `json:"teacher_name"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Campus        string `json:"campus"`
	Classroom     string `json:"classroom"`
	Link          string `json:"link,omitempty" db:"link"`
	Limit         int    `json:"limit, omitempty"`
	StudentsCount int    `json:"students_count"`
}
