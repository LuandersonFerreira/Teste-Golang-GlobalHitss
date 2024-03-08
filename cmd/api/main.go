package main

import (
	"fmt"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/configs"
	"github.com/LuandersonFerreira/teste-golang-global/internal/rabbitmq"
	"github.com/LuandersonFerreira/teste-golang-global/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	routes.SetRoutes(r)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
		if err != nil {
			panic(err)
		}
	}()

	err = rabbitmq.InitConsumer()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
