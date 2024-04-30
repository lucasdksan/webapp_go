package main

import (
	"fmt"
	"net/http"
	"webapp_go/src/config"
	"webapp_go/src/cookies"
	"webapp_go/src/router"
	"webapp_go/src/utils"
)

func main() {
	r := router.Generate()

	utils.Load_Templates()
	config.LoadingEnv()
	cookies.Configure()

	fmt.Printf("Escutando na porta %d", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)
}
