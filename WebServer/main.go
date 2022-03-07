package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi my name is Daiana, this is my first web with GO. Welcome to %s!", r.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http..ListenandServe(":8080",nil))

}