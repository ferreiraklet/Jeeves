package main

import (
        "bufio"
        "crypto/tls"
        "flag"
        "fmt"
        "net"
        "net/http"
        "net/url"
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
                     "       --payload-time,      The time from payload",
                     "       --proxy              Send traffic to a proxy",
                     "       -H, --headers	  Custom Headers",
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

        var proxy string
        flag.StringVar(&proxy,"proxy","","")

        var headers string
        flag.StringVar(&headers,"headers","","")
        flag.StringVar(&headers,"H","","")

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

                     if proxy != ""{
                            if headers != ""{
                                x := getParams(url, payloadTime, proxy, headers)
                                if x != "ERROR" {
                                    fmt.Println(x)
                                                }
                            }else{
                                x := getParams(url, payloadTime, proxy, "0")
                                if x != "ERROR" {
                                    fmt.Println(x)
                            }
                            }
                    }else{
                                if headers != ""{
                                    x := getParams(url, payloadTime, "0", headers)
                                    if x != "ERROR" {
                                        fmt.Println(x)
                                                    }
                                }else{
                                        x := getParams(url, payloadTime, "0", "0")
                                        if x != "ERROR" {
                                            fmt.Println(x)
                                                        }
                                    }

                            }

                }(u)
        }

        wg.Wait()

}

func getParams(turl string, pTime int, proxy string, headers string) string {

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
	
	_, err := url.ParseRequestURI(turl)
        if err != nil{
                return "ERROR"
        }
        if proxy != "0" {
            if p, err := url.Parse(proxy); err == nil {
                trans.Proxy = http.ProxyURL(p)
        }

        }
        before := time.Now().Second()
        res, err := http.NewRequest("GET", turl, nil)

        if headers != "0"{
            if strings.Contains(headers, ";"){
                    parts := strings.Split(headers, ";")
                    for _, q := range parts{
                        separatedHeader := strings.Split(q,":")
                        res.Header.Set(separatedHeader[0],separatedHeader[1])
        }
}else{
    sHeader := strings.Split(headers,":")
    res.Header.Set(sHeader[0], sHeader[1])
}
}

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
                return "\033[1;31mVulnerable To Time-Based SQLI "+turl+"\033[0;0m"
        }else{
                return "\033[1;30mNot Vulnerable to SQLI Time-Based "+turl+"\033[0;0m"
        }


}
