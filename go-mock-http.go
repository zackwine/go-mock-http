package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	port = flag.Int("port", 4532, "The port of the http server.")
)

func hostHTTP(port int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Printf("Failed to dump Post (%v)\n", err)
		}
		fmt.Printf("%s\n", requestDump)
	})

	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Printf("Failed to dump Delete (%v)\n", err)
		}
		fmt.Printf("%s\n", requestDump)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func main() {

	flag.Parse()

	hostHTTP(*port)
}
