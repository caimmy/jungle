package jungle

import (
	"net/http"
	"reflect"
	"github.com/caimmy/jungle/html"
)

var Global_JungleHttpServerHandler  JungleHttpServerHandler

func init() {
	Global_JungleHttpServerHandler = NewJungleHttpServerHandler()
}

func NewJungleHttpServerHandler() JungleHttpServerHandler {
	return JungleHttpServerHandler{
		routers: make(map[string] reflect.Type),

	}
}

type JungleRootApplication struct {
	Server 			*http.Server
	TemplatePath	string

	TemplateManager *html.TemplatesManager
}

var End_run chan bool

func (app *JungleRootApplication) Run() {
	main_http_server := http.Server{}

	main_http_server.Addr = ":8081"
	main_http_server.Handler = &Global_JungleHttpServerHandler
	go func() {
		main_http_server.ListenAndServe()
	}()
	<- End_run
}