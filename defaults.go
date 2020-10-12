package msgify

import "encoding/json"

var (
	DefaultPrefetchCount = 1
	DefaultPrefetchSize  = 0
)

type ToJson func(interface{}) ([]byte, error)

type FromJson func(string) (interface{}, error)

var toJsonDefault ToJson = func(msg interface{}) ([]byte, error) {
	return json.Marshal(msg)
}

// var toJsonDefault FromJson = func(b []byte) (, error) {
// 	return json.Marshal(msg)
// }
