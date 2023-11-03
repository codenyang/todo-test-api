package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	todo2 "go-todo/todo"
	"time"
)

func SetGin() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())
	setRoute(r)

	return r
}

func setRoute(r *gin.Engine) {

	r.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		t := todo2.GetTodos()
		c.JSON(200, t)
	})

	r.POST("/todo", func(c *gin.Context) {
		var todo todo2.ToDo
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "fail",
			})
			fmt.Println(err)
			return
		}
		todo2.AddTodo(todo)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.DELETE("/todo", func(c *gin.Context) {
		var todo todo2.ToDo
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "fail",
			})
			fmt.Println(err)
			return
		}
		res := todo2.DelTodo(todo)
		if res {
			c.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			c.JSON(404, gin.H{
				"message": "not found",
			})
		}

	})

	r.PUT("/todo", func(c *gin.Context) {
		var todo todo2.ToDo
		err := c.BindJSON(&todo)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "fail",
			})
			fmt.Println(err)
			return
		}
		res := todo2.CheckTodo(todo)
		if res {
			c.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "fail",
			})
		}

	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
