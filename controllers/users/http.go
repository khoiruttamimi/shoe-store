package users

import (
	"net/http"
	"shoe-store/app/middleware"
	"shoe-store/businesses/users"
	controller "shoe-store/controllers"
	"shoe-store/controllers/users/request"
	"shoe-store/controllers/users/response"

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

func (ctrl *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	token, err := ctrl.userUseCase.CreateToken(ctx, req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	response := struct {
		Role  string `json:"role"`
		Token string `json:"token"`
	}{Role: req.Role, Token: token}

	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	response, err := ctrl.userUseCase.Login(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	user, err := ctrl.userUseCase.GetProfile(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	response := response.FromDomain(user)
	return controller.NewSuccessResponse(c, response)
}
