package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler -stores the pointer to our comments service
type Handler struct {
	Router *mux.Router
}

//NewHandler - returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//setuproutes sets up all routes for all applications
func (h *Handler) SetUpRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive! The Gabe Server is alive")
	})
}
