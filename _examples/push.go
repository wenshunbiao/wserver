package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"wserver/conf"
	"wserver/server"
)

func main() {
	pushURL := "http://127.0.0.1:12345/push"
	contentType := "application/json"
	Authorization := conf.AppConf.Server.PushAuthKey

	for {
		pm := server.PushMessage{
			UserID:  "all",
			Event:   "topic1",
			Message: fmt.Sprintf("Hello in %s", time.Now().Format("2006-01-02 15:04:05.000")),
		}
		b, _ := json.Marshal(pm)
		fmt.Println(string(b))

		client := &http.Client{}
		req, err := http.NewRequest("POST", pushURL, bytes.NewReader(b))
		if err != nil {
			panic(err.Error())
		}
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Authorization", Authorization)
		response, err2 := client.Do(req)
		defer response.Body.Close()
		if err2 != nil {
			panic(err2.Error())
		}

		if response.StatusCode != 200 {
			fmt.Println(response.Status)
			return
		}

		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		time.Sleep(time.Second)
	}
}
