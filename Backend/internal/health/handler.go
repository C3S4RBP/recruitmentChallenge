package health

import (
	"Backend/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Services  map[string]string `json:"services"`
	Uptime    string            `json:"uptime"`
}

var startTime = time.Now()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Services:  make(map[string]string),
		Uptime:    time.Since(startTime).String(),
	}

	if err := checkDatabase(); err != nil {
		response.Status = "unhealthy"
		response.Services["database"] = "error: " + err.Error()
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		response.Services["database"] = "connected"
	}

	envStatus := checkEnvironment()
	if envStatus != "ok" {
		response.Status = "degraded"
		response.Services["environment"] = envStatus
	} else {
		response.Services["environment"] = "ok"
	}

	statusCode := http.StatusOK
	if response.Status == "unhealthy" {
		statusCode = http.StatusServiceUnavailable
	} else if response.Status == "degraded" {
		statusCode = http.StatusOK // Degraded pero aún funcional
	}

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error serializando respuesta", http.StatusInternalServerError)
	}
}

func checkDatabase() error {
	if db.DB == nil {
		return fmt.Errorf("conexión a base de datos no inicializada")
	}

	return nil
}

func checkEnvironment() string {
	requiredEnvVars := []string{
		"SERVER_DB",
		"USER_DB",
		"PASS_DB",
		"PORT_DB",
		"EXTERNAL_API",
		"TOKEN_EXTERNAL_API",
	}

	missingVars := []string{}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			missingVars = append(missingVars, envVar)
		}
	}

	if len(missingVars) > 0 {
		return "missing variables: " + strings.Join(missingVars, ", ")
	}

	return "ok"
}
