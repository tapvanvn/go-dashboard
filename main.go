package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tapvanvn/go-dashboard/hub"
	"github.com/tapvanvn/go-dashboard/route"
	"github.com/tapvanvn/go-dashboard/system"
	"github.com/tapvanvn/godashboard"
	engines "github.com/tapvanvn/godbengine"
	"github.com/tapvanvn/godbengine/engine"
	"github.com/tapvanvn/godbengine/engine/adapter"
	"github.com/tapvanvn/gopubsubengine"
	"github.com/tapvanvn/gopubsubengine/wspubsub"
	"github.com/tapvanvn/gorouter/v2"
	"github.com/tapvanvn/goutil"
)

type Handles []gorouter.RouteHandle
type Endpoint gorouter.EndpointDefine

var subscriber gopubsubengine.Subscriber = nil
var pubsubhub gopubsubengine.Hub = nil

func OnDashboardMessage(message string) {

	signal := &godashboard.Signal{}
	err := json.Unmarshal([]byte(message), signal)
	if err != nil {
		return
	}
	hub.Signal(signal)
}
func InitPubSub() {

	pubsubConnectString := system.Config.Hub.Endpoint

	h, err := wspubsub.NewWSPubSubHub(pubsubConnectString)

	if err != nil {

		panic(err)
	}
	pubsubhub = h
	sub, err := pubsubhub.SubscribeOn("dashboard")
	if err != nil {

		panic(err)
	}
	subscriber = sub
	subscriber.SetProcessor(OnDashboardMessage)
}

//Start start engine
func StartEngine(eng *engine.Engine) {

	connectString := system.Config.DocumentDB.ConnectionString
	var documentPool engine.DocumentPool = nil
	if system.Config.DocumentDB.Type == "mongodb" {

		mongoPool := &adapter.MongoPool{}
		err := mongoPool.InitWithDatabase(connectString, system.Config.DocumentDB.Database)
		if err != nil {
			panic("cannot init mongo")
		}

	} else if system.Config.DocumentDB.Type == "firestore" {
		firestorePool := &adapter.FirestorePool{}
		err := firestorePool.Init(connectString)
		if err != nil {
			panic("cannot init firestore")
		}
		documentPool = firestorePool
	}
	eng.Init(nil, documentPool, nil)
}

func InitRouter() {
	//MARK: init router

	var handers = map[string]gorouter.EndpointDefine{

		"":         {Handles: Handles{route.Root}},
		"unhandle": {Handles: Handles{route.Unhandle}},
	}

	var router = gorouter.Router{}
	routePrefix := "v1/"

	router.Init(routePrefix, string("{}"), handers)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("i am ok"))
	})
	http.Handle("/v1/", router)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(w, r)
	})
	cacheFileServer := goutil.NewCacheFileServer(http.Dir(system.RootPath + "/static"))

	fileServer := http.FileServer(cacheFileServer)

	http.Handle("/", fileServer)
}

func main() {

	var port = goutil.MustGetEnv("PORT")

	rootPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	system.RootPath = rootPath

	configPath := goutil.GetEnv("CONFIG_PATH")

	if configPath == "" {

		if rootPath != "/" {

			configPath = fmt.Sprintf("%s/config/%s", system.RootPath, "config.jsonc")

		} else {

			configPath = fmt.Sprintf("/config/%s", "config.jsonc")
		}
	}
	fmt.Println("config from:", configPath)

	system.LoadConfig(configPath)

	engines.InitEngineFunc = StartEngine

	_ = engines.GetEngine()

	InitRouter()

	go hub.Run()

	fmt.Println("listen on port", port)

	//init pubsub
	InitPubSub()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
