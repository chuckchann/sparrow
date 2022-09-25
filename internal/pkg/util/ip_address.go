package util

import (
	"io/ioutil"
	"net"
	"net/http"
)

const defaultIp = "127.0.0.1"

func InternalIp() string {
	//get all net device
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}

	return defaultIp
}

func ExternalIp() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(ip)
}
