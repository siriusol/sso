package cookie

import "encoding/json"

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(bs []byte, v interface{}) error {
	return json.Unmarshal(bs, v)
}
