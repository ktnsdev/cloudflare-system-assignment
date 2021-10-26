package controllers

import (
	"cloudflare-system-assessment/config"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetReadMe(w http.ResponseWriter, r *http.Request) {
	readme, err := ioutil.ReadFile("README.txt")
	if err != nil {
		config.Logger(r.URL.Path, http.StatusInternalServerError, "Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal Server Error"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
		return
	}

	config.Logger(r.URL.Path, http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(readme)
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
