package models

type TaskList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type TaskListInterface interface { //
}

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type TaskInterface interface { //
}

// many-to-many relation
type TaskItemToList struct {
	Id     int
	ListId int
	TaskId int
}
