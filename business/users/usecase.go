package users

import (
	"context"
	"fmt"
	"time"

	"github.com/daffaalex22/seleksi-deall/app/middlewares"
	"github.com/daffaalex22/seleksi-deall/helper/err"
	uuid "github.com/satori/go.uuid"
)

type UsersUseCase struct {
	ConfigJWT *middlewares.ConfigJWT
	repo      UsersRepoInterface
	ctx       time.Duration
}

func NewUsersUseCase(mdsRepo UsersRepoInterface, contextTimeout time.Duration, configJWT *middlewares.ConfigJWT) UsersUseCaseInterface {
	return &UsersUseCase{
		ConfigJWT: configJWT,
		repo:      mdsRepo,
		ctx:       contextTimeout,
	}
}

func (usecase *UsersUseCase) UsersLogin(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}

	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
	}

	user, err := usecase.repo.UsersLogin(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	user.Token, err = usecase.ConfigJWT.GenerateTokenJWT(user.ID, user.IsAdmin)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UsersUseCase) UsersGetAll(ctx context.Context) ([]Domain, error) {
	fmt.Println(ctx)
	result, err := usecase.repo.UsersGetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (usecase *UsersUseCase) UsersGetByID(ctx context.Context, id string) (Domain, error) {

	users, res := usecase.repo.UsersGetByID(ctx, id)
	if res != nil {
		return Domain{}, res
	}
	if (users.CreatedAt == time.Time{}) {
		return Domain{}, err.ErrNotFound
	}

	return users, nil
}

func (usecase *UsersUseCase) UsersAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.ID == "" {
		domain.ID = uuid.NewV4().String()
	}
	if domain.Name == "" {
		return Domain{}, err.ErrNameEmpty
	}
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
	}

	users, result := usecase.repo.UsersAdd(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	return users, nil
}

func (usecase *UsersUseCase) UsersUpdate(ctx context.Context, domain Domain) (Domain, error) {

	if domain.ID == "" {
		return Domain{}, err.ErrIDEmpty
	}
	if domain.Name == "" {
		return Domain{}, err.ErrNameEmpty
	}
	if domain.Email == "" {
		return Domain{}, err.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, err.ErrPasswordEmpty
	}

	users, result := usecase.repo.UsersUpdate(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	if (users.CreatedAt == time.Time{} && users.Email == "") {
		return Domain{}, err.ErrNotFound
	}
	return users, nil
}

func (usecase *UsersUseCase) UsersDelete(ctx context.Context, id string) error {
	result := usecase.repo.UsersDelete(ctx, id)
	if result != nil {
		return result
	}
	return nil
}
