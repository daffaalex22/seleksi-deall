package users

import (
	"context"

	"github.com/daffaalex22/seleksi-deall/helper/encrypt"

	"github.com/daffaalex22/seleksi-deall/business/users"
	"github.com/daffaalex22/seleksi-deall/helper/err"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(gormDb *gorm.DB) users.UsersRepoInterface {
	return &UsersRepository{
		db: gormDb,
	}
}

func (repo *UsersRepository) UsersLogin(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	result := repo.db.First(&user, "email = ?", domain.Email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	err := encrypt.CheckPassword(domain.Password, user.Password)

	if err != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (repo *UsersRepository) UsersGetAll(ctx context.Context) ([]users.Domain, error) {
	var newUsers []Users
	resultAdd := repo.db.Find(&newUsers)
	if resultAdd.Error != nil {
		return []users.Domain{}, resultAdd.Error
	}
	return ToDomainList(newUsers), nil
}

func (repo *UsersRepository) UsersGetByID(ctx context.Context, id string) (users.Domain, error) {
	var newUsers Users
	resultAdd := repo.db.Where("id = ?", id).Find(&newUsers)
	if resultAdd.Error != nil {
		return users.Domain{}, resultAdd.Error
	}
	return newUsers.ToDomain(), nil
}

func (repo *UsersRepository) UsersAdd(ctx context.Context, domain users.Domain) (users.Domain, error) {
	newUsers := FromDomain(domain)

	hashedPassword, err := encrypt.Hash(domain.Password)
	if err != nil {
		return users.Domain{}, err
	}

	newUsers.Password = hashedPassword

	resultAdd := repo.db.Create(&newUsers)
	if resultAdd.Error != nil {
		return users.Domain{}, resultAdd.Error
	}
	return newUsers.ToDomain(), nil
}

func (repo *UsersRepository) UsersUpdate(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users
	updateUser := FromDomain(domain)

	resultUpdate := repo.db.Model(&user).Where("id = ?", updateUser.ID).Updates(updateUser)
	if resultUpdate.Error != nil {
		return users.Domain{}, resultUpdate.Error
	}
	return updateUser.ToDomain(), nil
}

func (repo *UsersRepository) UsersDelete(ctx context.Context, id string) error {
	var targetTable Users
	result := repo.db.Where("id = ?", id).Delete(&targetTable)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}
