package product

import (
	controller "shoe-store/controllers"
	"shoe-store/controllers/product/request"
	"shoe-store/controllers/product/response"
	"shoe-store/domains"
	"shoe-store/domains/product"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase product.Usecase
}

func NewProductController(productUC product.Usecase) *ProductController {
	return &ProductController{
		productUseCase: productUC,
	}
}

func (ctrl *ProductController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Product{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, err)
	}

	resp, err := ctrl.productUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ProductController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, domains.ErrMissingID)
	}

	req := request.Product{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.productUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *ProductController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}
	resp, total, err := ctrl.productUseCase.GetAll(ctx, page, limit)
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	responseController := []response.Product{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	pagination := controller.PaginationRes(page, total, limit)
	return controller.NewSuccessResponseWithPagination(c, responseController, pagination)
}

func (ctrl *ProductController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controller.NewErrorResponse(c, domains.ErrParamID)
	}
	resp, err := ctrl.productUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
