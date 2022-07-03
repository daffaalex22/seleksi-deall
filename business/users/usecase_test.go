package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/daffaalex22/seleksi-deall/app/middlewares"
	"github.com/daffaalex22/seleksi-deall/business/users"
	_mockUserRepository "github.com/daffaalex22/seleksi-deall/business/users/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.UsersRepoInterface
var userService users.UsersUseCaseInterface
var userDomain users.Domain
var usersDomain []users.Domain

var configJWT middlewares.ConfigJWT

func setup() {
	configJWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	userService = users.NewUsersUseCase(&userRepository, time.Hour*1, &configJWT)
	userDomain = users.Domain{
		ID:        "my-id",
		Name:      "my-name",
		Email:     "my-email@mail.com",
		Password:  "my-secret-password",
		Token:     "my-json-web-token",
		IsAdmin:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	usersDomain = append(usersDomain, userDomain)
}

func TestLogin(t *testing.T) {

	setup()
	userRepository.On("UsersLogin",
		mock.Anything,
		mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()

	t.Run("Test case 1 | Valid Login", func(t *testing.T) {
		user, err := userService.UsersLogin(context.Background(), users.Domain{
			Email:    "my-email@mail.com",
			Password: "my-password",
		})

		assert.Nil(t, err)
		assert.Equal(t, "my-name", user.Name)
	})

	t.Run("Test case 2 | Error Login", func(t *testing.T) {
		userRepository.On("UsersLogin",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(users.Domain{}, errors.New("Unexpected Error")).Once()
		user, err := userService.UsersLogin(context.Background(), users.Domain{
			Email:    "my-email@mail.com",
			Password: "my-password",
		})

		assert.NotNil(t, err)
		assert.Equal(t, user, users.Domain{})
	})

	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {

		_, err := userService.UsersLogin(context.Background(), users.Domain{
			Email:    "",
			Password: "my-password",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Password Empty", func(t *testing.T) {

		_, err := userService.UsersLogin(context.Background(), users.Domain{
			Email:    "my-email@mail.com",
			Password: "",
		})
		assert.NotNil(t, err)
	})
}

func TestGetAllUser(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UsersGetAll",
			mock.Anything).Return(usersDomain, nil).Once()

		user, err := userService.UsersGetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("UsersGetAll",
			mock.Anything).Return([]users.Domain{}, errors.New("Unexpected Error")).Once()

		user, err := userService.UsersGetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, user, []users.Domain{})

		userRepository.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid", func(t *testing.T) {
		userRepository.On("UsersGetByID",
			mock.Anything,
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.UsersGetByID(context.Background(), userDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("UsersGetByID",
			mock.Anything,
			mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("Unexpected Error")).Once()

		user, err := userService.UsersGetByID(context.Background(), userDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, user, users.Domain{})
	})
}

func TestAddUser(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Add", func(t *testing.T) {
		userRepository.On("UsersAdd",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(userDomain, nil).Once()
		user, err := userService.UsersAdd(context.Background(), users.Domain{
			Name:     "my-name",
			Email:    "my-email@mail.com",
			Password: "my-password",
			IsAdmin:  true,
		})

		assert.Nil(t, err)
		assert.Equal(t, "my-name", user.Name)
	})

	t.Run("Test case 2 | Error Add", func(t *testing.T) {
		userRepository.On("UsersAdd",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(users.Domain{}, errors.New("Unexpected Error")).Once()
		user, err := userService.UsersAdd(context.Background(), users.Domain{})

		assert.Error(t, err)
		assert.Equal(t, user, users.Domain{})
	})

	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {
		_, err := userService.UsersAdd(context.Background(), users.Domain{
			Name:     "my-name",
			Email:    "",
			Password: "my-password",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Password Empty", func(t *testing.T) {
		_, err := userService.UsersAdd(context.Background(), users.Domain{
			Name:     "my-name",
			Email:    "my-email@mail.com",
			Password: "",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	setup()
	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UsersUpdate",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(users.Domain{
			ID:        userDomain.ID,
			Name:      "my-new-name",
			Email:     userDomain.Email,
			Password:  userDomain.Password,
			IsAdmin:   false,
			CreatedAt: userDomain.CreatedAt,
			UpdatedAt: userDomain.UpdatedAt,
		}, nil).Once()

		user, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "my-id",
			Name:     "my-new-name",
			Email:    "my-email@mail.com",
			Password: "my-password",
			IsAdmin:  false,
		})

		assert.NoError(t, err)
		assert.Equal(t, "my-new-name", user.Name)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("UsersUpdate",
			mock.Anything,
			mock.AnythingOfType("users.Domain")).Return(users.Domain{
			ID:       userDomain.ID,
			Name:     "my-new-name",
			Email:    userDomain.Email,
			Password: userDomain.Password,
			IsAdmin:  false,
		}, errors.New("Unexpected Error")).Once()

		user, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "my-id",
			Name:     "my-new-name",
			Email:    "my-email@mail.com",
			Password: "my-password",
			IsAdmin:  true,
		})

		assert.Error(t, err)
		assert.Equal(t, users.Domain{}, user)
	})

	t.Run("Test Case 4 | Invalid ID Empty", func(t *testing.T) {
		_, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "",
			Name:     "my-name",
			Email:    "my-email@mail.com",
			Password: "my-password",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Email Empty", func(t *testing.T) {
		_, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "my-id",
			Name:     "my-name",
			Email:    "",
			Password: "my-password",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Invalid Password Empty", func(t *testing.T) {
		_, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "my-id",
			Name:     "my-name",
			Email:    "my-email@mail.com",
			Password: "",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 6 | Invalid Name Empty", func(t *testing.T) {
		_, err := userService.UsersUpdate(context.Background(), users.Domain{
			ID:       "my-id",
			Name:     "",
			Email:    "my-email@mail.com",
			Password: "my-password",
			IsAdmin:  true,
		})
		assert.NotNil(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UsersDelete",
			mock.Anything,
			mock.AnythingOfType("string")).Return(nil).Once()

		err := userService.UsersDelete(context.Background(), userDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("UsersDelete",
			mock.Anything,
			mock.AnythingOfType("string")).Return(errors.New("Unexpected Error")).Once()

		err := userService.UsersDelete(context.Background(), userDomain.ID)

		assert.Error(t, err)
	})
}
