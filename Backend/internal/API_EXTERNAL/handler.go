package API_EXTERNAL

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var stocks []Stock

type SyncResponse struct {
	Cant   int `json:"Cant"`
	Estado int `json:"Estado"`
}

func SyncStocksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Iniciando sincronización de stocks")
	w.Header().Set("Content-Type", "application/json")
	response := SyncResponse{
		Cant:   0,
		Estado: 0,
	}

	getStosks(w, r, "")
	if len(stocks) > 0 {
		repo := NewStockRepository()
		// Usar el método más optimizado con UPSERT
		if err := repo.BatchUpsertStocks(stocks); err != nil {
			http.Error(w, "Error guardando los stocks en la base de datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response.Cant = len(stocks)
	response.Estado = http.StatusOK
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error serializando respuesta", http.StatusInternalServerError)
	}
	log.Println("Finalizandos sincronización de stocks")
}

func getStosks(w http.ResponseWriter, r *http.Request, q string) {
	externalAPI := os.Getenv("EXTERNAL_API")
	token := os.Getenv("TOKEN_EXTERNAL_API")

	if externalAPI == "" || token == "" {
		http.Error(w, "Faltan variables de entorno para la API externa", http.StatusInternalServerError)
		return
	}

	if q != "" {
		externalAPI = externalAPI + "?next_page=" + q
	}

	client := &http.Client{Timeout: 30 * time.Second}
	request, err := http.NewRequest("GET", externalAPI, nil)
	if err != nil {
		http.Error(w, "Error creando la petición: "+err.Error(), http.StatusInternalServerError)
		return
	}
	request.Header.Set("Authorization", token)

	resp, err := client.Do(request)
	if err != nil {
		http.Error(w, "Error al consumir la API externa: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Error de la API externa: %s", string(body)), http.StatusBadGateway)
		return
	}

	var apiResp StocksListResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		http.Error(w, "Error decodificando la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(apiResp.Items) == 0 {
		http.Error(w, "No se recibieron datos de la API externa", http.StatusNoContent)
		return
	}

	for _, s := range apiResp.Items {
		lastUpdated, _ := time.Parse(time.RFC3339, s.Time)
		stocks = append(stocks, Stock{
			Ticker:     s.Ticker,
			TargetFrom: s.TargetFrom,
			TargetTo:   s.TargetTo,
			Company:    s.Company,
			Action:     s.Action,
			Brokerage:  s.Brokerage,
			RatingFrom: s.RatingFrom,
			RatingTo:   s.RatingTo,
			Time:       s.Time,
			CreatedAt:  time.Now(),
			UpdatedAt:  lastUpdated,
		})
	}

	if apiResp.NextPage != "" {
		getStosks(w, r, apiResp.NextPage)
	}
}
