package jungle

import (
	"net/http"
	"reflect"
)

var Global_JungleHttpServerHandler  JungleHttpServerHandler

func init() {
	Global_JungleHttpServerHandler = NewJungleHttpServerHandler()
}

func NewJungleHttpServerHandler() JungleHttpServerHandler {
	return JungleHttpServerHandler{routers: make(map[string] reflect.Type)}
}

type root_app struct {
	Server 		*http.Server
}

var End_run chan bool

func (app *root_app) Run() {
	main_http_server := http.Server{}
	main_http_server.Addr = ":8081"
	main_http_server.Handler = &Global_JungleHttpServerHandler
	go func() {
		main_http_server.ListenAndServe()
	}()
	<- End_run
}