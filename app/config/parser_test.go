package config

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	// get parsed config
	c := Parse("../../config")

	// check validity
	t.Log(reflect.TypeOf(c))
}
