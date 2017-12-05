package jungle

import (
	"github.com/caimmy/jungle/html"
	"github.com/caimmy/jungle/plugins/logger"
)

// Config params for Top Application
var (
	JungleApp 			*JungleRootApplication
	TemplatesPath		string
	LogPath				string

	RedisServer			string
	RedisDb				int
)

func init() {
	TemplatesPath 		= "templates"
	LogPath				= "logs"
	JungleApp = NewJungleApp()
}

func Run() {
	JungleApp.Run()
}

func NewJungleApp() *JungleRootApplication {
	app := &JungleRootApplication{
		TemplateManager: html.NewTemplatesManager(),
		LoggerManager: logger.NewLoggingManager(LogPath),
	}
	return app
}

// raw add router information for http server
func Router(prefix string, controller ControllerInterface) {
	Global_JungleHttpServerHandler.Add(prefix, controller)
}
