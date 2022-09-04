package runtime

import (
	"fmt"
	"log"

	"github.com/tapvanvn/go-dashboard/common"
	"github.com/tapvanvn/go-dashboard/entity"
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
	"github.com/tapvanvn/godbengine/engine/adapter"
	"github.com/tapvanvn/goutil"
)

const (
	ENV_DEV   = 0
	ENV_PROD  = 1
	ENV_LOCAL = 2
)

var RootPath string = ""
var NodeName string = goutil.GenVerifyCode(5)
var Modules map[common.MODULE]bool = make(map[common.MODULE]bool)

var Environment = ENV_PROD
var Config *entity.Config = nil

func IsDev() bool {
	return Environment == ENV_DEV
}

func IsProd() bool {
	return Environment == ENV_PROD
}

func IsLocal() bool {
	return Environment == ENV_LOCAL
}

func InitEngine(configPath string, modules map[common.MODULE]bool) error {

	NodeName = goutil.GenVerifyCode(5)
	var err error = nil

	if Config, err = ReadConfig(configPath); err != nil {
		return err
	}

	engines.InitEngineFunc = startEngine
	_ = engines.GetEngine()

	for module, _ := range modules {
		if err := loadModule(module); err != nil {

			return err
		}
	}
	//Add route depend on module loaded

	return nil
}

func startEngine(eng *engine.Engine) {
	env := Config.Environment

	Environment = ENV_PROD

	switch env {
	case "dev":
		Environment = ENV_DEV

	case "local":
		Environment = ENV_LOCAL
	}
	//read redis define from env
	var memdb engine.MemPool = nil

	var documentDB engine.DocumentPool = nil
	var filePool engine.FilePool = nil
	if Config.DocDB != nil {

		connectString := Config.DocDB.ConnectionString
		databaseName := Config.DocDB.Database

		if Config.DocDB.Provider == "mongodb" {

			mongoPool := &adapter.MongoPool{}
			err := mongoPool.InitWithDatabase(connectString, databaseName)

			if err != nil {

				log.Fatal("cannot init mongo")
			}
			documentDB = mongoPool

			fileMongoPool := adapter.MongoFilePool{}
			fileMongoPool.Init("file", mongoPool)
			filePool = fileMongoPool

		} else {

			firestorePool := adapter.FirestorePool{}
			firestorePool.Init(connectString)
			documentDB = &firestorePool
		}
	}

	eng.Init(memdb, documentDB, filePool)

	if !IsProd() {
		adapter.SetMeasurement(true)
	}
}

func loadModule(moduleName common.MODULE) error {
	fmt.Println("load", moduleName)

	if _, loaded := Modules[moduleName]; loaded {
		return nil
	}
	deps := common.ModuleDependencySolver.GetDependency(string(moduleName))
	for _, dep := range deps {
		if err := loadModule(common.MODULE(dep)); err != nil {
			return err
		}
	}
	/*switch moduleName {
	case common.MODULE_CHECK_DATABASE:
		if err := repository.CheckRepository(); err != nil {

			return err
		}
	}*/
	Modules[moduleName] = true
	return nil
}
