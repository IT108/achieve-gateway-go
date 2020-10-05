package processor

import (
	. "../serialization"
	"encoding/json"
)
import broker "gopkg.in/IT108/achieve-broker-go.v0"

func Process(req *AppRequest) {
	if req == nil {
		return
	}
	data, _ := json.Marshal(req.Data)
	broker.WriteMsg(req.Service, req.Method, string(data))

	print(req.Service)
}
