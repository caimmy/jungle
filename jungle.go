package jungle



func Run() {

	new_root_app := root_app{}
	new_root_app.Run()

}

// raw add router information for http server
func Router(prefix string, controller ControllerInterface) {
	Global_JungleHttpServerHandler.Add(prefix, controller)
}
