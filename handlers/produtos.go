package handlers

import (
	"log"
	"net/http"
	"recomendacao/database"
	"time"

	"github.com/gin-gonic/gin"
)

type ProdutoBusca struct {
	CodigoInterno string `json:"codigo_interno" binding:"required"`
	CodigoBarras  string `json:"codigo_barras" binding:"required"`
}

type ProdutoAssociado struct {
	CodigoInterno string `json:"codigo_interno"`
	CodigoBarras  string `json:"codigo_barras"`
	TotalVendas   int    `json:"total_vendas"`
}

func BuscarProdutosVendidosJuntosHandler(c *gin.Context) {
	var busca ProdutoBusca

	if err := c.ShouldBindJSON(&busca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados de busca inválidos", "detalhes": err.Error()})
		return
	}

	dataLimite := time.Now().AddDate(0, -3, 0)

	query := `
		SELECT
			iv_assoc.codigo_interno,
			iv_assoc.codigo_barras,
			COUNT(iv_assoc.venda_id) AS total_vendas
		FROM 
			itens_venda iv_busca
		JOIN 
			vendas v ON iv_busca.venda_id = v.id
		JOIN 
			itens_venda iv_assoc ON iv_busca.venda_id = iv_assoc.venda_id
		WHERE
			iv_busca.codigo_interno = $1 
			AND iv_busca.codigo_barras = $2
			AND v.data_venda >= $3
			AND iv_assoc.codigo_interno != $1
		GROUP BY 
			iv_assoc.codigo_interno, iv_assoc.codigo_barras
		ORDER BY 
			total_vendas DESC
	`

	rows, err := database.DB.Query(query, busca.CodigoInterno, busca.CodigoBarras, dataLimite)
	if err != nil {
		log.Printf("Erro na consulta de produtos vendidos juntos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha ao consultar associações no banco de dados", "detalhes": err.Error()})
		return
	}
	defer rows.Close()

	var associados []ProdutoAssociado
	for rows.Next() {
		var p ProdutoAssociado
		if err := rows.Scan(&p.CodigoInterno, &p.CodigoBarras, &p.TotalVendas); err != nil {
			log.Printf("Erro ao escanear linha: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Falha interna ao processar resultados"})
			return
		}
		associados = append(associados, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro após iteração de linhas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao processar resultados da consulta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"produto_buscado": busca,
		"data_limite":     dataLimite.Format("2006-01-02"),
		"associados":      associados,
	})
}
