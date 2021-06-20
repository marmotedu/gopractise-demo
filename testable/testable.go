package testable

type UserStore interface {
	Create(user *User) error
	Get(name string) (*User, error)
	List() ([]*User, error)
	Delete(name string) error
}

type User struct {
	Name    string
	Address string
}

func CreateUser(ds UserStore, user *User) error {
	return ds.Create(user).Error
}
