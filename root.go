package jungle

import (
	"net/http"
	"reflect"
	"github.com/caimmy/jungle/html"
	"github.com/caimmy/jungle/plugins/logger"
	"fmt"
	"os"
	"os/signal"
	"github.com/caimmy/jungle/plugins/session"
)

var (
	Global_JungleHttpServerHandler  JungleHttpServerHandler
	End_Application 				chan os.Signal
)

func init() {
	Global_JungleHttpServerHandler = *NewJungleHttpServerHandler()
}

func NewJungleHttpServerHandler() *JungleHttpServerHandler {
	gen_jungleHttpServerHandler := JungleHttpServerHandler{}
	gen_jungleHttpServerHandler.routers = make(map[string] reflect.Type)
	gen_jungleHttpServerHandler.ws_routers = make(map[string] func(w http.ResponseWriter, r *http.Request))
	return &gen_jungleHttpServerHandler
}

type JungleRootApplication struct {
	Server 			*http.Server
	TemplatePath	string

	TemplateManager *html.TemplatesManager
	LoggerManager 	*logger.LoggingManager
	SessionManager	session.SessionMgrInterface
}

var End_run chan bool

func (app *JungleRootApplication) Run(listen_serv string) {
	main_http_server := http.Server{}

	main_http_server.Addr = listen_serv
	main_http_server.Handler = &Global_JungleHttpServerHandler
	go func() {
		main_http_server.ListenAndServe()
		fmt.Println("web server started")
	}()
	End_Application = make(chan os.Signal, 1)
	signal.Notify(End_Application, os.Interrupt, os.Kill)

	c := <- End_Application
	app.Cleanup()
	fmt.Println("Got signal: ", c)
}

// Do cleanup function for Jungle Application
func (app *JungleRootApplication) Cleanup() {
	JungleApp.LoggerManager.StopRecord()
}