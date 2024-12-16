package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const defaultPort string = "80"

var flagPort string

func main() {

	initFlags()
	addr := ":" + flagPort

	fmt.Printf("Starting myip service at port %s ...\n", flagPort)

	http.HandleFunc("/", handle)
	http.HandleFunc("/json", handleJSON)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("failed to start:", err)
		os.Exit(1)
	}
}

func initFlags() {

	flag.StringVar(&flagPort, "p", defaultPort, "Webserver listening port")
	flag.Parse()

	if flagPort == "" {
		flagPort = defaultPort
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", requestIP(r))
}

func handleJSON(w http.ResponseWriter, r *http.Request) {

	type response struct {
		IP string `json:"ip"`
	}

	var (
		resp = response{requestIP(r)}
	)

	respJSON, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

func requestIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	return r.RemoteAddr
}
