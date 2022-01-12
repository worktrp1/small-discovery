package smalldiscovery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func GetLANIP() []byte {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)

	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("Got LAN IP: %s\n", ipAddress.IP)

	return ipAddress.IP
}

func GetWANIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IPQuery
	json.Unmarshal(body, &ip)

	fmt.Printf("Got Query: %s\n", ip.Query)

	return ip.Query
}

func GetOSHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Failed to print hostname: %s\n", err)
	}

	return hostname
}
