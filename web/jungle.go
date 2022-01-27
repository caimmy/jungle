package web

import (
	"fmt"
	"os"
	"strings"

	"github.com/caimmy/jungle/html"
	"github.com/caimmy/jungle/plugins/blueprint"
	"github.com/caimmy/jungle/plugins/logger"
	"github.com/caimmy/jungle/plugins/session"
)

// Config params for Top Application
var (
	JungleApp     *JungleRootApplication
	TemplatesPath string
	LogPath       string
	SessionPath   string

	SessionOn bool
	// Session life duration on server 30Min
	SessDuration int64

	RedisServer string
	RedisDb     int
)

func init() {
	TemplatesPath = "templates"
	LogPath = "logs"
	SessionPath = "sessions"
	SessDuration = 1800
	SessionOn = false
}

func Run(listen_server string, log_path string) {
	JungleApp = NewJungleApp(log_path)
	JungleApp.Run(listen_server)
}

func NewJungleApp(log_path string) *JungleRootApplication {
	if log_path == "" {
		_app_path, _ := os.Getwd()
		log_path = strings.Join([]string{_app_path, "logs"}, fmt.Sprintf("%c", os.PathSeparator))
	}
	fmt.Println(log_path)
	app := &JungleRootApplication{
		TemplateManager: html.NewTemplatesManager(),
		LoggerManager:   logger.NewLoggingManager(log_path),
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

func AddBlueprint(prefix string, blueprint *blueprint.Blueprint) {
	Global_JungleHttpServerHandler.AddBlueprint(prefix, blueprint)
}
