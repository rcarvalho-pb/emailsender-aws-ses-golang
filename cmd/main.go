package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/router"
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/router/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	r := router.GenerateRouter()
	routes.Config(r)
	fmt.Printf("Serve's running on port :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
