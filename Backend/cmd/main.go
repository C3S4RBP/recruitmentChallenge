package main

import (
	"Backend/db"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hola, este es el api")

	//cargando archivo de variables de entorno
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error cargando archivo .env ", err)
    }
	//iniciaclizando la BD
	db.InitDB()
}



    

  