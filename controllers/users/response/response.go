package response

import (
	"time"

	"github.com/daffaalex22/seleksi-deall/business/users"
)

type UsersResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUsersResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		Id:        domain.Id,
		Email:     domain.Email,
		Name:      domain.Name,
		IsAdmin:   domain.IsAdmin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainList(domain []users.Domain) []UsersResponse {
	list := []UsersResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}

func FromDomainLogin(domain users.Domain) LoginUsersResponse {
	return LoginUsersResponse{
		Id:        domain.Id,
		Email:     domain.Email,
		Token:     domain.Token,
		Name:      domain.Name,
		IsAdmin:   domain.IsAdmin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainListLogin(domain []users.Domain) []LoginUsersResponse {
	list := []LoginUsersResponse{}
	for _, value := range domain {
		list = append(list, FromDomainLogin(value))
	}
	return list
}
