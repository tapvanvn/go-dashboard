package repository

import engines "github.com/tapvanvn/godbengine"

var CollectionItem string = "dashboard_item"
var CollectionAccount string = "dashboard_account"

func IsNoRecordError(err error) bool {

	eng := engines.GetEngine()

	documentPool := eng.GetDocumentPool()

	return documentPool.IsNoRecordError(err)
}
