package loggerjson

import (
	"context"
	"os"

	"github.com/kosatnkn/catalyst/v3/telemetry/log"
	"github.com/rs/zerolog"
)

var granularity map[string]zerolog.Level = map[string]zerolog.Level{
	log.LevelInfo:  zerolog.InfoLevel,
	log.LevelDebug: zerolog.DebugLevel,
	log.LevelWarn:  zerolog.WarnLevel,
	log.LevelError: zerolog.ErrorLevel,
}

// LoggerJSON is used to provide structured log messages.
type LoggerJSON struct {
	cfg Config
	l   zerolog.Logger
}

// NewLoggerJSON creates a new instance.
func NewLoggerJSON(cfg Config) (log.Logger, error) {
	if err := validateCfg(cfg); err != nil {
		return nil, err
	}

	a := &LoggerJSON{
		cfg: cfg,
		l:   zerolog.New(os.Stdout).With().Timestamp().Logger().Level(granularity[cfg.Level]),
	}

	return a, nil
}

// AddTraceID attaches a trace id to context that can be later read by the logger.
func (a *LoggerJSON) AddTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, log.ID, id)
}

// AppendTracePoint appends the given trace point to a trace path in context that can be later read by the logger.
func (a *LoggerJSON) AppendTracePoint(ctx context.Context, point string) context.Context {
	path := ctx.Value(log.TraceKey)
	if path == nil {
		return context.WithValue(ctx, log.TraceKey, point)
	}

	return context.WithValue(ctx, log.TraceKey, path.(string)+" - "+point)
}

// Error logs a message as of error type.
func (a *LoggerJSON) Error(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Error()).Msg(message)
}

// Debug logs a message as of debug type.
func (a *LoggerJSON) Debug(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Debug()).Msg(message)
}

// Info logs a message as of information type.
func (a *LoggerJSON) Info(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Info()).Msg(message)
}

// Warn logs a message as of warning type.
func (a *LoggerJSON) Warn(ctx context.Context, message string) {
	a.withCtxVals(ctx, a.l.Warn()).Msg(message)
}

// withCtxVals add special values set in the context by the logger to the log event.
func (a *LoggerJSON) withCtxVals(ctx context.Context, zlEvent *zerolog.Event) *zerolog.Event {
	if id, ok := ctx.Value(log.ID).(string); ok {
		zlEvent = zlEvent.Str("id", id)
	}

	if trace, ok := ctx.Value(log.TraceKey).(string); ok {
		zlEvent = zlEvent.Str("trace", trace)
	}

	return zlEvent
}
