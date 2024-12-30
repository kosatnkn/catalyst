package log

import (
	"context"
	"os"

	"github.com/kosatnkn/catalyst/v3/app/adapters"
	"github.com/rs/zerolog"
)

// JSONLogger is used to provide structured log messages.
type JSONLogger struct {
	cfg Config
	l   zerolog.Logger
}

// newJSONLogger creates a new Log adapter instance.
func newJSONLogger(cfg Config) (adapters.LogAdapterInterface, error) {
	a := &JSONLogger{
		cfg: cfg,
		l:   zerolog.New(os.Stdout).With().Timestamp().Logger().Level(granularity[cfg.Level]),
	}

	return a, nil
}

// AddTraceID attaches a trace id to context that can be later read by the logger.
func (a *JSONLogger) AddTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ID, id)
}

// AppendTracePoint appends the given trace point to a trace path in context that can be later read by the logger.
func (a *JSONLogger) AppendTracePoint(ctx context.Context, point string) context.Context {
	path := ctx.Value(TraceKey)
	if path == nil {
		return context.WithValue(ctx, TraceKey, point)
	}

	return context.WithValue(ctx, TraceKey, path.(string)+" - "+point)
}

// Error logs a message as of error type.
func (a *JSONLogger) Error(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Error()).Msg(message)
}

// Debug logs a message as of debug type.
func (a *JSONLogger) Debug(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Debug()).Msg(message)
}

// Info logs a message as of information type.
func (a *JSONLogger) Info(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Info()).Msg(message)
}

// Warn logs a message as of warning type.
func (a *JSONLogger) Warn(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Warn()).Msg(message)
}

// withCtxVals add special values set in the context by the logger to the log event.
func (a *JSONLogger) withCtxVals(ctx context.Context, zlEvent *zerolog.Event) *zerolog.Event {
	if id, ok := ctx.Value(ID).(string); ok {
		zlEvent = zlEvent.Str("id", id)
	}

	if trace, ok := ctx.Value(TraceKey).(string); ok {
		zlEvent = zlEvent.Str("trace", trace)
	}

	return zlEvent
}
