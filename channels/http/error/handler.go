package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/catalyst/v2/app/adapters"

	middlewareErrs "github.com/kosatnkn/catalyst/v2/channels/http/middleware/errors"
	unpackerErrs "github.com/kosatnkn/catalyst/v2/channels/http/request/unpackers/errors"
	transformerErrs "github.com/kosatnkn/catalyst/v2/channels/http/response/transformers/errors"
	domainErrs "github.com/kosatnkn/catalyst/v2/domain/errors"
	repositoryErrs "github.com/kosatnkn/catalyst/v2/externals/repositories/errors"
	serviceErrs "github.com/kosatnkn/catalyst/v2/externals/services/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, log adapters.LogAdapterInterface) (interface{}, int) {

	switch err.(type) {
	case *transformerErrs.TransformerError:
		logError(ctx, log, err)
		return formatGenericError(err), http.StatusInternalServerError
	case *middlewareErrs.MiddlewareError,
		*domainErrs.DomainError,
		*repositoryErrs.RepositoryError,
		*serviceErrs.ServiceError:
		logError(ctx, log, err)
		return formatGenericError(err), http.StatusBadRequest
	case *unpackerErrs.UnpackerError:
		logError(ctx, log, err)
		return formatUnpackerError(err), http.StatusUnprocessableEntity
	default:
		logError(ctx, log, err)
		return formatUnknownError(err), http.StatusInternalServerError
	}
}

// HandleValidatorErrors specifically handles validation errors thrown by the validator.
func HandleValidatorErrors(ctx context.Context, errs map[string]string, log adapters.LogAdapterInterface) (interface{}, int) {

	log.Error(ctx, "Validation Errors", errs)
	return formatValidatorErrors(errs), http.StatusUnprocessableEntity
}
