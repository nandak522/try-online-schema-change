package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/none-da/try-online-schema-change/reader-web-app/pkg/reader"
)

// StartTime gives the start time of server
var StartTime = time.Now()

const defaultAppPort string = "8080"

func uptime() string {
	elapsedTime := time.Since(StartTime)
	return fmt.Sprintf("%d:%d:%d", int(math.Round(elapsedTime.Hours())), int(math.Round(elapsedTime.Minutes())), int(math.Round(elapsedTime.Seconds())))
}

func homePage(w http.ResponseWriter, req *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, fmt.Sprintf("[HOST: %s] (uptime: %s)]", host, uptime()))
}

func readData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, reader.ReadData())
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/read", readData)

	http.ListenAndServe(fmt.Sprintf(":%s", defaultAppPort), nil)
}
