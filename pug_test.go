package pug_test

import (
	"fmt"
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

	ddd := pug.GetString("default")
	fmt.Println(ddd)
	if ddd != "ddd" {
		t.Error("GetStringDefault")
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
