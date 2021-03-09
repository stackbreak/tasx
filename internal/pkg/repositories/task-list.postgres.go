package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/stackbreak/tasx/internal/pkg/models"
)

type PgTaskList struct {
	db *sqlx.DB
}

func (r *PgTaskList) CreateOne(personId int, list *models.TaskList) (int, error) {
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

func (r *PgTaskList) GetAll(personId int) ([]models.TaskList, error) {
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

func (r *PgTaskList) GetOne(personId, taskListId int) (*models.TaskList, error) {
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

func (r *PgTaskList) DeleteOne(personId, taskListId int) error {
	stmt := fmt.Sprintf(`
		delete from %s
		where
			person_id = $1
			and id = $2
	`,
		tableTaskList,
	)

	_, err := r.db.Exec(stmt, personId, taskListId)

	return err
}
