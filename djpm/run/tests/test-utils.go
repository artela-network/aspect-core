package tests

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

func GetTestTarget(target string) ([]byte, error) {
	readFile, _ := os.ReadFile("./test-config.json")
	get := jsoniter.Get(readFile, "params")
	param := make(map[string]interface{})
	get.ToVal(&param)

	path := jsoniter.Get(readFile, "target", target).ToString()
	expand := os.Expand(path, func(k string) string {
		if val, ok := param[k]; ok {
			return val.(string)
		}
		return ""
	})
	return os.ReadFile(expand)
}
