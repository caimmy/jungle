package jungle

import (
	"fmt"
	"jungle/html"
	"jungle/plugins/logger"
	"jungle/plugins/session"
	"net/http"
	"os"
	"os/signal"
	"reflect"
)

var (
	Global_JungleHttpServerHandler JungleHttpServerHandler
	End_Application                chan os.Signal
)

func init() {
	Global_JungleHttpServerHandler = NewJungleHttpServerHandler()
}

func NewJungleHttpServerHandler() JungleHttpServerHandler {
	return JungleHttpServerHandler{
		routers: make(map[string]reflect.Type),
	}
}

type JungleRootApplication struct {
	Server       *http.Server
	TemplatePath string

	TemplateManager *html.TemplatesManager
	LoggerManager   *logger.LoggingManager
	SessionManager  session.SessionMgrInterface
}

var End_run chan bool

func (app *JungleRootApplication) Run(listen_serv string) {
	main_http_server := http.Server{}

	main_http_server.Addr = listen_serv
	main_http_server.Handler = &Global_JungleHttpServerHandler
	go func() {
		main_http_server.ListenAndServe()
		fmt.Println("asdfasfsadf end")
	}()
	End_Application = make(chan os.Signal, 1)
	signal.Notify(End_Application, os.Interrupt, os.Kill)

	c := <-End_Application
	app.Cleanup()
	fmt.Println("Got signal: ", c)
}

// Do cleanup function for Jungle Application
func (app *JungleRootApplication) Cleanup() {
	JungleApp.LoggerManager.StopRecord()
}
