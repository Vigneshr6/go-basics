package main

import (
	"github.com/gin-gonic/gin"
	"mycompany.com/go/problems"
)

func startServer() {
	r := gin.Default()
	r.POST("/count", problems.CountHandler)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
