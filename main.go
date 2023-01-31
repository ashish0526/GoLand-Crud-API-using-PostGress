package main

import (
	"fmt"
	"go-postgress/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Server on port 8083 ...")

	log.Fatal(http.ListenAndServe(":8083", r))

}
