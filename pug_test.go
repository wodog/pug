package pug_test

import (
	"testing"

	"github.com/wodog/pug"
)

var config = `{
	"PORT": 8080,
	"MONGO": "{{_ .MONGO localhost:27017/dc-server}}",
	"TRACKER": "{{_ .TRACKER}}"
	}`
var configResult = `{
	"PORT": 8080,
	"MONGO": "localhost:27017/dc-server",
	"TRACKER": ""
	}`

func init() {
	pug.ParseString(config)
}

func TestGetInt(t *testing.T) {
	if pug.GetInt("PORT") != 8080 {
		t.Fail()
	}
}

func TestGetNil(t *testing.T) {
	if pug.GetString("TRACKER") != "" {
		t.Fail()
	}
}

func TestGetDefault(t *testing.T) {
	if pug.GetString("MONGO") != "localhost:27017/dc-server" {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	if pug.String() != configResult {
		t.Fail()
	}
}
