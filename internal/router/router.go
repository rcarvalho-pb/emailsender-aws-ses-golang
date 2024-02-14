package router

import (
	"github.com/gorilla/mux"
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/router/routes"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	routes.Config(r)
	return r
}
