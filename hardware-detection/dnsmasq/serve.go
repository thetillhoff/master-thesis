package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	ip   = "0.0.0.0"
	port = 80
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	fmt.Println("Listening on " + ip + ":" + strconv.Itoa(port) + " ...")
	err := http.ListenAndServe(ip+":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
