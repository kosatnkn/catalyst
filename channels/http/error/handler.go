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

		log.Error(ctx, "Server Error", err)
		return formatGenericError(err), http.StatusInternalServerError

	case *httpErrs.MiddlewareError,
		*domainErrs.DomainError,
		*externalErrs.RepositoryError,
		*externalErrs.ServiceError:

		log.Error(ctx, "Other Error", err)
		return formatGenericError(err), http.StatusBadRequest

	case *httpErrs.ValidationError:

		log.Error(ctx, "Validation Error", err)
		return formatValidationError(err), http.StatusUnprocessableEntity

	default:

		log.Error(ctx, "Unknown Error", err)
		return formatUnknownError(err), http.StatusInternalServerError
	}
}

// HandleValidatorErrors specifically handles validation errors thrown by the validator.
func HandleValidatorErrors(ctx context.Context, errs map[string]string, log adapters.LogAdapterInterface) (interface{}, int) {

	e := formatValidatorErrors(errs)

	log.Error(ctx, "Validation Errors", errs)

	return e, http.StatusUnprocessableEntity
}
