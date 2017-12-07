package jungle

import (
	"github.com/caimmy/jungle/html"
	"github.com/caimmy/jungle/plugins/logger"
	"github.com/caimmy/jungle/plugins/session"
	"github.com/caimmy/jungle/plugins/Blueprint"
)

// Config params for Top Application
var (
	JungleApp 			*JungleRootApplication
	TemplatesPath		string
	LogPath				string
	SessionPath			string

	SessionOn			bool
	// Session life duration on server 30Min
	SessDuration		int64

	RedisServer			string
	RedisDb				int
)

func init() {
	TemplatesPath 		= "templates"
	LogPath				= "logs"
	SessionPath			= "sessions"
	SessDuration		= 1800
	SessionOn			= false
}

func Run() {
	JungleApp = NewJungleApp()
	JungleApp.Run()
}

func NewJungleApp() *JungleRootApplication {
	app := &JungleRootApplication{
		TemplateManager: html.NewTemplatesManager(),
		LoggerManager: logger.NewLoggingManager(LogPath),
		//SessionManager: session.NewSessionManager(session.FILE_SESSION, SessDuration, map[string]interface{}{"path": SessionPath}),
	}

	if SessionOn {
		app.SessionManager = session.NewSessionManager(session.FILE_SESSION, SessDuration, map[string]interface{}{"path": SessionPath})
	}

	return app
}

// raw add router information for http server
func Router(prefix string, controller ControllerInterface) {
	Global_JungleHttpServerHandler.Add(prefix, controller)
}

func AddBlueprint(prefix string, blueprint *Blueprint.Blueprint) {
	Global_JungleHttpServerHandler.AddBlueprint(prefix, blueprint)
}
