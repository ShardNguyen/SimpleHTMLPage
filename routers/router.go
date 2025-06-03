package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func Run() {
	userRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Serve index.html at root
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.ListenAndServe(":8080", r)
}
