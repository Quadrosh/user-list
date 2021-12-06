package user

// RepoInterface user interface for interaction with repository
type RepoInterface interface {
	CreateUser(user *User) error
	FindUsersByID(id string) ([]User, error)
	FindUsersByFirstName(firstName string) ([]User, error)
	FindUsersByLastName(lastName string) ([]User, error)
	FindUsersByAge(age int) ([]User, error)
	FilterUsersByRange(dateFrom, dateTo, ageFrom, ageTo int) ([]User, error)
	CreateUserTableIfNotExists() error
}

// ServiceInterface is the interface for interaction of user logic with API
type ServiceInterface interface {
	Create(user *User) error
	Find(property string, value interface{}) ([]User, error)
	Filter(dateFrom, dateTo, ageFrom, ageTo int) ([]User, int, error)
}
