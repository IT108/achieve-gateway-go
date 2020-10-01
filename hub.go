package main

import "./processor"
import "./serialization"

// Hub maintains the set of active clients and broadcasts messages to the
// clients
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Authorized users.
	users map[string][]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
		users:      make(map[string][]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:

			h.registerClient(client)
		case client := <-h.unregister:
			h.unregisterClient(client)
		case message := <-h.broadcast:
			processor.Process(serialization.Deserialize(string(message)))

			//for _, client := range h.clients {
			//	select {
			//	case client.send <- message:
			//	default:
			//		close(client.send)
			//		delete(h.clients, client.clientId)
			//	}
			//}
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	h.clients[client.clientId] = client
	println(client.clientId)
}

func (h *Hub) unregisterClient(client *Client) {
	if _, ok := h.clients[client.clientId]; ok {
		delete(h.clients, client.clientId)
		close(client.send)
	}
}
