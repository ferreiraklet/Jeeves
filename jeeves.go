package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	flag.Usage = func() {
		help := []string{
			"",
			"",
			"Usage:",
			"+=======================================================+",
			"       --payload-time,            The time from payload",
			"       -h                   Show This Help Message",
			"",
			"+=======================================================+",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
	}

}

func main() {


	var payloadTime int
	flag.IntVar(&payloadTime, "payload-time", 0,"")
	flag.Parse()
	
	var urls []string
	std := bufio.NewScanner(os.Stdin)
	for std.Scan() {
		var line string = std.Text()
		hline := strings.Replace(line, "%2F", "/", -1)
		line = hline

		urls = append(urls, line)

	}
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {

			defer wg.Done()

			x := getParams(url, payloadTime)
			if x != "ERROR" {
				fmt.Println(x)
			}

		}(u)
	}

	wg.Wait()

}

func getParams(url string, pTime int) string {

	var trans = &http.Transport{
		MaxIdleConns:      30,
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			KeepAlive: time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: trans,
	}

	before := time.Now().Second()
	res, err := http.NewRequest("GET", url, nil)
	res.Header.Set("Connection", "close")
	resp, err := client.Do(res)
	// res.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")
	after := time.Now().Second()
	
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	if err != nil {
		return "ERROR"
	}


	if (after - before) >= pTime{
		return "\033[1;31mVulnerable To Time-Based SQLI "+url+"\033[0;0m"
	}else{
		return "\033[1;30mNot Vulnerable to SQLI Time-Based "+url+"\033[0;0m"
	}


}

