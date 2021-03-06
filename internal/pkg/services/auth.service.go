package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/stackbreak/tasx/internal/pkg/models"
	"github.com/stackbreak/tasx/internal/pkg/tokens"
)

var ErrInvalidCredentials = errors.New("auth.service: invalid credentials")

type tokenClaims struct {
	RegisteredClaims *tokens.Registered
	PersonId         int `json:"person_id"`
}

func (s *Services) AuthServiceGetPersonByCredentials(username, password string) (*models.Person, error) {
	person, err := s.repo.Person.GetOne(username)
	if err != nil {
		if errors.Is(err, models.ErrUsernameNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	return person, nil
}

func (s *Services) AuthServiceCreatePerson(person *models.Person) (int, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(person.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	person.Password = string(passHash)
	return s.repo.Person.CreateOne(person)
}

func (s *Services) AuthServiceGenerateToken(username, password string) (string, error) {
	person, err := s.AuthServiceGetPersonByCredentials(username, password)
	if err != nil {
		return "", err
	}

	token, err := s.tokens.CreateWithClaims(&tokenClaims{
		RegisteredClaims: s.tokens.GetRegisteredClaimsBase(-1),
		PersonId:         person.Id,
	})
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func (s *Services) AuthServiceParseToken(tokenStr string) (int, error) {
	var claims tokenClaims

	err := s.tokens.ParseToClaims(tokenStr, &claims)
	if err != nil {
		return -1, err
	}

	return claims.PersonId, nil
}
