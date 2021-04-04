package api

// Simple user model, type is either "candidate"
// or "interviewer".
type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" binding:"required"`
	Mail string `json:"mail" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// Slot model. WeekNumber as the name implies
// is the number of the week in a year (1 to 53),
// leaving the better representation to the front-end.
type Slot struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name" binding:"required"`
	WeekNumber uint   `json:"week" binding:"required"`
	Day        string `json:"day" binding:"required"`
	HourStart  uint   `json:"hour_start" binding:"required"`
	HourEnd    uint   `json:"hour_end" binding:"required"`
}

// Main feature Model, scheduling a candidate
// with one or more interviewers (split by ",")
type ScheduleInput struct {
	Candidate    string `json:"candidate" binding:"required"`
	Interviewers string `json:"interviewers" binding:"required"`
}
