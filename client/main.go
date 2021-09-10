/**
 * Auth :   liubo
 * Date :   2021/9/10 15:28
 * Comment:
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:20210", "127.0.0.1:20210")

func req(action string) {
	var url = "http://" + *addr + "/" + action
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var resp, e = client.Get(url)
	if e == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			var d, _ = ioutil.ReadAll(resp.Body)
			fmt.Println(action, " resp:", string(d))
		} else {
			fmt.Println(action, " resp, err=", resp.StatusCode)
		}
	} else {
		fmt.Println(url, " req failed")
	}
}

func main() {
	flag.Parse()

	var wg = sync.WaitGroup{}

	var echoAction = func(s string) {
		req("api/" + s + "/echo?msg=" + s)
		wg.Done()
	}

	var cnt = 0
	for {
		cnt++
		fmt.Println("---------- " + strconv.Itoa(cnt)  + " ----------")

		wg.Add(3)
		echoAction("s1")
		echoAction("s2")
		echoAction("s3")
		wg.Wait()

		time.Sleep(time.Second * 2)
	}
}