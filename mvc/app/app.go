package app

import (
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", mvc.controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
