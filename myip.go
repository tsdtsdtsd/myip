package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

const defaultPort int = 1323

func main() {

	var flagPort int
	flag.IntVar(&flagPort, "p", defaultPort, "Webserver listening port")
	flag.Parse()

	if flagPort == 0 {
		flagPort = defaultPort
	}

	port := ":" + strconv.Itoa(flagPort)

	http.HandleFunc("/", handle)

	fmt.Printf("Starting myip service at port %d ...\n", flagPort)
	http.ListenAndServe(port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", userIP(r))
}

func userIP(r *http.Request) string {
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
