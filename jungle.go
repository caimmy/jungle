package jungle

import (
	"net/http"
	"log"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "asdfasdfads", 404)
}


func Run() {

	MainServer := &http.Server{}
	MainServer.Handler = &JungleController{}
	MainServer.Addr = ":8080"
	err := MainServer.ListenAndServe()
	if (err == nil) {
		log.Fatal("MainServer ListenAndServe : ", err.Error())
	}

}