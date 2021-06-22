package repository

import engines "github.com/tapvanvn/godbengine"

var CollectionItem string = "dashboard_item"

func IsNoRecordError(err error) bool {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.IsNoRecordError(err)
}
