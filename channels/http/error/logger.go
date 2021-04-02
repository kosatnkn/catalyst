package error

import (
	"context"
	"errors"

	"github.com/kosatnkn/catalyst/v2/app/adapters"
)

// logError logs the error with trace.
func logError(ctx context.Context, log adapters.LogAdapterInterface, err error) {

	log.Error(ctx, formatForLog(err.Error()))
	logTrace(ctx, log, err)
}

// logTrace logs the error race.
func logTrace(ctx context.Context, log adapters.LogAdapterInterface, err error) {

	trace := []string{
		err.Error(), // add the top most error of the error chain
	}

	// unwrap error in a loop to get previous errors in the chain
	for {
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

	log.Debug(ctx, formatForLog(trace[0]), formatLogTrace(trace[1:]))
}
