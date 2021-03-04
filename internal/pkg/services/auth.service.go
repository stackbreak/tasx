package services

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/cristalhq/jwt/v3"
	"github.com/stackbreak/tasx/internal/pkg/models"
)

var ErrInvalidCredentials = errors.New("auth.service: invalid credentials")

func (s *Services) AuthServiceGetPersonByCredentials(username, password string) (*models.Person, error) {
	person, err := s.repo.Person.GetPerson(username)
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
	passHash, err := generatePasswordHash(person.Password)
	if err != nil {
		return -1, err
	}

	person.Password = passHash
	return s.repo.Person.CreatePerson(person)
}

func (s *Services) AuthServiceGenerateToken(username, password string) (string, error) {
	person, err := s.AuthServiceGetPersonByCredentials(username, password)
	if err != nil {
		return "", err
	}

	token, err := jwtCreateToken(jwtSecret, person.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generatePasswordHash(password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passHash), nil
}

/* -------------------------------------------------------------------------- */
// TODO:
// move to separate pkg

var (
	jwtSecret = []byte(`super@HYPER!secret`)
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func jwtCreateToken(secret []byte, userId int) (string, error) {
	jwtSigner, err := jwt.NewSignerHS(jwt.HS256, secret)
	if err != nil {
		return "", fmt.Errorf("jwt: unable to create signer. %w", err)
	}

	jwtBuilder := jwt.NewBuilder(jwtSigner)

	claims := &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userId,
	}

	token, err := jwtBuilder.Build(claims)
	if err != nil {
		return "", fmt.Errorf("jwt: error building token. %w", err)
	}

	return token.String(), nil
}
