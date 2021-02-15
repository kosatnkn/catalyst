package error

import (
	"context"
	"errors"

	"github.com/kosatnkn/catalyst/app/adapters"
)

// logError logs the error with trace.
func logError(ctx context.Context, log adapters.LogAdapterInterface, err error) {

	trace := []string{err.Error()}

	// limit unwraping depth
	for i := 0; i < 5; i++ {

		err = errors.Unwrap(err)
		if err == nil {
			break
		}

		trace = append(trace, err.Error())
	}

	if len(trace) == 1 {
		log.Error(ctx, formatForLog(trace[0]))
		return
	}

	log.Error(ctx, formatForLog(trace[0]), formatLogTrace(trace[1:]))
}
