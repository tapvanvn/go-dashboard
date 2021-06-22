package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tapvanvn/go-dashboard/route/response"
	"github.com/tapvanvn/gorouter/v2"
)

//Unhandle handle unhandling route
func Unhandle(context *gorouter.RouteContext) {

	fmt.Println("cannot handle:", context.Path)

	context.W.WriteHeader(http.StatusNotFound)

	responseData := response.Data{Success: false,
		ErrorCode: 0,
		Message:   "Route to nowhere",
		Data:      nil}

	if data, err := json.Marshal(responseData); err == nil {

		context.W.Write(data)
	}
	context.Handled = true
}

//Root handle root
func Root(context *gorouter.RouteContext) {

	if context.Action == "healthz" {
		context.W.WriteHeader(http.StatusOK)
		context.W.Write([]byte("i am ok"))
		context.Handled = true
		return
	} else if context.Action == "new" {

	} else if context.Action == "update" {

	}
}
