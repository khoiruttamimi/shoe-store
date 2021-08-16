package transaction

import (
	"net/http"
	"shoe-store/controllers/transaction/request"
	"shoe-store/controllers/transaction/response"
	"shoe-store/domains/transaction"

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

	req := request.Transaction{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	resp, err := ctrl.transactionUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TransactionController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.transactionUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Transaction{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}
