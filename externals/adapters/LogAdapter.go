package adapters

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"github.com/kosatnkn/catalyst/domain/globals"
)

// LogAdapter is used to wrap
type LogAdapter struct {
	cfg config.LogConfig
	lf  *os.File
}

// New creates a new Postgres adapter instance.
func (a *LogAdapter) New(config config.LogConfig) (adapters.LogAdapterInterface, error) {

	a.cfg = config
	a.lf = a.initLogFile()

	return a, nil
}

// Error logs a message as of error type.
func (a *LogAdapter) Error(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, "ERROR", message, options)
}

// Debug logs a message as of debug type.
func (a *LogAdapter) Debug(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, "DEBUG", message)
}

// Info logs a message as of information type.
func (a *LogAdapter) Info(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, "INFO", message, options)
}

// Warn logs a message as of warning type.
func (a *LogAdapter) Warn(ctx context.Context, message string, options ...interface{}) {
	a.log(ctx, "WARN", message, options)
}

// Destruct will close the logger gracefully releasing all resources.
func (a *LogAdapter) Destruct() {
	if a.cfg.File {
		a.lf.Close()
	}
}

// Initialize the log file.
func (a *LogAdapter) initLogFile() *os.File {

	ld := a.cfg.Directory

	f, err := os.OpenFile(filepath.Join(ld, "out.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	return f
}

// Logs a message using the following format.
// <date> <time_in_24h_foramt_plus_milliseconds> [<message_type>] [<uuid>] [<prefix>] [<message>] [<additional_inforamtion>]
// ex:
//		2019/01/14 12:13:29.435517 [ERROR] [b2e1bfc7-11ed-40e5-ab08-abeadef079e6] [usecases.TestUsecase.TestFunc] [error message] [key1: value1, ...]
func (a *LogAdapter) log(ctx context.Context, logLevel string, message string, options ...interface{}) {

	// check whether the message should be logged
	if !a.isLoggable(logLevel) {
		return
	}

	nowVal := time.Now().Format("2006/01/02 15:04:05.000000")
	uuidVal := ctx.Value(globals.UUIDKey)
	prefixVal := ctx.Value(globals.PrefixKey)
	logLevelVal := a.setTag(logLevel)

	formattedMessage := fmt.Sprintf("%s %s [%s] [%v] [%v] [%v]", nowVal, logLevelVal, uuidVal, prefixVal, message, options)

	a.logToConsole(formattedMessage)
	a.logToFile(formattedMessage)
}

// Check whether the message should be logged depending on the log level setting.
func (a *LogAdapter) isLoggable(logLevel string) bool {

	l := map[string]int{
		"ERROR": 1,
		"DEBUG": 2,
		"WARN":  3,
		"INFO":  4,
	}

	return l[logLevel] >= l[a.cfg.Level]
}

// Generate tags based on color configuration settings.
func (a *LogAdapter) setTag(logLevel string) interface{} {

	if a.cfg.Colors {
		var logLevelVal aurora.Value

		switch logLevel {
		case "ERROR":
			logLevelVal = aurora.BgRed("[ERROR]")
			break
		case "DEBUG":
			logLevelVal = aurora.BgGreen("[DEBUG]")
			break
		case "INFO":
			logLevelVal = aurora.BgCyan("[INFO]")
			break
		case "WARN":
			logLevelVal = aurora.BgBrown("[WARN]")
			break
		default:
			logLevelVal = aurora.BgMagenta("[UNKNOWN]")
			break
		}

		return logLevelVal
	}

	return "[" + logLevel + "]"
}

// Logs a message to the console.
func (a *LogAdapter) logToConsole(message string) {

	if a.cfg.Console {
		fmt.Println(message)
	}
}

// Logs a message to a file.
func (a *LogAdapter) logToFile(message string) {

	if !a.cfg.File {
		return
	}

	_, err := a.lf.WriteString(message + "\n")

	if err != nil {
		fmt.Println(err)
	}
}
