package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CompletedAt int64  `json:"completedAt"`
}

type Tmpcmd struct{}
