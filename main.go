package main

import (
	"go-todo/router"
)

func main() {
	r := router.SetGin()

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
