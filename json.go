package msgify

import "encoding/json"

type ToJson func(interface{}) ([]byte, error)

type FromJson func(string) (interface{}, error)

var toJsonDefault ToJson = func(msg interface{}) ([]byte, error) {
	return json.Marshal(msg)
}

// var toJsonDefault FromJson = func(b []byte) (, error) {
// 	return json.Marshal(msg)
// }
