package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tapvanvn/go-dashboard/entity"
	"github.com/tapvanvn/go-dashboard/hub"
	"github.com/tapvanvn/go-dashboard/route"
	"github.com/tapvanvn/go-dashboard/system"
	"github.com/tapvanvn/go-dashboard/utility"
	"github.com/tapvanvn/gopubsubengine"
	"github.com/tapvanvn/gopubsubengine/wspubsub"
	"github.com/tapvanvn/gorouter/v2"
)

type Handles []gorouter.RouteHandle
type Endpoint gorouter.EndpointDefine

var subscriber gopubsubengine.Subscriber = nil
var pubsubhub gopubsubengine.Hub = nil

func OnDashboardMessage(message string) {
	signal := &entity.Signal{}
	err := json.Unmarshal([]byte(message), signal)
	if err != nil {
		return
	}

}
func InitPubSub() {

	pubsubConnectString := utility.MustGetEnv("CONNECT_STRING_WSPUBSUB")

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

func main() {
	var port = utility.MustGetEnv("PORT")
	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	system.RootPath = rootPath
	configFile := utility.GetEnv("CONFIG")
	if configFile == "" {
		configFile = "config.json"
	}
	//MARK: init system config
	jsonFile2, err := os.Open(rootPath + "/config/" + configFile)

	if err != nil {

		panic(err)
	}

	defer jsonFile2.Close()
	bytes, err := ioutil.ReadAll(jsonFile2)
	systemConfig := entity.Config{}

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &systemConfig)
	if err != nil {
		panic(err)
	}
	system.Config = &systemConfig

	if err != nil {

		panic(err)
	}
	//MARK: init router
	jsonFile, err := os.Open(rootPath + "/config/route.json")

	if err != nil {

		panic(err)
	}

	defer jsonFile.Close()

	bytes, _ = ioutil.ReadAll(jsonFile)
	var handers = map[string]gorouter.EndpointDefine{

		"":         {Handles: Handles{route.Root}},
		"unhandle": {Handles: Handles{route.Unhandle}},
	}

	var router = gorouter.Router{}

	router.Init("v1/", string(bytes), handers)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("i am ok"))
	})
	http.Handle("/v1/", router)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(w, r)
	})

	fileServer := http.FileServer(utility.FileSystem{FS: http.Dir(rootPath + "/static")})
	http.Handle("/", fileServer)

	go hub.Run()

	fmt.Println("listen on port", port)

	//init pubsub
	InitPubSub()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
