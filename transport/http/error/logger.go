package error

import (
	"context"
	"errors"

	"github.com/kosatnkn/catalyst/v2/app/adapters"
)

// logError logs the error with trace.
func logError(ctx context.Context, log adapters.LogAdapterInterface, err error) {
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

	log.Error(ctx, formatForLog(trace[0]))

	// when error has an embedded error chain log the error chan as a trace
	// debug mode is used to log the error trace so that the trace will be printed only when
	// the application is running in the debug mode.
	if len(trace) > 1 {
		log.Debug(ctx, formatForLog(trace[0]), formatLogTrace(trace[1:]))
	}
}
