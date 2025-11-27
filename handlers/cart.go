package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/hannanmiah/golang-tutorial/models"
)

type CartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{db: db}
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var cartItems []models.Cart
	if err := h.db.Where("user_id = ?", userID).
		Preload("Product").
		Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart_items": cartItems})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := h.db.First(&product, req.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Stock < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient stock for product " + product.Name,
		})
		return
	}

	var existingCart models.Cart
	if err := h.db.Where("user_id = ? AND product_id = ?", userID, req.ProductID).
		First(&existingCart).Error; err == nil {
		
		newQuantity := existingCart.Quantity + req.Quantity
		if newQuantity > product.Stock {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Insufficient stock for product " + product.Name,
			})
			return
		}

		if err := h.db.Model(&existingCart).Update("quantity", newQuantity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
			return
		}

		h.db.Preload("Product").First(&existingCart, existingCart.ID)
		c.JSON(http.StatusOK, gin.H{
			"message":  "Cart item updated successfully",
			"cart_item": existingCart,
		})
		return
	}

	cartItem := models.Cart{
		UserID:    userID.(uint),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.db.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	h.db.Preload("Product").First(&cartItem, cartItem.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Item added to cart successfully",
		"cart_item": cartItem,
	})
}

func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	var cartItem models.Cart

	if err := h.db.Where("id = ? AND user_id = ?", id, userID).
		First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	var req UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := h.db.First(&product, cartItem.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if req.Quantity > product.Stock {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Insufficient stock for product " + product.Name,
		})
		return
	}

	if err := h.db.Model(&cartItem).Update("quantity", req.Quantity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	h.db.Preload("Product").First(&cartItem, id)
	c.JSON(http.StatusOK, gin.H{
		"message":   "Cart item updated successfully",
		"cart_item": cartItem,
	})
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	var cartItem models.Cart

	if err := h.db.Where("id = ? AND user_id = ?", id, userID).
		First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	if err := h.db.Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := h.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
}