package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		addr := req.RemoteAddr[:strings.LastIndexByte(req.RemoteAddr, ':')]
		addr = strings.Trim(addr, "[]")
		res.Write([]byte(addr))
		log.Println(addr)

	})
	log.Fatal(http.ListenAndServe(":2460", nil))
}
