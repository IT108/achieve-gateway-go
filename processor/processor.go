package processor

import (
	"encoding/json"
	broker "github.com/IT108/achieve-broker-go"
	. "github.com/IT108/achieve-gateway-go/serialization"
	. "github.com/IT108/achieve-gateway-go/utils"
	"log"
)

var GateId string

const GATE_ID_LEN = 10

func Process(req *AppRequest) {
	req.Data["GateId"] = GateId
	req.Data["Sender"] = req.ClientId
	req.Data["Method"] = req.Method
	req.Data["RequestId"] = req.RequestId
	if req == nil {
		return
	}
	data, _ := json.Marshal(req.Data)
	broker.WriteMsg(req.Service, req.Method, string(data))

	print(req.Service)
}

func GenerateGateId() {
	GateId = "gid_" + RandStringRunes(GATE_ID_LEN)
	log.Println("Gate ID: ", GateId)
}
