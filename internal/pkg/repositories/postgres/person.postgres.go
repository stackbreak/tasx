package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

type PersonRepo struct {
	db *sqlx.DB
}

func (r *PersonRepo) CreateOne(person *models.Person) (int, error) {
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

func (r *PersonRepo) GetOne(username string) (*models.Person, error) {
	stmt := fmt.Sprintf(`
		select
			*
		from %s
		where
			username=$1
	`,
		tablePerson)

	var person models.Person
	err := r.db.Get(&person, stmt, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUsernameNotFound
		}
		return nil, err
	}

	return &person, nil
}
