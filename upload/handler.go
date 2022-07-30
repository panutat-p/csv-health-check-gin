package upload

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/panutat-p/line-remoteinterview-gin/color"
	"log"
	"net/http"
	"time"
)

func Handler(c *gin.Context) {
	start := time.Now()

	txId := c.Request.Header.Get("Transaction-ID")

	var file File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println(color.Blue(txId), "🟨 Bad request:", err)
		return
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(file.Data)
	if err != nil {
		log.Println(color.Blue(txId), "🟧 Failed to decode base64:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := Convert(rawDecodedText)
	if err != nil {
		log.Println(color.Blue(txId), "🟧 Failed to convert CSV data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ch := make(chan bool)
	for _, v := range data {
		go Check(ch, v)
	}

	var up int
	var down int
	for i := 0; i < len(data); i++ {
		r := <-ch
		if r {
			up += 1
		} else {
			down += 1
		}
	}

	elapsed := time.Since(start)

	p := ResponsePayload{
		Total:            len(data),
		Up:               up,
		Down:             down,
		ElapsedTimeMilli: int(elapsed.Milliseconds()),
	}
	c.JSON(http.StatusOK, &p)
	log.Println(color.Blue(txId), "🟩 success, elapsed time:", elapsed)
}
