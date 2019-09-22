package config

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	// get parsed config
	c := Parse("../../config")

	// check data structure type
	need := reflect.TypeOf(&Config{})
	got := reflect.TypeOf(c)

	if got != need {

		t.Errorf("Required %v, got %v", need, got)
	}
}
