# Soxy - a very fast tool for checking open SOCKS proxies in Golang 
I was looking for some open socks proxies, and so I needed to test them - but really fast. So I wrote on in Go!

### Installation
If you have a properly configured GOPATH and $GOPATH/bin is in your PATH, then run this command for a one-liner install, thank you golang!
```
go get -u github.com/pry0cc/soxy
```

### Usage
`proxies.txt`
```
8.8.8.8:3128
8.8.8.8:8080
```

```
cat proxies.txt | soxy | tee alive.txt
```

### Credit
I pulled the proxy checking code and some of the multi-threading out of https://github.com/trigun117/ProxyChecker, so credit to trigun117! 
