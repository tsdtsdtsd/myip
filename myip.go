package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const defaultPort string = "80"

var flagPort string

func main() {

	initFlags()
	addr := ":" + flagPort

	slog.Info("Starting myip service", "port", flagPort)

	http.HandleFunc("/", handle)
	http.HandleFunc("/json", handleJSON)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("failed to start", "error", err)
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
	fmt.Fprint(w, requestIP(r))
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
