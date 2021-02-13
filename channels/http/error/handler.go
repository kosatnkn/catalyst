package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/app/adapters"
	"github.com/kosatnkn/catalyst/channels/http/response/mappers"

	baseErrs "github.com/kosatnkn/catalyst/app/errors"
	httpErrs "github.com/kosatnkn/catalyst/channels/http/errors"
	domainErrs "github.com/kosatnkn/catalyst/domain/errors"
	externalErrs "github.com/kosatnkn/catalyst/externals/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, log adapters.LogAdapterInterface) (interface{}, int) {

	var e mappers.Error
	var status int

	switch err.(type) {

	case *baseErrs.ServerError,
		*httpErrs.TransformerError:

		e = format(err)
		status = http.StatusInternalServerError

		log.Error(ctx, "Server Error", err)

		break

	case *httpErrs.MiddlewareError,
		*externalErrs.RepositoryError,
		*externalErrs.ServiceError,
		*domainErrs.DomainError:

		e = format(err)
		status = http.StatusBadRequest

		log.Error(ctx, "Other Error", err)

		break

	case *httpErrs.ValidationError:

		e = format(err)
		status = http.StatusUnprocessableEntity

		log.Error(ctx, "Unpacker Error", err)

		break

	default:

		e = format(err)
		status = http.StatusInternalServerError

		log.Error(ctx, "Unknown Error", err)

		break
	}

	return e, status
}

// HandleValidationErrors specifically handles validation errors thrown by the validator.
func HandleValidationErrors(ctx context.Context, errs map[string]string, log adapters.LogAdapterInterface) (interface{}, int) {

	e := formatValidationErrors(errs)

	log.Error(ctx, "Validation Errors", errs)

	return e, http.StatusUnprocessableEntity
}
