// to run this: inside of your go workspace use the command `go run server.go` assuming youre at the proper level in the directory structure
package main

import "net/http"

func main() {
	http.HandleFunc("/", frontPageFunc)
	// I can also add more routes here such as http.HandleFunc("/users/johnEappleSeed/mostFrequent", anotherFunc)
	http.ListenAndServe(":8080", nil)
}

func frontPageFunc(w http.ResponseWriter, req *http.Request) { // w http.ResponseWriter is going to write back to the request and req is the request
	w.Write([]byte("Welcome to Food.Now"))
}
