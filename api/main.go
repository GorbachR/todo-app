package main

import (
	"flag"
	"fmt"
	"github.com/GorbachR/todo-app/api/data"
	"github.com/GorbachR/todo-app/api/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	listendAddr := flag.String("listenAddr", ":3000", "the server address")

	r := gin.Default()
	routes.SetupRoutes(r)
	data.InitDbConnection()

	fmt.Println("Server running on port:", *listendAddr)
	log.Fatal(r.Run("127.0.0.1" + *listendAddr))
}
