package services

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (s *Services) AuthServiceCreatePerson(person *models.Person) (int, error) {
	passHash, err := generatePasswordHash(person.Password)
	if err != nil {
		return -1, err
	}

	person.Password = passHash
	return s.repo.Person.CreatePerson(person)
}

func generatePasswordHash(password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passHash), nil
}
