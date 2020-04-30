package registration

type User  interface {
	ID() string
	Login() string
}

type UsersRepository interface {
	HasLogin(login string) (bool, error)
	CreateUser(login string) User
}

type RestV1User struct {
	Users UsersRepository
}

// TODO New* is known to be a DI target, so pick up the interface spec and generate an adapter somewhere  which allows the wanted subtype-polymorphism
func NewRestV1User(users UsersRepository)*RestV1User {
	return nil
}

// GetUserById maps to GET /api/v1/user/{id}
func (r *RestV1User) GetById(id string) error {
	return nil
}

// PostRegisterWithLogin maps to POST /api/v1/user/{login}
func (r *RestV1User) PostRegisterWithLogin(login string) (User, error) {
	panic("")
}
