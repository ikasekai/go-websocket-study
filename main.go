package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello go web</body></html>\n")
}

func main() {
	var httpServer http.Server
	//http.HandleFunc("/", handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("start http listening :8080")
	httpServer.Addr = ":8080"
	log.Println(httpServer.ListenAndServe())
}
