package repository

import (
	"github.com/tapvanvn/go-dashboard/entity"
	engines "github.com/tapvanvn/godbengine"
)

func GetRegistry(rootName string) (*entity.Registry, error) {
	pool := engines.GetEngine().GetDocumentPool()
	registry := &entity.Registry{}
	err := pool.Get(CollectionRegistry, rootName, registry)
	if err != nil {
		if IsNoRecordError(err) {
			registry = entity.NewRegistry(rootName)
		} else {
			return nil, err
		}
	} else {
		registry.Hashes = make(map[string]*entity.Branch)
	}
	return registry, nil
}

func PutRegistry(registry *entity.Registry) error {

	pool := engines.GetEngine().GetDocumentPool()

	return pool.Put(CollectionRegistry, registry)
}
