// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Map Callbacks from model
	callbacks map[string]func([]byte) []byte
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		callbacks:  make(map[string]func([]byte) []byte),
	}
}

type Message struct {
	Action string `json:"action"` //command, result, notify
	Data   string `json:"data"`
}

func (h *Hub) AddCallback(name string, cb func([]byte) []byte) {
	h.callbacks[name] = cb
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:

			var mes Message
			err := json.Unmarshal(message, &mes)
			if err != nil {
				fmt.Println(err)
				return
			}

			var message2 []byte
			var mes2 Message

			if f, ok := h.callbacks[mes.Action]; ok {
				mes2.Action = mes.Action
				mes2.Data = string(f([]byte(mes.Data)))
			}
			if message2, err = json.Marshal(&mes2); err != nil {
				continue
			}

			for client := range h.clients {
				select {
				case client.send <- message2:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}
