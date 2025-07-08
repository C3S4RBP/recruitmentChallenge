package Router

import (
	"Backend/internal/API_EXTERNAL"
	"Backend/internal/health"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/sync-stocks", handleMethod(http.MethodPost, API_EXTERNAL.SyncStocksHandler))
	http.HandleFunc("/stocks", handleMethod(http.MethodGet, API_EXTERNAL.GetStocksHandler))
	http.HandleFunc("/health", handleMethod(http.MethodGet, health.HealthHandler))
}

func handleMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
