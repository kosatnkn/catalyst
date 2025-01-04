package error

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kosatnkn/catalyst/v3/app/adapters"

	middlewareErrs "github.com/kosatnkn/catalyst/v3/app/transport/http/middleware/errors"
	unpackerErrs "github.com/kosatnkn/catalyst/v3/app/transport/http/request/unpackers/errors"
	domainErrs "github.com/kosatnkn/catalyst/v3/domain/errors"
	repositoryErrs "github.com/kosatnkn/catalyst/v3/externals/repositories/errors"
	transformerErrs "github.com/kosatnkn/catalyst/v3/transport/http/response/transformers/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, log adapters.LogAdapterInterface) (interface{}, int) {
	switch err.(type) {
	case *transformerErrs.TransformerError:
		logError(ctx, log, err)
		return formatGenericError(err), http.StatusInternalServerError
	case *middlewareErrs.MiddlewareError,
		*domainErrs.DomainError,
		*repositoryErrs.RepositoryError:
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
	log.Error(ctx, fmt.Sprintf("Validation Errors: %v", errs))
	return formatValidatorErrors(errs), http.StatusUnprocessableEntity
}
