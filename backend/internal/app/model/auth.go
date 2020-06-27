package model

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// Validate ...
func (a *Auth) Validate() error {

	a.Email = strings.Trim(a.Email, " \t\r\n")
	if err := validation.ValidateStruct(
		a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.By(requiredIf(a.EncryptedPassword == "")), validation.Length(8, 100)),
	); err != nil {
		return errors.New("not valid data")
	}
	return nil
}

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

// BeforeCreate ...
func (a *Auth) BeforeCreate() error {

	if len(a.Password) > 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		a.EncryptedPassword = string(hash)
		a.Password = ""
	}
	return nil
}

// ComparePassword ...
func (a *Auth) ComparePassword(pwd string) error {

	err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}
