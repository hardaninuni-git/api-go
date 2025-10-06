package routes

import (
	"net/http" // 👈 tambahkan ini
	"rhiona-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// === Tambahkan middleware CORS ===
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	// =================================

	api := r.Group("/api")
	{
		api.POST("/customers", controllers.CreateCustomer)
		api.GET("/customers", controllers.GetCustomers)
		api.GET("/customers/:id", controllers.GetCustomerByID)
		api.PUT("/customers/:id", controllers.UpdateCustomer)
		api.DELETE("/customers/:id", controllers.DeleteCustomer)
	}

	return r
}
