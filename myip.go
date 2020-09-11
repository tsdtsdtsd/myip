package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
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

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userIP := net.ParseIP(ip)

	if userIP == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "userIP: %s\n", spew.Sdump(userIP))
	fmt.Fprintf(w, "userIP.String: %s\n", userIP.String())
}
