package serialization

import (
	"encoding/json"
	"log"
)

func Deserialize(data string) *AppRequest {
	req := new(AppRequest)
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		req.Error = err.Error()
		log.Print("Deserialization: ", err)
		return req
	}
	return req
}
