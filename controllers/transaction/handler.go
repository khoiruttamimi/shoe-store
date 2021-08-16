package transaction

import (
	"errors"
	"net/http"
	"shoe-store/app/middleware"
	"shoe-store/controllers/transaction/request"
	"shoe-store/controllers/transaction/response"
	"shoe-store/domains/transaction"
	"strconv"

	controller "shoe-store/controllers"

	echo "github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionUsecase transaction.Usecase
}

func NewTransactionController(bu transaction.Usecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: bu,
	}
}

func (ctrl *TransactionController) Shopping(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	reqDomain := req.ToDomain()
	reqDomain.UserID = userID
	resp, err := ctrl.transactionUsecase.Store(ctx, reqDomain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TransactionController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.transactionUsecase.GetAll(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Transaction{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *TransactionController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("invalid parameter id"))
	}
	resp, err := ctrl.transactionUsecase.GetByID(ctx, id, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
