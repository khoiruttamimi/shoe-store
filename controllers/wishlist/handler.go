package wishlist

import (
	"shoe-store/app/middleware"
	"shoe-store/controllers/wishlist/request"
	"shoe-store/controllers/wishlist/response"
	"shoe-store/domains/wishlist"

	controller "shoe-store/controllers"

	echo "github.com/labstack/echo/v4"
)

type WishlistController struct {
	wishlistUsecase wishlist.Usecase
}

func NewWishlistController(bu wishlist.Usecase) *WishlistController {
	return &WishlistController{
		wishlistUsecase: bu,
	}
}

func (ctrl *WishlistController) AddWishlist(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	req := request.Wishlist{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, err)
	}

	reqDomain := req.ToDomain()
	reqDomain.UserID = userID
	resp, err := ctrl.wishlistUsecase.Store(ctx, reqDomain)
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *WishlistController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.wishlistUsecase.GetAll(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	responseController := []response.Wishlist{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
