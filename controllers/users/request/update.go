package requests

import "github.com/daffaalex22/seleksi-deall/business/users"

type UserUpdate struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (user *UserUpdate) ToDomain() users.Domain {
	return users.Domain{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}
