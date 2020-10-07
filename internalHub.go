package main

import (
	"./processor"
	broker "gopkg.in/IT108/achieve-broker-go.v0"
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