package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	// input
	dir := "../../config/"

	// run
	c := Parse(dir)

	// check
	need := reflect.TypeOf(&Config{})
	got := reflect.TypeOf(c)

	if got != need {
		t.Errorf("Required %v, got %v", need, got)
	}
}

func TestGetConfigDirTailing(t *testing.T) {

	// input
	s := string(os.PathSeparator)
	pathTailing := fmt.Sprintf(".%sconfig%s", s, s)

	// run
	got := getConfigDir(pathTailing)

	// check
	need := pathTailing

	if need != got {
		t.Errorf("Required %v, got %v", need, got)
	}
}

func TestGetConfigDirNoTailing(t *testing.T) {

	// input
	s := string(os.PathSeparator)
	pathTailing := fmt.Sprintf(".%sconfig%s", s, s)
	pathNoTailing := fmt.Sprintf(".%sconfig", s)

	// run
	got := getConfigDir(pathNoTailing)

	// check
	need := pathTailing

	if need != got {
		t.Errorf("Required %v, got %v", need, got)
	}
}
