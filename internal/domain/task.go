package domain

import "time"

type Task struct {
	ID           *int       `json:"id" ksql:"_id"`
	Title        *string    `json:"title" ksql:"title"`
	Date         *time.Time `json:"date" ksql:"date"`
	Class        *string    `json:"class" ksql:"class"`
	Content      *string    `json:"content" ksql:"content"`
	Contempled   *bool      `json:"contempled" ksql:"contempled"`
	Satisfactory *bool      `json:"satisfactory" ksql:"satisfactory"`
	Obs          *string    `json:"obs" ksql:"obs"`
	UserID       *int       `json:"userId" ksql:"user_id"`
}

type CreateForm struct {
	Title   string `json:"title"`
	Class   string `json:"class"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type Agenda []Task

type AgendaMap map[string]Task
