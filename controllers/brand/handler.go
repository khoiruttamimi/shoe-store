package brand

import (
	"net/http"
	"shoe-store/controllers/brand/request"
	"shoe-store/controllers/brand/response"
	"shoe-store/domains/brand"

	controller "shoe-store/controllers"

	echo "github.com/labstack/echo/v4"
)

type BrandController struct {
	brandUsecase brand.Usecase
}

func NewBrandController(bu brand.Usecase) *BrandController {
	return &BrandController{
		brandUsecase: bu,
	}
}

func (ctrl *BrandController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Brand{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.brandUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *BrandController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.brandUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Brand{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
