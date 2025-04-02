package main

import (
	"series-tracker/pkg/db"
	"series-tracker/pkg/handlers"
	"series-tracker/pkg/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Series Tracker API
// @version 1.0
// @description API para gestionar series de TV
// @host localhost:8080
// @BasePath /api
func main() {
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// Conectar a la base de datos
	database, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&models.Series{})
	handler := &handlers.Handler{DB: database}

	// Rutas de la API
	api := router.Group("/api")
	{
		api.GET("/series", handler.GetAllSeries)
		api.POST("/series", handler.CreateSeries)
		api.GET("/series/:id", handler.GetSeriesByID)
		api.PUT("/series/:id", handler.UpdateSeries)
		api.DELETE("/series/:id", handler.DeleteSeries)
		api.PATCH("/series/:id/status", handler.UpdateStatus)
		api.PATCH("/series/:id/episode", handler.IncrementEpisode)
		api.PATCH("/series/:id/upvote", handler.Upvote)
		api.PATCH("/series/:id/downvote", handler.Downvote)
	}

	// Ruta para la documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar el servidor
	router.Run(":8080")
}
