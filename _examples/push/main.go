package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"wserver/server"
)

func main() {
	pushURL := "http://127.0.0.1:12345/push"
	contentType := "application/json"

	for {
		pm := server.PushMessage{
			UserID:  "all",
			Event:   "topic1",
			Message: fmt.Sprintf("Hello in %s", time.Now().Format("2006-01-02 15:04:05.000")),
		}
		b, _ := json.Marshal(pm)

		response, err := http.DefaultClient.Post(pushURL, contentType, bytes.NewReader(b))
		if err != nil {
			fmt.Println(err)
			return
		}
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		time.Sleep(time.Second)
	}
}
