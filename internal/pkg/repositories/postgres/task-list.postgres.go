package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/stackbreak/tasx/internal/pkg/models"
)

type TaskListRepo struct {
	db *sqlx.DB
}

func (r *TaskListRepo) CreateOne(personId int, list *models.TaskList) (int, error) {
	stmt := fmt.Sprintf(`
		insert into %s
			(title, description, person_id)
		values
			($1, $2, $3)
		returning id
	`,
		tableTaskList)

	row := r.db.QueryRow(
		stmt,
		list.Title,
		list.Description,
		personId,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err // TODO: replace to common error
	}

	return id, nil
}

func (r *TaskListRepo) GetAll(personId int) ([]models.TaskList, error) {
	stmt := fmt.Sprintf(`
		select
			id,
			title,
			description
		from %s
		where
			person_id = $1
	`,
		tableTaskList,
	)

	var lists []models.TaskList
	err := r.db.Select(&lists, stmt, personId)

	return lists, err
}

func (r *TaskListRepo) GetOne(personId, taskListId int) (*models.TaskList, error) {
	stmt := fmt.Sprintf(`
		select
			id,
			title,
			description
		from %s
		where
			person_id = $1
			and id = $2
	`,
		tableTaskList,
	)

	var oneList models.TaskList
	err := r.db.Get(&oneList, stmt, personId, taskListId)

	return &oneList, err
}

func (r *TaskListRepo) DeleteOne(personId, taskListId int) error {
	stmt := fmt.Sprintf(`
		delete from %s
		where
			person_id = $1
			and id = $2
		returning id
	`,
		tableTaskList,
	)

	var id int
	err := r.db.QueryRow(stmt, personId, taskListId).Scan(&id)
	return err
}

func (r *TaskListRepo) UpdateOne(personId, taskListId int, inputData *models.InputUpdateTaskList) error {
	setVals := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inputData.Title != nil {
		setVals = append(setVals, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inputData.Title)
		argId++
	}

	if inputData.Description != nil {
		setVals = append(setVals, fmt.Sprintf("description=$%d", argId))
		args = append(args, *inputData.Description)
		argId++
	}

	setStmt := strings.Join(setVals, ", ")

	stmt := fmt.Sprintf(`
		update %s
		set %s
		where
			person_id = $%d
			and id = $%d
	`,
		tableTaskList,
		setStmt,
		argId,
		argId+1,
	)

	args = append(args, personId, taskListId)

	_, err := r.db.Exec(stmt, args...)

	return err
}
