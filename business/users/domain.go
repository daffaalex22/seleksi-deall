package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        string
	Name      string
	Email     string
	Token     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersUseCaseInterface interface {
	UsersLogin(ctx context.Context, domain Domain) (Domain, error)
	UsersGetByID(ctx context.Context, id string) (Domain, error)
	UsersGetAll(ctx context.Context) ([]Domain, error)
	UsersAdd(ctx context.Context, domain Domain) (Domain, error)
	UsersUpdate(ctx context.Context, domain Domain) (Domain, error)
	UsersDelete(ctx context.Context, id string) error
}

type UsersRepoInterface interface {
	UsersLogin(ctx context.Context, domain Domain) (Domain, error)
	UsersGetByID(ctx context.Context, id string) (Domain, error)
	UsersGetAll(ctx context.Context) ([]Domain, error)
	UsersAdd(ctx context.Context, domain Domain) (Domain, error)
	UsersUpdate(ctx context.Context, domain Domain) (Domain, error)
	UsersDelete(ctx context.Context, id string) error
}
