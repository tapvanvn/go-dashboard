package repository

import (
	"github.com/tapvanvn/go-dashboard/entity"
	engines "github.com/tapvanvn/godbengine"
)

func GetItem(id string) (*entity.Item, error) {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	doc := &entity.Item{}

	err := documentPool.Get(CollectionItem, id, doc)

	if err != nil {

		return nil, err

	}
	return doc, nil
}

func PutItem(item *entity.Item) error {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.Put(CollectionItem, item)
}
