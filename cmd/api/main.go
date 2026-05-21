package main

import (
	"log"
	"os"

	"github.com/GN15H/paytrack/internal/accounts"
	"github.com/GN15H/paytrack/internal/auth"
	"github.com/GN15H/paytrack/internal/categories"
	"github.com/GN15H/paytrack/internal/db"
	"github.com/GN15H/paytrack/internal/summary"
	"github.com/GN15H/paytrack/internal/transactions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using environment variables")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	authService := auth.NewService(database)
	authHandler := auth.NewHandler(authService)

	accountsService := accounts.NewService(database)
	accountsHandler := accounts.NewHandler(accountsService)

	categoriesService := categories.NewService(database)
	categoriesHandler := categories.NewHandler(categoriesService)

	transactionsService := transactions.NewService(database)
	transactionsHandler := transactions.NewHandler(transactionsService)

	summaryService := summary.NewService(database)
	summaryHandler := summary.NewHandler(summaryService)

	r := gin.Default()

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(auth.Middleware())
	{
		protected.GET("/accounts", accountsHandler.FindAll)
		protected.POST("/accounts", accountsHandler.Create)
		protected.GET("/accounts/:id", accountsHandler.FindByID)
		protected.DELETE("/accounts/:id", accountsHandler.Delete)

		protected.GET("/accounts/:id/transactions", transactionsHandler.FindAll)
		protected.POST("/accounts/:id/transactions", transactionsHandler.Create)
		protected.DELETE("/accounts/:id/transactions/:txID", transactionsHandler.Delete)

		protected.POST("/accounts/transfer", transactionsHandler.Transfer)

		protected.GET("/categories", categoriesHandler.FindAll)
		protected.POST("/categories", categoriesHandler.Create)

		protected.GET("/summary", summaryHandler.GetSummary)
		protected.GET("/summary/by-category", summaryHandler.GetByCategory)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
