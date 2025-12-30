package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost        = 12
	minFirstNameLent  = 2
	minLastNameLent   = 4
	minPasswordLength = 8
)

type CreateUserParams struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.Firstname) > 0 {
		m["firstName"] = p.Firstname
	}
	if len(p.Lastname) > 0 {
		m["lastName"] = p.Lastname
	}
	return m
}
func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Firstname) < minFirstNameLent {
		errors["firstName"] = fmt.Sprintf("Firstname too short!(it should be %d characters long.)", minFirstNameLent)
	}
	if len(params.Lastname) < minLastNameLent {
		errors["lastName"] = fmt.Sprintf("lastname is too short!(it should be %d characters long.)", minLastNameLent)
	}
	if len(params.Password) < minPasswordLength {
		errors["password"] = fmt.Sprintf("Password too short!(it should be %d characters long.)", minPasswordLength)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("Email is invalid!")
	}
	return errors
}
func isEmailValid(email string) bool {
	var emailRegex = regexp.MustCompile(`^[A-Za-z0-9.!#$%&'*+/=?^_{|}~-]+@[A-Za-z0-9-]+(?:\.[A-Za-z0-9-]+)+$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` //these are maps for bson and json bson is for mongodb usage and omitempty doesnt show that argument to public in json or bson
	Firstname         string        `json:"firstName" bson:"firstName"`
	Lastname          string        `json:"lastName" bson:"lastName"`
	Email             string        `json:"email" bson:"email"`
	EncryptedPassword string        `json:"-" bson:"EncryptedPassword"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Firstname:         params.Firstname,
		Lastname:          params.Lastname,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
