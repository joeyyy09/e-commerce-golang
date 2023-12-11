package main

import(
	"github.com/joeyyy09/e-commerce-golang/routes"
	"github.com/joeyyy09/e-commerce-golang/controllers"
	"github.com/joeyyy09/e-commerce-golang/database"
	"github.com/joeyyy09/e-commerce-golang/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	port:= os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app:= controllers.newApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client, "Users"))

	router : gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}