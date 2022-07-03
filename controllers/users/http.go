package users

import (
	"github.com/daffaalex22/seleksi-deall/app/middlewares"
	"github.com/daffaalex22/seleksi-deall/business/users"
	"github.com/daffaalex22/seleksi-deall/controllers"
	requests "github.com/daffaalex22/seleksi-deall/controllers/users/request"
	"github.com/daffaalex22/seleksi-deall/controllers/users/response"
	"github.com/daffaalex22/seleksi-deall/helper/err"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	usecase users.UsersUseCaseInterface
}

func NewUsersController(userUsecase users.UsersUseCaseInterface) *UsersController {
	return &UsersController{
		usecase: userUsecase,
	}
}

func (controller *UsersController) UsersLogin(c echo.Context) error {
	_, res := middlewares.ValidateAuthorization(c, []int{})
	if res != nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, result := controller.usecase.UsersLogin(ctx, userLogin.ToDomain())

	if result != nil {
		codeErr := err.ErrorGetUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}

	return controllers.SuccessResponse(c, response.FromDomainLogin(user))
}

func (controller *UsersController) UsersGetAll(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleAdmin})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	ctx := c.Request().Context()

	data, result := controller.usecase.UsersGetAll(ctx)
	if result != nil {
		codeErr := err.ErrorGetUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}

	return controllers.SuccessResponse(c, response.FromDomainList(data))
}

func (controller *UsersController) UsersGetByID(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleAdmin})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	id := c.Param("id")

	ctx := c.Request().Context()
	data, result := controller.usecase.UsersGetByID(ctx, id)
	if result != nil {
		codeErr := err.ErrDeleteUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}

	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *UsersController) UsersAdd(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleAdmin})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	req := requests.UserAdd{}
	c.Bind(&req)
	ctx := c.Request().Context()

	_, result := controller.usecase.UsersAdd(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorAddUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessOkResponse(c)
}

func (controller *UsersController) UsersUpdate(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleAdmin})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	req := requests.UserUpdate{}
	c.Bind(&req)

	ctx := c.Request().Context()
	_, result := controller.usecase.UsersUpdate(ctx, req.ToDomain())

	if result != nil {
		codeErr := err.ErrorUpdateUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}

	return controllers.SuccessOkResponse(c)
}

func (controller *UsersController) UsersDelete(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleAdmin})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	id := c.Param("id")

	ctx := c.Request().Context()
	result := controller.usecase.UsersDelete(ctx, id)
	if result != nil {
		codeErr := err.ErrDeleteUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}
	return controllers.SuccessOkResponse(c)
}

func (controller *UsersController) UsersGetMyData(c echo.Context) error {
	token, res := middlewares.ValidateAuthorization(c, []int{middlewares.RoleUser})
	if res != nil || token == nil {
		return controllers.ErrorResponse(c, err.ErrUnathorizedCheck(res), "error request", res)
	}

	id := token.UserID

	ctx := c.Request().Context()
	data, result := controller.usecase.UsersGetByID(ctx, id)
	if result != nil {
		codeErr := err.ErrDeleteUsersCheck(result)
		return controllers.ErrorResponse(c, codeErr, "error request", result)
	}

	return controllers.SuccessResponse(c, response.FromDomain(data))
}
