package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	// input
	dir := "../../configs/"

	// run
	c := Parse(dir)

	// check
	need := reflect.TypeOf(&Config{})
	got := reflect.TypeOf(c)
	if got != need {
		t.Errorf("Required %v, got %v", need, got)
	}
}

func TestParseInvalidDir(t *testing.T) {
	// check panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to panic but did not")
		}
	}()

	// input
	dir := "./invalid_config_dir/"

	// run
	_ = Parse(dir)
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

func TestParseConfigUnmatchedUnpacker(t *testing.T) {
	// check panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to panic but did not")
		}
	}()

	// input
	file := "../../configs/app.yaml"
	var unpacker string // some data type that cannot be unmarshalled in to

	// run
	parseConfig(file, &unpacker)

	t.Log(unpacker)
}
