package user

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrUserNotFound error of not found user
	ErrUserNotFound = errors.New("User Not Found")
	// ErrRequestInvalid error of invalid request
	ErrRequestInvalid = errors.New("Request Invalid")
	// ErrServerError server error
	ErrServerError = errors.New("Server Error")
)

type service struct {
	userlistRepo RepoInterface
}

// NewService returns service, which implements ServiceInterface inteface
func NewService(repo RepoInterface) ServiceInterface {
	return &service{
		repo,
	}
}

// Create creates new User
func (r *service) Create(user *User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrRequestInvalid, "Invalid request")
	}
	if user.ID == "" {
		user.ID = shortid.MustGenerate()
	}
	user.RecordingDate = time.Now().UTC().Unix()
	response := r.userlistRepo.CreateUser(user)
	return response
}

// Find returns list or users by property
func (r *service) Find(property string, value interface{}) ([]User, error) {

	if property != "id" &&
		property != "first_name" &&
		property != "last_name" &&
		property != "age" {
		return nil, errs.Wrap(ErrRequestInvalid, "Invalid property! (Valid properties: id, first_name, last_name, age)")
	}
	switch property {
	case "id":
		users, err := r.userlistRepo.FindUsersByID(fmt.Sprintf("%s", value))
		return users, err
	case "first_name":
		users, err := r.userlistRepo.FindUsersByFirstName(fmt.Sprintf("%s", value))
		return users, err
	case "last_name":
		users, err := r.userlistRepo.FindUsersByLastName(fmt.Sprintf("%s", value))
		return users, err
	case "age":
		var intVal int
		switch value.(type) {
		case string:
			intVal, _ = strconv.Atoi(fmt.Sprintf("%s", value))
		default:
			intVal = int(value.(float64))
		}
		users, err := r.userlistRepo.FindUsersByAge(intVal)
		return users, err
	}
	return nil, nil
}

// Find returns list or users by property
func (r *service) Filter(dateFrom, dateTo, ageFrom, ageTo int) ([]User, int, error) {

	users, err := r.userlistRepo.FilterUsersByRange(dateFrom, dateTo, ageFrom, ageTo)

	return users, len(users), err

}
