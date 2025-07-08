package db

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	connString := os.Getenv("SERVER_DB")
	environment := os.Getenv("ENVIRONMENT")
	var config *gorm.Config
	if environment == "PRD" {
		config = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		}
	} else {
		config = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	database, err := gorm.Open(postgres.Open(fmt.Sprintf(connString, os.Getenv("USER_DB"), os.Getenv("PASS_DB"), os.Getenv("PORT_DB"))), config)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	log.Println("✅ Base de datos conectada correctamente")
	DB = database
}

func MigrateModels(models ...interface{}) {
	if DB == nil {
		log.Fatal("La base de datos no está inicializada")
	}

	for _, model := range models {
		log.Println("sincronizando tablas ", reflect.TypeOf(model))
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatal("Error migrando modelo ", reflect.TypeOf(model), ": ", err)
		}
	}

	log.Println("✅ Migración de tablas completada")
}

// RunOptimizationMigrations ejecuta migraciones SQL para optimizar el rendimiento
func RunOptimizationMigrations() {
	if DB == nil {
		log.Fatal("La base de datos no está inicializada")
	}

	// Leer y ejecutar el archivo de migraciones
	migrationSQL := `
	-- Índice para consultas por empresa
	CREATE INDEX IF NOT EXISTS idx_stocks_company ON stocks(company);

	-- Índice para consultas por broker
	CREATE INDEX IF NOT EXISTS idx_stocks_brokerage ON stocks(brokerage);

	-- Índice para consultas por acción
	CREATE INDEX IF NOT EXISTS idx_stocks_action ON stocks(action);

	-- Índice compuesto para consultas frecuentes
	CREATE INDEX IF NOT EXISTS idx_stocks_company_brokerage ON stocks(company, brokerage);

	-- Índice para consultas por fecha de actualización
	CREATE INDEX IF NOT EXISTS idx_stocks_updated_at ON stocks(updated_at);

	-- Índice para consultas por fecha de creación
	CREATE INDEX IF NOT EXISTS idx_stocks_created_at ON stocks(created_at);
	`

	// Ejecutar las migraciones
	if err := DB.Exec(migrationSQL).Error; err != nil {
		log.Printf("Error ejecutando migraciones de optimización: %v", err)
		return
	}

	log.Println("✅ Migraciones de optimización completadas")
}
