package handlers

import (
	"fmt"
	"log"
	"net/http"
	"recomendacao/database"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemVenda struct {
	CodigoInterno     string `json:"codigo_interno" binding:"required"`
	CodigoBarras      string `json:"codigo_barras" binding:"required"`
	QuantidadeVendida int    `json:"quantidade_vendida" binding:"required,min=1"`
}

type Venda struct {
	DataVenda string      `json:"data_venda" binding:"required"`
	Produtos  []ItemVenda `json:"produtos" binding:"required,min=1"`
}

func SalvarVendaHandler(c *gin.Context) {
	var venda Venda

	if err := c.ShouldBindJSON(&venda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos", "detalhes": err.Error()})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao iniciar a transação", "detalhes": err.Error()})
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	var vendaID int
	dataVendaParsed, err := time.Parse("2006-01-02 15:04:05", venda.DataVenda)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Formato de data inválido. Use YYYY-MM-DD HH:MM:SS", "detalhes": err.Error()})
		return
	}

	insertVendaSQL := `INSERT INTO vendas (data_venda) VALUES ($1) RETURNING id`
	err = tx.QueryRow(insertVendaSQL, dataVendaParsed).Scan(&vendaID)
	if err != nil {
		tx.Rollback()
		log.Printf("Erro ao inserir venda: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao salvar o cabeçalho da venda", "detalhes": err.Error()})
		return
	}

	insertItemSQL := `INSERT INTO itens_venda (venda_id, codigo_interno, codigo_barras, quantidade_vendida) VALUES ($1, $2, $3, $4)`
	for _, item := range venda.Produtos {
		_, err := tx.Exec(insertItemSQL, vendaID, item.CodigoInterno, item.CodigoBarras, item.QuantidadeVendida)
		if err != nil {
			tx.Rollback()
			log.Printf("Erro ao inserir item: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"erro": fmt.Sprintf("Falha ao salvar o item com código %s", item.CodigoInterno), "detalhes": err.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao confirmar a transação", "detalhes": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensagem": "Venda salva com sucesso!",
		"id_venda": vendaID,
		"produtos": len(venda.Produtos),
	})
}
