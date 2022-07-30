package upload

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/panutat-p/line-remoteinterview-gin/color"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const TIME_OUT = 3 * time.Second

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

func Check(ch chan bool, url string) {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("%s%s", "http://", url)
	}
	c := &http.Client{
		Timeout: TIME_OUT,
	}
	_, err := c.Get(url)
	if err != nil {
		if os.IsTimeout(err) {
			ch <- false
			fmt.Println(color.Gray(err.Error()))
			return
		}

		ch <- false
		fmt.Println(color.Gray(err.Error()))
		return
	}
	ch <- true
}
