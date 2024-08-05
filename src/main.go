package main

import (
	"ApiRestaurant/src/database"
	"ApiRestaurant/src/routes"
	"ApiRestaurant/src/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error conenctandote a la base de datos:", err)
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
