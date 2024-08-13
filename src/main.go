// main.go
package main

import (
	"ApiRestaurant/src/database"
	"ApiRestaurant/src/routes"
	"ApiRestaurant/src/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @security Bearer
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
