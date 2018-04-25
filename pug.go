package pug

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

var config map[string]interface{}

// Parse parse a file
func Parse(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile(`{{_ \.(\w+)}}`)
	allSub := reg.FindAllSubmatch(b, -1)
	for _, sub := range allSub {
		rawValue := sub[0]
		env := string(sub[1])
		envValue := []byte(os.Getenv(env))
		b = bytes.Replace(b, rawValue, envValue, 1)
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
}

// GetString get string value
func GetString(s string) string {
	v := config[s].(string)
	return v
}

// GetInt get int value
func GetInt(s string) int {
	v := (config[s].(float64))
	return int(v)
}

// GetStringSlice get stringslice value
func GetStringSlice(s string) []string {
	var ss []string
	vs := config[s].([]interface{})
	for _, v := range vs {
		ss = append(ss, v.(string))
	}
	return ss
}
