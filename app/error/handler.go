package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/app/error/types"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	domainError "github.com/kosatnkn/catalyst/domain/error"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, logger adapters.LogAdapterInterface) ([]byte, int) {

	var errMessage []byte
	var status int

	switch err.(type) {

	case *types.ServerError:
		logger.Error(ctx, "Server Error", err)
		errMessage = format(err)
		status = http.StatusInternalServerError
		break

	case *types.AdapterError,
		*types.MiddlewareError,
		*types.RepositoryError,
		*types.ServiceError,
		*domainError.DomainError:
		logger.Error(ctx, "Other Error", err)
		errMessage = format(err)
		status = http.StatusBadRequest
		break

	case *types.ValidationError:
		logger.Error(ctx, "Unpacker Error", err)
		errMessage = format(err)
		status = http.StatusUnprocessableEntity
		break

	default:
		logger.Error(ctx, "Unknown Error", err)
		errMessage = format(err)
		status = http.StatusInternalServerError
		break
	}

	return errMessage, status
}

// HandleValidationErrors specifically handles validation errors thrown by the validator.
func HandleValidationErrors(ctx context.Context, errs map[string]string, logger adapters.LogAdapterInterface) ([]byte, int) {

	errMessage := formatValidationErrors(errs)

	logger.Error(ctx, "Validation Errors", string(errMessage))

	return errMessage, http.StatusUnprocessableEntity
}
