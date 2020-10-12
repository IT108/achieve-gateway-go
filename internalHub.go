package main

import (
	broker "github.com/IT108/achieve-broker-go"
	"github.com/IT108/achieve-gateway-go/processor"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type gateRouter struct {
	router broker.RouterBase
}

func (receiver *gateRouter) RunAction(data *kafka.Message) {
	key := string(data.Key)
	msg := &Message{
		clientID: key,
		msg:      data.Value,
	}
	clientsHub.send <- msg
}

func startConsumer() {
	base := broker.RouterBase{}
	authRouter := &gateRouter{base}
	broker.AssignRouter(broker.RouterInterface(authRouter))
	broker.Subscribe([]string{processor.GateId}, processor.GateId)
}