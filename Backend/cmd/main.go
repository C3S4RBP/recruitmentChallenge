package main

import (
	"Backend/Router"
	"Backend/db"
	"Backend/internal/API_EXTERNAL"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//cargando archivo de variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env ", err)
	}

	//iniciaclizando la BD
	db.InitDB()

	// Migrar las tablas automáticamente
	db.MigrateModels(&API_EXTERNAL.Stock{})

	// Ejecutar migraciones de optimización
	db.RunOptimizationMigrations()

	// Configurar rutas HTTP
	Router.SetupRoutes()

	port := os.Getenv("PORT_SERVER")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor iniciado en puerto %s\n", port)

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
