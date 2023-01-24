package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestBlogServiceRoute(t *testing.T) {
	router := gin.Default()

	router.POST("/", Hello)
	router.Run(":8000")
}

func Hello(ctx *gin.Context) {

	data, err := ctx.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("data: %v\n", string(data))

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})

}
