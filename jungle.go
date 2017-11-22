package jungle

import (
	"net/http"
	"io"
	"log"
)

func indexPage(w http.ResponseWriter, r http.Request) {
	io.WriteString(w, "hello jungle!")
}

func run() {

	http.HandleFunc("/", indexPage)
	err := http.ListenAndServe("8080", nil)
	if (err != nil) {
		log.Fatal("ListenAndServ: ", err.Error())
	}
}