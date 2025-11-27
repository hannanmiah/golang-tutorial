package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/hannanmiah/golang-tutorial/config"
	"github.com/hannanmiah/golang-tutorial/handlers"
	"github.com/hannanmiah/golang-tutorial/middleware"
	"github.com/hannanmiah/golang-tutorial/models"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to E-Commerce API",
			"version": "1.0.0",
		})
	})

	userHandler := handlers.NewUserHandler(db)
	productHandler := handlers.NewProductHandler(db)
	cartHandler := handlers.NewCartHandler(db)
	orderHandler := handlers.NewOrderHandler(db)

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", userHandler.Profile)
		
		protected.GET("/products", productHandler.GetProducts)
		protected.GET("/products/:id", productHandler.GetProduct)
		protected.POST("/products", productHandler.CreateProduct)
		protected.PUT("/products/:id", productHandler.UpdateProduct)
		protected.DELETE("/products/:id", productHandler.DeleteProduct)
		protected.GET("/my-products", productHandler.GetMyProducts)
		
		protected.GET("/cart", cartHandler.GetCart)
		protected.POST("/cart", cartHandler.AddToCart)
		protected.PUT("/cart/:id", cartHandler.UpdateCartItem)
		protected.DELETE("/cart/:id", cartHandler.RemoveFromCart)
		protected.DELETE("/cart", cartHandler.ClearCart)
		
		protected.GET("/orders", orderHandler.GetOrders)
		protected.GET("/orders/:id", orderHandler.GetOrder)
		protected.POST("/orders", orderHandler.CreateOrder)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		admin.GET("/orders", orderHandler.GetAllOrders)
		admin.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
	}

	fmt.Printf("E-Commerce API Server is running on port %s\n", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}