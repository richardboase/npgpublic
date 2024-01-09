package assetlayer

import (
	"fmt"
	"reflect"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Body       interface{} `json:"body"`
	Error      string      `json:"error"`
	Errors     string      `json:"errors"`
}

func assertMapStringInterface(x interface{}) (map[string]interface{}, error) {
	m, ok := x.(map[string]interface{})
	if !ok {
		t := "nil"
		if x != nil {
			t = reflect.TypeOf(x).String()
		}
		return nil, fmt.Errorf("assertMapStringInterface: failed to assert type: %s", t)
	}
	return m, nil
}

func assertInterfaceArray(x interface{}) ([]interface{}, error) {
	m, ok := x.([]interface{})
	if !ok {
		t := "nil"
		if x != nil {
			t = reflect.TypeOf(x).String()
		}
		return nil, fmt.Errorf("assertInterfaceArray: failed to assert type: %s", t)
	}
	return m, nil
}
