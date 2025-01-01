package config

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	// input
	file := "../../configs/app.yaml"

	// run
	c := read(file)

	// check
	need := "[]uint8"
	got := reflect.TypeOf(c).String()
	if got != need {
		t.Errorf("Required %v, got %v", need, got)
	}
}

func TestReadInvalidFile(t *testing.T) {
	// check panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to panic but did not")
		}
	}()

	// input
	file := "./invalid_config_dir/invalid_file.yaml"

	// run
	_ = read(file)
}
