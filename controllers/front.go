package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func EncodeResponseAsJSON(data interface{}, writer io.Writer) {
	enc := json.NewEncoder(writer)
	_ = enc.Encode(data)
}
