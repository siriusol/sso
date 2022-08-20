package notation

import "encoding/json"

func JSONStr(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
