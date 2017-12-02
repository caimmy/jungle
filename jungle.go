package jungle

var (
	JungleApp 			*JungleRootApplication
	TemplatesPath		string
)

func init() {
	JungleApp = NewJungleApp()
}

func Run() {
	JungleApp.Run()
}

func NewJungleApp() *JungleRootApplication {
	app := &JungleRootApplication{}
	return app
}

// raw add router information for http server
func Router(prefix string, controller ControllerInterface) {
	Global_JungleHttpServerHandler.Add(prefix, controller)
}
