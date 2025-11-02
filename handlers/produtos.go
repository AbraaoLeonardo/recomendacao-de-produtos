package handlers

import (
	"encoding/json"
	"net/http"
	"recomendacao/database"
	"recomendacao/models"
)

func GetVendasHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, produto_id, quantidade, data_venda FROM vendas"
	rows, err := database.QueryDB(query)
	if err != nil {
		database.HandleDBError(w, err)
		return
	}
	defer rows.Close()

	var vendas []models.Venda
	for rows.Next() {
		var venda models.Venda
		err := rows.Scan(&venda.ID, &venda.ProdutoID, &venda.Quantidade, &venda.DataVenda)
		if err != nil {
			database.HandleDBError(w, err)
			return
		}
		vendas = append(vendas, venda)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}