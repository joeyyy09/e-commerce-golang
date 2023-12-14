package controllers

import (
)

var(
	ErrCantFindProduct = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't decode product")
	ErrUserIdNotValid = errors.New("User id not valid")
	ErrCantUpdateUser = errors.New("Can't update user")
	ErrCantRemoveCartItem = errors.New("Can't remove cart item")
	ErrCantGetItem = errors.New("Can't get item")
	ErrCantBuyCartItem = errors.New("Can't buy cart item")
)

func AddProductToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "AddProductToCart",
		})
	}
}

func RemoveCartItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "RemoveCartItem",
		})
	}
}

func BuyItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "BuyItemFromCart",
		})
	}
}

func InstantBuyer() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "InstantBuyer",
		})
	}
}