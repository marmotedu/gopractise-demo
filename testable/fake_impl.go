package testable

type fake struct {
}

func NewFake() *fake {
	return &fake{}
}

func (f *fake) Create(user *User) error {
	return nil
}

func (f *fake) Get(name string) (*User, error) {
	return &User{"colin", "shenzhen"}
}

func (f *fake) List() ([]*User, error) {
	return []*User{&User{"colin", "shenzhen"}}, nil
}

func (f *fake) Delete(name string) error {
	return nil
}
