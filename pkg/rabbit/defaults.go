package rabbit

import "encoding/json"

var (
	DefaultPrefetchCount = 1
	DefaultPrefetchSize  = 0
)

type ToJSON func(interface{}) ([]byte, error)

type FromJSON func(string) (interface{}, error)

var toJSONDefault ToJSON = func(msg interface{}) ([]byte, error) {
	return json.Marshal(msg)
}

// var toJsonDefault FromJson = func(b []byte) (, error) {
// 	return json.Marshal(msg)
// }
