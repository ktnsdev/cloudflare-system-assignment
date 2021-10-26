package controllers

import (
	"cloudflare-system-assessment/config"
	"cloudflare-system-assessment/internal/helper"
	"fmt"
	"net/http"
	"time"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	responseString := "Average Encoding Runtime (/auth): " + helper.GetAverageEncodingRuntime().String()
	responseString += "\nAverage Decoding Runtime (/verify, middleware): " + helper.GetAverageDecodingRuntime().String() + "\nAll logs:"

	runtimes := helper.GetAllRuntimes()
	for i := 0; i < len(runtimes); i++ {
		responseString += "\n" + runtimes[i].Timestamp.Format(time.StampMilli) + "\t" + runtimes[i].Path + "\t" + runtimes[i].Runtime.String()
	}

	config.Logger(r.URL.Path, http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(responseString))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
