package pug_test

import (
	"os"
	"testing"

	"github.com/wodog/pug"
)

func Test(t *testing.T) {
	os.Setenv("HELLO", "world")
	pug.Parse("config.json")

	hello := pug.GetString("HELLO")
	if hello != "world" {
		t.Error("GetString")
	}

	num := pug.GetInt("num")
	if num != 123456 {
		t.Error("GetInt")
	}

	array := pug.GetStringSlice("array")
	if len(array) != 4 {
		t.Error("GetStringSlice")
	}
}
