// This package is used to test the Loggly package
package main

import (
	"fmt"
	loggly "github.com/jamespearly/loggly"
	"net/http"
	"bufio"
)

func main() {

	var tag string
	tag = "My-Go-Demo"

	client := loggly.New(tag)

	resp, err := http.Get("https://developer.nps.gov/api/v1/activities/parks?parkCode=acad&api_key=HABe5chvX4Zswp47KAZ5WvDriLUtSXdmqm88dEpA")

	scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 1000; i++ {
        fmt.Println(scanner.Text())
    }

	if err != nil {
		err := client.EchoSend("error", "err")
		fmt.Println("err:", err)
    }
	if resp != nil{
		resp := client.EchoSend("info", "resp")
		fmt.Println("resp:", resp)
	}

}