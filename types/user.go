package types

import (
	"fmt"
	"regexp"
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	minFirstNameLen = 3
	minLastNameLen  = 3
	minPasswordLen  = 8
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	// CreatedAT time.Time          `bson:"created_at" json:"created_at"`
	// UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateUserParams struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	// CreatedAT  string `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["first_name"] = fmt.Sprintf("نام حداقل باید %d کاراکتر باشد.", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["last_name"] = fmt.Sprintf("نام خانوادگی حداقل باید %d کاراکتر باشد.", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("رمز عبور حداقل باید %d کاراکتر باشد.", minPasswordLen)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = ("ایمیل صحیح نیست.")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegx.MatchString(e)
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(encpw),
		// CreatedAT: time.Now(),
	}, nil
}
