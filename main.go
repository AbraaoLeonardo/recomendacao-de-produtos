package main

import (
	"log"

	"recomendacao/database"
	"recomendacao/handlers"
	"recomendacao/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := utils.LoadConfigFromFile("./config/api.conf")
	serverPort := ":" + cfg.ServerPort

	database.InitDB(cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	defer database.CloseDB()

	r := gin.Default()

	r.POST("/vendas", handlers.SalvarVendaHandler)

	log.Printf("Servidor iniciado na porta %s. Utilizando Connection Pooling...", cfg.ServerPort)
	if err := r.Run(serverPort); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
