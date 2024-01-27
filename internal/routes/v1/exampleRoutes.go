package routes

import (
	"github.com/Dune-Global/unite-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func ExampleRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.ExampleController)
}