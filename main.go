package main

import (
	"fmt"
	"net/http"
	"webapp_go/src/router"
	"webapp_go/src/utils"
)

func main() {
	r := router.Generate()

	utils.Load_Templates()

	fmt.Println("Escutando na porta 3000")
	http.ListenAndServe(":3000", r)
}
