package main

import "github.com/gin-gonic/gin"

func startServer() {
	r := gin.Default()
	r.POST("/count", countHandler)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
