package utils

import (
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

// LocalIPByTCP 有外网的情况下, 通过tcp访问获得本机 IP 地址
func LocalIPByTCP() string {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

// LocalIP get local ip
func LocalIP() (ipAddr string) {
	return LocalIPWithVal("")
}

// LocalIPWithVal if local IP is "", return defaultVal
func LocalIPWithVal(defaultVal string) string {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		return defaultVal
	}
	for _, addr := range addrSlice {
		if IPNet, ok := addr.(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if nil != IPNet.IP.To4() {
				return IPNet.IP.String()
			}
		}
	}
	return defaultVal
}

// PublicIP get public ip
func PublicIP() (string, error) {
	timeout := time.Duration(10)
	conn, err := net.DialTimeout("tcp", "ns1.dnspod.net:6666", timeout*time.Second)
	defer func() {
		if x := recover(); x != nil {
			log.Println("Can't get public ip", x)
		}
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil {
		var bytes []byte
		deadline := time.Now().Add(timeout * time.Second)
		err = conn.SetDeadline(deadline)
		if err != nil {
			return "", err
		}
		bytes, err = ioutil.ReadAll(conn)
		if err == nil {
			return string(bytes), nil
		}
	}
	return "", err
}
