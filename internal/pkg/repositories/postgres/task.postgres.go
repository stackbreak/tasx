package postgres

import (
	"fmt"

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
