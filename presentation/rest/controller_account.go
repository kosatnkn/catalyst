package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kosatnkn/catalyst/v3/domain/entities"
	"github.com/kosatnkn/catalyst/v3/domain/usecases"
	"github.com/kosatnkn/catalyst/v3/infra"
)

// NOTE: This is the sample controller that we have created in order to demonstrate the implementation.
// A controller is used to group handlers that needs access to additional resources like the same set of persistence.
// It can also be used to logically group handlers together.

// accountController is the collection of handlers used to manipulate accounts.
type accountController struct {
	accounts *usecases.AccountUseCases
}

func newAccountController(ctr *infra.Container) *accountController {
	return &accountController{
		accounts: usecases.NewAccountUseCases(ctr),
	}
}

func (c *accountController) Get(ctx *gin.Context) {
	// get data from request
	paging, err := paging(ctx)
	if err != nil {
		// set error to Gin context so that the logging middleware can access it
		bindErrorToCtx(ctx, err, http.StatusUnprocessableEntity)
		return
	}

	filters, err := filters(ctx.Query("filters"))
	if err != nil {
		bindErrorToCtx(ctx, err, http.StatusUnprocessableEntity)
		return
	}

	// pass on to the usecase
	data, err := c.accounts.GetAccounts(context.Background(), filters, paging)
	if err != nil {
		bindErrorToCtx(ctx, err, http.StatusBadRequest)
		return
	}

	// respond
	ctx.JSON(http.StatusOK, responseData(data, nil))
}

func (c *accountController) Create(ctx *gin.Context) {
	// get data from request and validate (gin does this under the hood using 'go-playground/validator/v10')
	var reqBody accountRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		bindErrorToCtx(ctx, err, http.StatusUnprocessableEntity)
		return
	}

	// assign data to account struct
	a := entities.Account{
		Owner: reqBody.Owner,
	}

	// pass on to the usecase
	data, err := c.accounts.CreateAccount(context.Background(), a)
	if err != nil {
		bindErrorToCtx(ctx, err, http.StatusBadRequest)
		return
	}

	// respond
	ctx.JSON(http.StatusOK, responseData(data, nil))
}
