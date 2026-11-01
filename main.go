package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	targetPort := "12345"
	listenPort := "8080"

	if len(os.Args) < 2 {
		log.Fatal("Address cannot be empty")
	} else if len(os.Args) == 3 {
		targetPort = os.Args[1]
		listenPort = os.Args[2]
	} else {
		targetPort = os.Args[1]
	}

	listen := "0.0.0.0:" + listenPort
	target := "http://127.0.0.1:" + targetPort

	u, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	log.Printf("proxying %s -> %s", target, listen)
	log.Fatal(http.ListenAndServe(listen, proxy))

}
