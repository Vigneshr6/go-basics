package problems

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mycompany.com/go/algorithms"
)

// most used words and times of occurrence in the text
// input:- I am manish manish is good person manish like cricket player manish like music manish is now holyboll player// output// [// {// "word" : "manish",// "count":5// },// {// "word" : "like",// "count":2// }

type tekionRequest struct {
	Input string `json:"input"`
}

type tekionResponse struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func CountHandler(ctx *gin.Context) {
	var ij tekionRequest
	err := ctx.BindJSON(&ij)
	if err != nil {
		panic(err)
	}
	keys, countMap := algorithms.GetWordCount(ij.Input)
	fmt.Println(keys)
	var response []tekionResponse
	for i, k := range keys {
		if i > 9 {
			break
		}
		response = append(response, tekionResponse{k, countMap[k]})
	}
	ctx.JSON(http.StatusOK, response)
}
