package entity

import "github.com/tapvanvn/godashboard"

type Item struct {
	Name           string                       `json:"name" bson:"name"`
	Title          string                       `json:"title" bson:"title"`
	SignalTime     int64                        `json:"signal_time" bson:"signal_time"`
	SignalDuration int64                        `json:"signal_duration" bson:"signal_duration"`
	Signal         map[string]godashboard.Param `json:"signal" bson:"signal"`
}

func (doc Item) GetID() string {
	return doc.Name
}
