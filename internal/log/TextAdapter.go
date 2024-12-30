package log

import (
	"context"
	"fmt"
	"time"

	"github.com/gookit/color"
	"github.com/kosatnkn/catalyst/v3/app/adapters"
)

// TextAdapter is used to provide structured log messages.
type TextAdapter struct {
	cfg Config
}

// newTextAdapter creates a new Log adapter instance.
func newTextAdapter(cfg Config) (adapters.LogAdapterInterface, error) {
	a := &TextAdapter{
		cfg: cfg,
	}

	return a, nil
}

// AddTraceID attaches a trace id to context that can be later read by the logger.
func (a *TextAdapter) AddTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ID, id)
}

// AppendTracePoint appends the given trace point to a trace path in context that can be later read by the logger.
func (a *TextAdapter) AppendTracePoint(ctx context.Context, point string) context.Context {
	path := ctx.Value(TraceKey)
	if path == nil {
		return context.WithValue(ctx, TraceKey, point)
	}

	return context.WithValue(ctx, TraceKey, path.(string)+">"+point)
}

// Error logs a message as of error type.
func (a *TextAdapter) Error(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, levelError, message, options...)
}

// Debug logs a message as of debug type.
func (a *TextAdapter) Debug(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, levelDebug, message, options...)
}

// Info logs a message as of information type.
func (a *TextAdapter) Info(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, levelInfo, message, options...)
}

// Warn logs a message as of warning type.
func (a *TextAdapter) Warn(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, levelWarn, message, options...)
}

// Destruct will close the logger gracefully releasing all resources.
func (a *TextAdapter) Destruct() {
	// NOTE: nothing to implement
}

// log logs a message using the following format.
// <date> <time_in_24h_foramt_plus_milliseconds> [<log_level>] [<uuid>] [<trace_points>] [<message>] : [<additional_information>]
// ex:
//
//	2019/01/14 12:13:29.435517 [ERROR] [b2e1bfc7-11ed-40e5-ab08-abeadef079e6] [usecases.TestUsecase.TestFunc] [error message] : [key1: value1, ...]
func (a *TextAdapter) log(ctx context.Context, logLevel string, message string, options ...interface{}) {
	// check whether the message should be logged
	if !a.isLoggable(logLevel) {
		return
	}

	a.toConsole(a.formatMessage(ctx, logLevel, message, options...))
}

// formatMessage formats the log message.
func (a *TextAdapter) formatMessage(ctx context.Context, logLevel string, message string, options ...interface{}) string {
	var now = time.Now().Format("2006/01/02 15:04:05.000000")
	var level = a.setTag(logLevel)
	var uuid = "NONE"
	var trace = "NONE"

	id, ok := ctx.Value(ID).(string)
	if ok {
		uuid = id
	}

	points, ok := ctx.Value(TraceKey).(string)
	if ok {
		trace = points
	}

	if len(options) == 0 {
		return fmt.Sprintf("%s %s [%s] [%s] [%s]", now, level, uuid, trace, message)
	}

	return fmt.Sprintf("%s %s [%s] [%s] [%s] : %v", now, level, uuid, trace, message, options)
}

// isLoggable checks whether the message should be logged depending on the granularity of the log level.
func (a *TextAdapter) isLoggable(level string) bool {
	return granularity[level] >= granularity[a.cfg.Level]
}

// setTag generates tags based on color configuration settings.
func (a *TextAdapter) setTag(level string) interface{} {
	if a.cfg.Colors {
		switch level {
		case levelError:
			return color.New(color.FgRed).Sprint("[ERROR]")
		case levelDebug:
			return color.Debug.Sprint("[DEBUG]")
		case levelInfo:
			return color.Info.Sprint("[INFO]")
		case levelWarn:
			return color.New(color.FgYellow).Sprint("[WARN]")
		default:
			return "[" + level + "]"
		}
	}

	return "[" + level + "]"
}

// toConsole logs a message to the console.
func (a *TextAdapter) toConsole(message string) {
	fmt.Println(message)
}
