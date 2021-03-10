package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

type TaskRepo struct {
	db *sqlx.DB
}

func (r *TaskRepo) CreateOne(listId int, task *models.Task) (int, error) {
	tx := r.db.MustBegin()

	stmtCreateItem := fmt.Sprintf(`
		insert into %s
			(title, description)
		values
			($1, $2)
		returning id
	`,
		tableTask,
	)

	itemid := new(int)
	row := tx.QueryRow(stmtCreateItem, task.Title, task.Description)
	err := row.Scan(itemid)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	stmtLinkToList := fmt.Sprintf(`
		insert into %s
			(task_id, task_list_id)
		values
			($1, $2)
	`,
		tableTaskToList,
	)

	_, err = tx.Exec(stmtLinkToList, *itemid, listId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	return *itemid, tx.Commit()
}

func (r *TaskRepo) GetAll(personId, listId int) ([]models.Task, error) {
	stmt := fmt.Sprintf(`
		select
			t.id,
			t.title,
			t.description,
			t.is_done
		from %s t
		join %s ttl
			on t.id = ttl.task_id
		join %s tl
			on tl.id = ttl.task_list_id
		where
			tl.id = $1
			and person_id = $2
	`,
		tableTask,
		tableTaskToList,
		tableTaskList,
	)

	items := make([]models.Task, 0)
	err := r.db.Select(&items, stmt, listId, personId)

	return items, err
}

func (r *TaskRepo) GetOne(personId, taskId int) (*models.Task, error) {
	stmt := fmt.Sprintf(`
		select
			t.id,
			t.title,
			t.description,
			t.is_done
		from %s t
		join %s ttl
			on t.id = ttl.task_id
		join %s tl
			on tl.id = ttl.task_list_id
		where
			t.id = $1
			and person_id = $2
	`,
		tableTask,
		tableTaskToList,
		tableTaskList,
	)

	task := new(models.Task)
	err := r.db.Get(task, stmt, taskId, personId)

	return task, err
}

func (r *TaskRepo) DeleteOne(personId, taskId int) error {
	stmt := fmt.Sprintf(`
		delete from %s t
		using
			%s ttl,
			%s tl
		where
			t.id = $1
			and t.id = ttl.task_id
			and tl.id = ttl.task_list_id
			and tl.person_id = $2
		returning t.id
	`,
		tableTask,
		tableTaskToList,
		tableTaskList,
	)

	var id int
	err := r.db.QueryRow(stmt, taskId, personId).Scan(&id)
	return err
}

func (r *TaskRepo) UpdateOne(personId, taskId int, inputData *models.InputUpdateTask) error {
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

	if inputData.IsDone != nil {
		setVals = append(setVals, fmt.Sprintf("is_done=$%d", argId))
		args = append(args, *inputData.IsDone)
		argId++
	}

	setStmt := strings.Join(setVals, ", ")

	stmt := fmt.Sprintf(`
		update %s t
		set %s
		from
			%s tl,
			%s ttl
		where
			t.id = $%d
			and t.id = ttl.task_id
			and tl.id = ttl.task_list_id
			and tl.person_id = $%d
		returning t.id
	`,
		tableTask,
		setStmt,
		tableTaskList,
		tableTaskToList,
		argId,
		argId+1,
	)

	args = append(args, taskId, personId)

	var id int
	err := r.db.QueryRow(stmt, args...).Scan(&id)
	return err
}
