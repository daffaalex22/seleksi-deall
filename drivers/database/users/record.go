package users

import (
	"time"

	"github.com/daffaalex22/seleksi-deall/business/users"
	"gorm.io/gorm"
)

type Users struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Token     string
	IsAdmin   bool
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (user Users) ToDomain() users.Domain {
	return users.Domain{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Token:     user.Token,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Email:     domain.Email,
		Name:      domain.Name,
		Token:     domain.Token,
		Password:  domain.Password,
		IsAdmin:   domain.IsAdmin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToDomainList(datamds []Users) []users.Domain {
	All := []users.Domain{}
	for _, v := range datamds {
		All = append(All, v.ToDomain())
	}
	return All
}
