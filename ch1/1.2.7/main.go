package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// hello world V2.0
func main() {
	fmt.Println("Please visit http://127.0.0.1:12345")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		s := fmt.Sprintf("hello world! time : %s", time.Now().String())
		fmt.Fprintf(writer, "%v\n", s)
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
