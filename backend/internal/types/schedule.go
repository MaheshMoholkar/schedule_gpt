package types

type Subject struct {
	TeacherName string `json:"teacher_name"`
	SubjectName string `json:"subject_name"`
	Occurrence  int    `json:"occurrence"`
}

type TeacherSubject struct {
	Subjects []Subject `json:"subjects"`
}

type DailySchedule struct {
	DayName  string    `json:"day_name"`
	Subjects []Subject `json:"subjects"`
}

type WeeklySchedule struct {
	DailySchedule []DailySchedule `json:"daily_schedule"`
}
