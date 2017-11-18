package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Here we go!")


	http.HandleFunc("/",handleIndex) //set router
	http.ListenAndServe(":9090",nil) //set listen port
}

func handleIndex(w http. ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "<html><header><title>This is title</title></header><body>Hello world</body></html>")
}
