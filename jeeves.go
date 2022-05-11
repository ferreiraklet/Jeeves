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
        "strconv"
        "time"
)

func init() {
        flag.Usage = func() {
                help := []string{
                     "",
                     "",
                     "Usage:",
                     "+=======================================================+",
                     "       -t, --payload-time,  The time from payload",
                     "       -c                   Set Concurrency, Default: 25",
                     "       -p, --proxy          Send traffic to a proxy",
                     "       -H, --headers        Custom Headers",
                     "       -d, --data           Sending Post request with data",
                     "       -h                   Show This Help Message",
                     "",
                     "+=======================================================+",
                     "",
                }
                fmt.Println(`

         /\/\
        /  \ \
       / /\ \ \
       \/ /\/ /
       / /\/ /\
      / /\ \/\ \
     / / /\ \ \ \
  /\/ / / /\ \ \ \/\
 /  \/ / /  \ \ \ \ \
/ /\ \/ /    \ \/\ \ \
\/ /\/ /      \/ /\/ /
/ /\/ /\      / /\/ /\
\ \ \/\ \    / /\ \/ /
 \ \ \ \ \  / / /\  /
  \/\ \ \ \/ / / /\/
     \ \ \ \/ / /
      \ \/\ \/ /
       \/ /\/ /
       / /\/ /\
       \ \ \/ /
        \ \  /
         \/\/

`)
                fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
        }

}

func main() {


        var concurrency int
        flag.IntVar(&concurrency, "c", 25, "")

        var payloadTime int
        flag.IntVar(&payloadTime, "payload-time", 0,"")
        flag.IntVar(&payloadTime, "t", 0, "")

        var proxy string
        flag.StringVar(&proxy,"proxy","","")
        flag.StringVar(&proxy, "p","","")

        var headers string
        flag.StringVar(&headers,"headers","","")
        flag.StringVar(&headers,"H","","")

        var data string
        flag.StringVar(&data, "d", "","")
        flag.StringVar(&headers, "data", "","")

        flag.Parse()

        std := bufio.NewScanner(os.Stdin)

        //buf
        alvos := make(chan string)
        var wg sync.WaitGroup

        for i:=0;i<concurrency;i++ {
                wg.Add(1)
                go func() {

                     defer wg.Done()
                     for alvo := range alvos{
                        
                        if !strings.HasPrefix(alvo, "http"){
                            continue
                        }
                        _, err := url.Parse(alvo)
                        if err != nil{
                            continue
                        }
                        if data != ""{
                            bodyReq(alvo, payloadTime, proxy, headers, data)
                        
                        }else{
                            x := jeeves(alvo, payloadTime, proxy, headers)
                            if x != "ERROR" {
                                fmt.Println(x)
                                        }
                        }
                    }

                }()
        }

        for std.Scan() {
            var line string = std.Text()
            alvos <- line

        }
        close(alvos)
        wg.Wait()

}

func jeeves(turl string, pTime int, proxy string, headers string) string {

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
                CheckRedirect: func(req *http.Request, via []*http.Request) error {
                    return http.ErrUseLastResponse
                    },
        }

        
        if proxy != "" {
            if p, err := url.Parse(proxy); err == nil {
                trans.Proxy = http.ProxyURL(p)
        }

        }
        before := time.Now().Second()
        res, err := http.NewRequest("GET", turl, nil)

        if headers != ""{
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

        //res.Header.Set("Connection", "close")
        resp, err := client.Do(res)
        if resp.StatusCode >= 300{
            scstring := strconv.Itoa(resp.StatusCode)
            return "\033[1;30mNeed Manual Analisys "+scstring+" - "+turl+"\033[0;0m"}
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


func bodyReq(turl string, pTime int, proxy string, headers string, data string) string {

        var trans = &http.Transport{
                MaxIdleConns:      30,
                //IdleConnTimeout:   time.Second,
                DisableKeepAlives: true,
                TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
                DialContext: (&net.Dialer{
                     KeepAlive: time.Second,
                }).DialContext,
        }

        client := &http.Client{
                Transport: trans,
        }

        if proxy != "" {
            if p, err := url.Parse(proxy); err == nil {
                trans.Proxy = http.ProxyURL(p)
        }

        }


        requestBody := url.Values{}

        postparts := strings.Split(data, "&")
        for _, q := range postparts{
            sep := strings.Split(q, "=")
            //fmt.Println(sep)
            name := sep[0]
            value := sep[1]
            requestBody.Set(name, value)
        }
    
        qwe := requestBody.Encode()

        decodedUrl, err := url.QueryUnescape(string(qwe))
        if err != nil{
            return "ERROR"
        }

        before := time.Now().Second()
        res, err := http.NewRequest("POST", turl, strings.NewReader(decodedUrl))



        if headers != ""{
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

        resp, err := client.Do(res)

        if resp.StatusCode >= 300{
            scstring := strconv.Itoa(resp.StatusCode)
            return "\033[1;30mNeed Manual Analisys "+scstring+" - "+turl+"\033[0;0m"}
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
