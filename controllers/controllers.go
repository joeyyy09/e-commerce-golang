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

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
	 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 defer cancel()

	 var user models.User
	 if err := c.BindJSON(&user); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	 }

	 validationError:= Validate.Struct(user)
	 if validationError != nil {	
		 c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
		 return
	 }

	 count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	 if err != nil {
		log.Panic(err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 return
	 }

	 if count > 0 {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		 return
	 }

	 count, err = UserCollection.CountDocuments(ctx, bson.M{"Phone ": user.Phone})

	 defer cancel()

	 if err != nil {
		log.Panic(err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 return
	 }

	 if count > 0 {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Phone already exists"})
		 return
	 }

	 


	}
}

func Login() gin.HandlerFunc {
}

func ProductViewerAdmin() gin.HandlerFunc {
}

func SearchProduct() gin.HandlerFunc {
}

func SearchProductbyQuery() gin.HandlerFunc {
}