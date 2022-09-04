package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tapvanvn/go-dashboard/common"
	"github.com/tapvanvn/go-dashboard/runtime"
	"github.com/tapvanvn/go-dashboard/utility"
	"github.com/tapvanvn/godashboard"
	"github.com/tapvanvn/goutil"
)

var count = 1

func GetSignalParams() map[string]godashboard.Param {
	report := map[string]godashboard.Param{}

	report["lastest"] = godashboard.Param{
		Type:  "int",
		Value: []byte(fmt.Sprintf("%d", count)),
	}
	count++
	return report
}
func reportLive() {
	signal := &godashboard.Signal{ItemName: "sender." + runtime.NodeName,
		Params: GetSignalParams(),
	}
	godashboard.Report(signal)
}

func main() {
	var port = goutil.MustGetEnv("PORT")

	rootPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	runtime.RootPath = rootPath
	configPath := utility.GetGeneralConfigPath(rootPath)

	if err := runtime.InitEngine(configPath, common.EmptyModules); err != nil {
		panic(err)
	}

	dashboard := &godashboard.Dashboard{
		Type:             runtime.Config.Hub.Type,
		ConnectionString: runtime.Config.Hub.Endpoint,
	}
	godashboard.AddDashboard(dashboard)

	goutil.Schedule(reportLive, time.Second*1)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
