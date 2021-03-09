package models

type TaskList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	PersonId    int    `json:"person_id,omitempty" db:"person_id"`
}

type Task struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	IsDone      bool   `json:"is_done" db:"is_done"`
}

// many-to-many relation
type TaskItemToList struct {
	Id     int `db:"id"`
	ListId int `db:"task_list_id"`
	TaskId int `db:"task_id"`
}
