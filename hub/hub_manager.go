package hub

import (
	"encoding/json"
	"strings"

	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/go-dashboard/repository"
	"github.com/tapvanvn/godashboard"
)

var __hubmap map[string]*Hub = map[string]*Hub{}

var unregister chan *Client = make(chan *Client)
var register chan *Client = make(chan *Client)
var clients map[*Client]bool = make(map[*Client]bool)
var broadcast chan []byte = make(chan []byte)

func GetHub(itemID string) *Hub {
	itemID = strings.ToLower(strings.TrimSpace(itemID))
	if hub, ok := __hubmap[itemID]; ok {
		return hub
	}
	item, err := repository.GetItem(itemID)
	if err == nil {
		item = &entity.Item{
			Name:           itemID,
			Title:          itemID,
			SignalTime:     0,
			SignalDuration: 1,
			Signal:         make(map[string]godashboard.Param),
		}
	}
	hub := &Hub{Item: item, LastWriteTime: 0}
	__hubmap[itemID] = hub
	return hub
}
func Signal(signal *godashboard.Signal) {

	itemID := strings.TrimSpace(signal.ItemName)
	if len(itemID) > 0 {

		h := GetHub(itemID)
		h.Signal(signal.Params)
	}
	parseParams := map[string]string{}
	for key, param := range signal.Params {
		if param.Type == "string" {
			parseParams[key] = string(param.Value)
		} else if param.Type == "int" {
			parseParams[key] = string(param.Value)
		} else if param.Type == "uint265" {
			parseParams[key] = string(param.Value)
		} else {
			parseParams[key] = string(param.Value)
		}
	}
	clientSignal := &entity.ClientSignal{
		ItemName: itemID,
		Params:   parseParams,
	}
	data, err := json.Marshal(clientSignal)
	if err != nil {
		return
	}
	broadcast <- data
}

func Run() {
	for {
		select {
		case client := <-register:
			clients[client] = true

		case client := <-unregister:

			delete(clients, client)

		case message := <-broadcast:

			for client := range clients {
				select {
				case client.send <- message:
				default:
					delete(clients, client)
				}
			}

		}
	}
}
