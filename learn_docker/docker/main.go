package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello docker!")
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.ListenAndServe(":8012",nil)
}
