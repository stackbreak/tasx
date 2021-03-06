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
