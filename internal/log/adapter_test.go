package log_test

import (
	"context"
	"testing"

	"github.com/kosatnkn/catalyst/v3/app/adapters"
	"github.com/kosatnkn/catalyst/v3/internal/log"
)

func newLogger(t *testing.T, level string) adapters.LogAdapterInterface {

	cfg := log.Config{
		Level:  level,
		Colors: true,
	}

	l, err := log.NewAdapter(cfg)
	if err != nil {
		t.Fatalf("Error creating logger %v", err)
	}

	return l
}

func TestInvalidConfig(t *testing.T) {

	// invalid levels
	tcs := []struct {
		name  string
		level string
	}{
		{"Non existant", "ERR"},
		{"Different case", "Error"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			cfg := log.Config{
				Level:  tc.level,
				Colors: true,
			}

			_, err := log.NewAdapter(cfg)
			if err == nil {
				t.Errorf("Expected to return an error but did not")
			}
		})
	}

}

func TestMessage(t *testing.T) {

	l := newLogger(t, "INFO")

	l.Error(context.Background(), "Hello")
	l.Error(context.Background(), "Hello", "Additional 1", "Additional 2")

	l.Warn(context.Background(), "Hello")
	l.Warn(context.Background(), "Hello", "Additional 1", "Additional 2")

	l.Debug(context.Background(), "Hello")
	l.Debug(context.Background(), "Hello", "Additional 1", "Additional 2")

	l.Info(context.Background(), "Hello")
	l.Info(context.Background(), "Hello", "Additional 1", "Additional 2")
}

func TestLogLevels(t *testing.T) {

	tcs := []struct {
		name  string
		level string
	}{
		{"Error level logger", "ERROR"},
		{"Warn level logger", "WARN"},
		{"Debug level logger", "DEBUG"},
		{"Info level logger", "INFO"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			l := newLogger(t, tc.level)

			l.Error(context.Background(), "Hello")
			l.Error(context.Background(), "Hello", "Additional 1", "Additional 2")

			l.Warn(context.Background(), "Hello")
			l.Warn(context.Background(), "Hello", "Additional 1", "Additional 2")

			l.Debug(context.Background(), "Hello")
			l.Debug(context.Background(), "Hello", "Additional 1", "Additional 2")

			l.Info(context.Background(), "Hello")
			l.Info(context.Background(), "Hello", "Additional 1", "Additional 2")
		})
	}
}
