package hardware_inspection

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"code.cloudfoundry.org/bytefmt"
)

func getMemInfo(ipAddress string) uint64 {
	var (
		err    error
		memory uint64 = 0
		resp   *http.Response
		body   []byte
		url    string = "http://" + ipAddress + "/meminfo.txt"
	)
	resp, err = httpClient.Get(url)
	if err != nil { // If request didn't work
		if debug {
			log.Println("INF Couldn't access '" + url + "'.")
		}
		return memory // return empty size, since request failed
	}
	// Read full content of response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ERR Can't read response body;", err)
	}
	resp.Body.Close()

	// body:
	// MemTotal:        2030280 kB
	// ... (other lines) ...
	lines := strings.Split(string(body), "\n")
	line := lines[0]
	line = strings.ReplaceAll(line, " ", "")
	memory, err = bytefmt.ToBytes(strings.Split(line, ":")[1])
	if err != nil {
		log.Fatalln("ERR Can't parse as tosca.size '"+strings.Split(line, ":")[1]+"';", err)
	}

	if debug {
		log.Println("INF Found memory size for ip", ipAddress)
	}

	return memory
}
