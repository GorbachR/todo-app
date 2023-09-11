package main

import (
	"flag"
	"fmt"
	"github.com/GorbachR/todo-app/api/server"
	"log"
)

func main() {

	listendAddr := flag.String("listenAddr", ":3000", "the server address")

	server := server.SetupServer(*listendAddr)

	fmt.Println("Server running on port:", *listendAddr)
	log.Fatal(server.ListenAndServe())
}
