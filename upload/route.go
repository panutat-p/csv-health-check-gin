package upload

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
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
	fmt.Println(txId)

	rawDecodedText, err := base64.StdEncoding.DecodeString(file.Data)
	if err != nil {
		fmt.Println("🟥 cannot decode base64")
		fmt.Println(err)
	}

	data, err := Convert(rawDecodedText)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func Convert(data []byte) ([]string, error) {
	buf := bytes.NewBuffer(data)
	r := csv.NewReader(buf)
	sl := make([]string, 0)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, v := range records {
		sl = append(sl, v[0])
	}
	return sl, nil
}
