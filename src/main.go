package main

import (
	"ApiRestaurant/src/database"
	"ApiRestaurant/src/routes"
	"ApiRestaurant/src/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	var db *sql.DB
	var err error
	maxRetries := 10
	retryInterval := 5 * time.Second
	for i := 0; i < maxRetries; i++ {
		db, err = database.Connect()
		if err != nil {
			log.Printf("Error connecting to database: %v", err)
			time.Sleep(retryInterval)
			continue
		}
		if err = db.Ping(); err != nil {
			log.Printf("Error pinging database: %v. Retrying in %v...", err, retryInterval)
			db.Close() // Cierra la conexión actual si el ping falla
			time.Sleep(retryInterval)
			continue
		}
		log.Println("Successfully connected to the database!")
		break
	}
	defer db.Close()
	hamburguesaService := services.NewHamburguesaService(db)
	pedidosService := services.NewPedidoServices(db)
	userService := services.NewUserServices(db)
	r := gin.Default()

	routes.SetupHamburguesaRoutes(r, hamburguesaService, userService)
	routes.SetupPedidoRoutes(r, pedidosService, userService)
	routes.SetupUserRoutes(r, userService)

	r.Run()

}
