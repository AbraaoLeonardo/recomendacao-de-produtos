package main

import (
	"log"

	"recomendacao/database"
	"recomendacao/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	defer database.CloseDB() 

	r := gin.Default()

	r.POST("/vendas", handlers.SalvarVendaHandler)

	log.Println("Servidor iniciado na porta 8080. Utilizando Connection Pooling...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}