package upload

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Handler(c *gin.Context) {
	var file File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	txId := c.Request.Header.Get("Transaction-ID")
	log.Println("Transaction ID:", txId)

	rawDecodedText, err := base64.StdEncoding.DecodeString(file.Data)
	if err != nil {
		fmt.Println("🟥 cannot decode base64")
		fmt.Println(err)
	}

	data, err := Convert(rawDecodedText)
	if err != nil {
		log.Println(err)
	}

	ch := make(chan Result)
	for _, v := range data {
		go Check(ch, v)
	}

	result := make([]Result, 0)
	for i := 0; i < len(data); i++ {
		r := <-ch
		result = append(result, r)
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": result,
	})
}
