package pug

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

type pug struct {
	gjson.Result
	b []byte
}

var p *pug

// Parse parse from reader
func Parse(r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	b = Format(b)
	p = &pug{
		gjson.ParseBytes(b),
		b,
	}
}

// ParseString parse from string
func ParseString(s string) {
	Parse(strings.NewReader(s))
}

// ParseFile parse from file
func ParseFile(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	Parse(bytes.NewReader(b))
}

// Format format b
func Format(b []byte) []byte {
	reg := regexp.MustCompile(`{{_ \.(\w+)( [a-zA-Z0-9\:/\-_.]+)?}}`)
	allSub := reg.FindAllSubmatch(b, -1)
	for _, sub := range allSub {
		rawValue := sub[0]
		env := string(sub[1])
		envValue := []byte(os.Getenv(env))
		envValueDefault := bytes.Trim(sub[2], " ")
		if len(envValue) == 0 {
			envValue = envValueDefault
		}
		b = bytes.Replace(b, rawValue, envValue, 1)
	}
	return b
}

// GetString get string value
func GetString(s string) string {
	return p.Get(s).String()
}

// GetInt get int value
func GetInt(s string) int64 {
	return p.Get(s).Int()
}

// String return raw data
func String() string {
	return string(p.b)
}
