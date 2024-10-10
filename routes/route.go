// routes/routes.go
package routes

import (
	"BasicTrade/controllers"
	middleware "BasicTrade/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", controllers.RegisterAdmin)
		authGroup.POST("/login", controllers.LoginAdmin)
	}

	productGroup := r.Group("/products")
	productGroup.Use(middleware.AuthMiddleware())
	{
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.PUT("/:productUUID", controllers.UpdateProduct)
		productGroup.DELETE("/:productUUID", controllers.DeleteProduct)
		productGroup.GET("/:productUUID", controllers.GetProductByUUID)
		productGroup.GET("/", controllers.GetAllProducts)

		productGroup.POST("/variants", controllers.CreateVariant)
		productGroup.PUT("/variants/:variantUUID", controllers.UpdateVariant)
		productGroup.DELETE("/variants/:variantUUID", controllers.DeleteVariant)
		productGroup.GET("/variants/:variantUUID", controllers.GetVariantByUUID)
		productGroup.GET("/variants", controllers.GetAllVariants)
	}

	return r
}
