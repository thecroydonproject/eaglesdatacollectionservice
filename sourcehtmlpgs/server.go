package main

import (
	"log"
	"net/http"
)


func main() {

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/home/aabdi/workspace/go/src/github.com/abdul2/eaglesdatacollectionservice/chung_1/"))))

}