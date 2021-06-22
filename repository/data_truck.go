package repository

import (
	"fmt"

	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
)

//SaveDataTruck save data truck to database
func SaveDataTruck(dataTruck *engine.DataTruck) error {

	engine := engines.GetEngine()

	documentPool := engine.GetDocumentPool()

	transaction := documentPool.MakeTransaction()

	transaction.Begin()

	for _, item := range dataTruck.Items {

		fmt.Println("put item to " + item.Collection)
		transaction.Put(item.Collection, item.Document)
	}

	return transaction.Commit()
}
