package upload

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var dir, _ = os.Getwd()

func CreateFile(c *gin.Context) {
	var file File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	txId := c.Request.Header.Get("Transaction-ID")
	fmt.Println(txId)

	rawDecodedText, err := base64.StdEncoding.DecodeString(file.File)
	if err != nil {
		fmt.Println("🟥 cannot decode base64")
		fmt.Println(err)
	}
	d1 := rawDecodedText
	err = os.WriteFile(fmt.Sprintf("%s/%s.csv", dir, txId), d1, 0644)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
