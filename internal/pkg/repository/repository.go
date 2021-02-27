package repository

type Authorization interface { //
}

type TaskList interface { //
}

type TaskListItem interface { //
}

type Repository struct {
	Authorization
	TaskList
	TaskListItem
}

func NewRepository() *Repository {
	return &Repository{}
}
