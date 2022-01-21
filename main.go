package main

import (
	"gin-postgresql/controllers" // new
	"gin-postgresql/models"                  // new

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
		r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	

	db := models.SetupModels() // new

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)

	r.POST("/books", controllers.CreateBook) // create

	r.GET("/books/:id", controllers.FindBook) // find by id

	r.PATCH("/books/:id", controllers.UpdateBook) // update by id

	r.DELETE("/books/:id", controllers.DeleteBook) // delete by id

	r.Run()
}
