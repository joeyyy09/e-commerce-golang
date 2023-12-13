package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joeyyy09/e-commerce-golang/models"
	"github.com/joeyyy09/e-commerce-golang/utils"
	"net/http"
	"strconv"
)

func HashPassword(password string) string {
}

func VerifyPassword(userPassword string, givenPassword string) (bool,string) {
}

func Login() gin.HandlerFunc {
}

func ProductViewerAdmin() gin.HandlerFunc {
}

func SearchProduct() gin.HandlerFunc {
}

func SearchProductbyQuery() gin.HandlerFunc {
}