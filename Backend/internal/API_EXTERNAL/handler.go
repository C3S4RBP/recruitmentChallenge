package API_EXTERNAL

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := 1
	pageSize := 10
	company := ""

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	if companyParam := r.URL.Query().Get("company"); companyParam != "" {
		company = companyParam
	}

	offset := (page - 1) * pageSize

	repo := NewStockRepository()

	var stocks []Stock
	var total int64
	var err error

	if company != "" {
		stocks, err = repo.GetStocksByTicker(company, pageSize, offset)
		if err != nil {
			http.Error(w, "Error consultando stocks por company: "+err.Error(), http.StatusInternalServerError)
			return
		}
		total, err = repo.GetStocksCountByTicker(company)
		if err != nil {
			http.Error(w, "Error obteniendo conteo de stocks: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		stocks, err = repo.GetAllStocks(pageSize, offset)
		if err != nil {
			http.Error(w, "Error consultando stocks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		total, err = repo.GetTotalStocksCount()
		if err != nil {
			http.Error(w, "Error obteniendo conteo total de stocks: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	response := PaginatedStocksResponse{
		Stocks:     stocks,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error serializando respuesta", http.StatusInternalServerError)
		return
	}
}
