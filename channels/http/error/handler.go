package error

import (
	"context"
	"net/http"

	baseErrs "github.com/kosatnkn/catalyst/app/errors"
	httpErrs "github.com/kosatnkn/catalyst/channels/http/errors"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	domainErrs "github.com/kosatnkn/catalyst/domain/errors"
	externalErrs "github.com/kosatnkn/catalyst/externals/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, logger adapters.LogAdapterInterface) ([]byte, int) {

	var errMessage []byte
	var status int

	switch err.(type) {

	case *baseErrs.ServerError,
		*httpErrs.TransformerError:
		logger.Error(ctx, "Server Error", err)
		errMessage = format(err)
		status = http.StatusInternalServerError
		break

	case *externalErrs.AdapterError,
		*httpErrs.MiddlewareError,
		*externalErrs.RepositoryError,
		*externalErrs.ServiceError,
		*domainErrs.DomainError:
		logger.Error(ctx, "Other Error", err)
		errMessage = format(err)
		status = http.StatusBadRequest
		break

	case *httpErrs.ValidationError:
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
