package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	timeout = time.Duration(5 * time.Second)
)

// IP is a struct for storing the JSON output from apify
type IP struct {
	IP string
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(50)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		proxy := strings.ToLower(sc.Text())
		wg.Add(1)
		go CheckProxySOCKS(proxy, &wg)
	}

	wg.Wait()
}

//CheckProxySOCKS Check proxies on valid
func CheckProxySOCKS(proxyy string, wg *sync.WaitGroup) (err error) {

	defer wg.Done()

	d := net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}

	//Sending request through proxy
	dialer, _ := proxy.SOCKS5("tcp", proxyy, nil, &d)
	var ip IP
	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial:              dialer.Dial,
		},
	}

	response, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	json.Unmarshal([]byte(body), &ip)
	sp := strings.Split(proxyy, ":")
	respIp := sp[0]
	port := sp[1]

	if ip.IP == respIp {
		fmt.Printf("%s:%s\n", respIp, port)
	}
	defer response.Body.Close()
	return nil
}
