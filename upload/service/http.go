package service

import (
	"fmt"
	"github.com/panutat-p/line-remoteinterview-gin/logging"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const TimeOut = 3 * time.Second

func Check(ch chan bool, url string) {
	if !strings.HasPrefix(url, "http") {
		url = fmt.Sprintf("%s%s", "http://", url)
	}
	c := &http.Client{
		Timeout: TimeOut,
	}
	r, err := c.Get(url)
	if err != nil {
		if os.IsTimeout(err) {
			ch <- false
			fmt.Println(logging.Gray(err.Error()))
			return
		}

		ch <- false
		fmt.Println(logging.Gray(err.Error()))
		return
	}
	fmt.Println(r.StatusCode)
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))
	ch <- true
}
