package processor

import (
	. "../serialization"
	. "../utils"
	"encoding/json"
	broker "gopkg.in/IT108/achieve-broker-go.v0"
	"log"
)

var GateId string

const GATE_ID_LEN = 10

func Process(req *AppRequest) {
	req.Data["GateId"] = GateId
	req.Data["Sender"] = req.ClientId
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