package main

import (
	"dyx_xy/server/common"
	"dyx_xy/server/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, statusCode int, message any) {
	w.WriteHeader(statusCode)
	if statusCode != 200 {
		if message == "" {
			message = "fail"
		}

		log.Printf("[warn] statusCode=%d message='%v'", statusCode, message)
	}

	w.Write([]byte(fmt.Sprint(message)))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Print("form:", r.Form)

	fmt.Fprintf(w, "hello dyx_xy!")
}

func login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	{
		var name, password string
		var ok bool
		name, password, ok = r.BasicAuth()
		if ok {
			user.Name = name
			user.Password = password
		} else {
			bytes, err := io.ReadAll(r.Body)
			common.OnError(err, "")
			json.Unmarshal(bytes, &user)
		}
	}

	if ok, err := user.Verify(); err != nil {
		handle(w, 500, "")
	} else if ok {
		handle(w, 200, "passed verify")
	} else {
		handle(w, 400, "failed to pass verify")
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	bytes, err := io.ReadAll(r.Body)
	common.OnError(err, "")
	if err = json.Unmarshal(bytes, &user); err != nil {
		w.WriteHeader(400)
		return
	}

	if err := user.Rigerster(); err != nil {
		handle(w, 400, err)
	} else {
		handle(w, 200, "")
		log.Printf("successfully register user '%s'", user.Name)
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	err := http.ListenAndServe(":8888", nil)
	common.OnError(err, "")
}
