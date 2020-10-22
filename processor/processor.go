package processor

import (
	"encoding/json"
	broker "github.com/IT108/achieve-broker-go"
	. "github.com/IT108/achieve-gateway-go/serialization"
	. "github.com/IT108/achieve-gateway-go/utils"
	achieve_models_go "github.com/IT108/achieve-models-go"
	"log"
	"strconv"
)

var GateId string

const GATE_ID_LEN = 10

func Process(req *AppRequest) {
	request := achieve_models_go.Request{
		RequestId: strconv.FormatInt(req.RequestId, 10),
		Method:    req.Method,
		Sender:    req.ClientId,
		User:      "",
		GateId:    GateId,
		Data:      "",
	}

	req.Data["request"] = request
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
