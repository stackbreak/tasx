package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

type PgPerson struct {
	db *sqlx.DB
}

func (r *PgPerson) CreatePerson(person *models.Person) (int, error) {
	stmt := fmt.Sprintf(`
		insert into %s
			(name, username, password_hash)
		values
			($1, $2, $3)
		returning id
	`,
		tablePerson)

	row := r.db.QueryRow(
		stmt,
		person.Name,
		person.Username,
		person.Password,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return -1, err // TODO: replace to common error
	}

	return id, nil
}
