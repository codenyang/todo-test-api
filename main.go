package main

import (
	"go-todo/router"
)

func main() {
	r := router.SetGin()

	r.Run()
}
