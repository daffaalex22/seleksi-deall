package requests

import "github.com/daffaalex22/seleksi-deall/business/users"

type UserAdd struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (user *UserAdd) ToDomain() users.Domain {
	return users.Domain{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}
