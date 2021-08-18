package users

import (
	"net/http"
	"shoe-store/app/middleware"
	controller "shoe-store/controllers"
	"shoe-store/controllers/users/request"
	"shoe-store/controllers/users/response"
	"shoe-store/domains/users"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.UserRegister{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqDomain := req.ToDomain()
	reqDomain.Role = "customer"
	user, token, err := ctrl.userUseCase.Register(ctx, reqDomain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := response.GetAuthResponse(user, token)
	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) RegisterAdmin(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.UserRegister{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqDomain := req.ToDomain()
	reqDomain.Role = "admin"
	user, token, err := ctrl.userUseCase.Register(ctx, reqDomain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := response.GetAuthResponse(user, token)
	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.UserLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, token, err := ctrl.userUseCase.Login(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := response.GetAuthResponse(user, token)
	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	user, err := ctrl.userUseCase.GetProfile(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	response := response.GetProfileResponse(user)
	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	req := request.UpdateProfile{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	reqDomain := req.ToDomain()
	reqDomain.ID = userID
	user, err := ctrl.userUseCase.UpdateProfile(ctx, reqDomain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	response := response.GetProfileResponse(user)
	return controller.NewSuccessResponse(c, response)
}
