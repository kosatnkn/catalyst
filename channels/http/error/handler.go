package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/app/adapters"

	httpErrs "github.com/kosatnkn/catalyst/channels/http/errors"
	domainErrs "github.com/kosatnkn/catalyst/domain/errors"
	externalErrs "github.com/kosatnkn/catalyst/externals/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, log adapters.LogAdapterInterface) (interface{}, int) {

	switch err.(type) {

	case *httpErrs.TransformerError:

		logError(ctx, log, err)
		return formatGenericError(err), http.StatusInternalServerError

	case *httpErrs.MiddlewareError,
		*domainErrs.DomainError,
		*externalErrs.RepositoryError,
		*externalErrs.ServiceError:

		logError(ctx, log, err)
		return formatGenericError(err), http.StatusBadRequest

	case *httpErrs.ValidationError:

		logError(ctx, log, err)
		return formatValidationError(err), http.StatusUnprocessableEntity

	default:

		logError(ctx, log, err)
		return formatUnknownError(err), http.StatusInternalServerError
	}
}

// HandleValidatorErrors specifically handles validation errors thrown by the validator.
func HandleValidatorErrors(ctx context.Context, errs map[string]string, log adapters.LogAdapterInterface) (interface{}, int) {

	e := formatValidatorErrors(errs)

	log.Error(ctx, "Validation Errors", errs)

	return e, http.StatusUnprocessableEntity
}
