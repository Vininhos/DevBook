package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))

	fmt.Println("Rodando WebApp na porta 3000...")
}
