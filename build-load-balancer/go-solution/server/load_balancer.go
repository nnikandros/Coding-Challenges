package server

import (
	"fmt"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", RootHandler)

	return mux
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello world")

}
