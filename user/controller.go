package user

import "github.com/worldiety/mercurius/ee"

type Controller struct {
	users Repository
}

func NewController(users Repository) *Controller {
	return &Controller{users}
}

func (c *Controller) ShowUsersREST(limit int, offset int) (*UsersResponse, error) {
	users, err := c.users.FindAll(100, 0)
	return &UsersResponse{users}, err
}

func (c *Controller) ShowUsersHTML(limit int, offset int) ee.Renderer {
	users, err := c.users.FindAll(100, 0)

}
