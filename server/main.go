package main

import (
	"dyx_xy/server/common"
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Print("form:", r.Form)

	fmt.Fprintf(w, "hello dyx_xy!")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8888", nil)
	common.OnError(err, "")
}
