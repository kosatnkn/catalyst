package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/app/transport/response"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	domainError "github.com/kosatnkn/catalyst/domain/error"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, w http.ResponseWriter, logger adapters.LogAdapterInterface) {

	switch err.(type) {
	case *types.ServerError:
		logger.Error(ctx, "Server Error", err)
		response.Send(w, format(err), http.StatusInternalServerError)
		break
	case *types.AdapterError, *types.MiddlewareError, *types.RepositoryError, *types.ServiceError, *domainError.DomainError:
		logger.Error(ctx, "Oter Error", err)
		response.Send(w, format(err), http.StatusBadRequest)
		break
	case *types.ValidationError:
		logger.Error(ctx, "Validation Structure Error", err)
		response.Send(w, format(err), http.StatusUnprocessableEntity)
		break
	default:
		logger.Error(ctx, "Unknown Error", err)
		response.Send(w, format(err), http.StatusInternalServerError)
		break
	}

	return
}

// HandleValidationErrors specifically handles validation errors thrown by the validator.
func HandleValidationErrors(ctx context.Context, errs map[string]string, w http.ResponseWriter, logger adapters.LogAdapterInterface) {

	errMessage := formatValidationErrors(errs)

	logger.Error(ctx, "Validation Errors", string(errMessage))
	response.Send(w, errMessage, http.StatusUnprocessableEntity)

	return
}
