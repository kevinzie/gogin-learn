package mappings

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/s12i/gin-throttle"
	"ramo/bookstore/controllers"
)

var Router *gin.Engine

func CreateUrlMappings() {
	r := gin.Default()
	maxEventsPerSec := 1000
	maxBurstSize := 20

	r.Use(middleware.Throttle(maxEventsPerSec, maxBurstSize))

	// Routes
	api := r.Group("api")
	{
		books := api.Group("books")
		{
			books.GET("", controllers.FindBooks)
			books.POST("", controllers.CreateBook)

			books.GET(":id", controllers.FindBook)
			books.PATCH(":id", controllers.UpdateBook)
			books.DELETE(":id", controllers.DeleteBook)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})
	r.Run()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	// Run the server
}
