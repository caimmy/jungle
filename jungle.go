package jungle

import (
	"github.com/caimmy/jungle/html"
	"github.com/caimmy/jungle/plugins/logger"
	"github.com/caimmy/jungle/plugins/session"
)

// Config params for Top Application
var (
	JungleApp 			*JungleRootApplication
	TemplatesPath		string
	LogPath				string
	SessionPath			string

	RedisServer			string
	RedisDb				int
)

func init() {
	TemplatesPath 		= "templates"
	LogPath				= "logs"
	SessionPath			= "sessions"
	JungleApp = NewJungleApp()
}

func Run() {
	JungleApp.Run()
}

func NewJungleApp() *JungleRootApplication {
	app := &JungleRootApplication{
		TemplateManager: html.NewTemplatesManager(),
		LoggerManager: logger.NewLoggingManager(LogPath),
		SessionManager:session.NewSessionManager(session.FILE_SESSION, 0, map[string]interface{}{"path": SessionPath}),
	}
	return app
}

// raw add router information for http server
func Router(prefix string, controller ControllerInterface) {
	Global_JungleHttpServerHandler.Add(prefix, controller)
}
