package processor

import . "../serialization"
import "github.com/it108/achieve-broker-go"

func Process(req *AppRequest) {
	// TODO: kafka request
	if req == nil {
		return
	}
	achieve_broker_go.WriteMsg(req.Service, req.Data)

	print(req.Service)
}
