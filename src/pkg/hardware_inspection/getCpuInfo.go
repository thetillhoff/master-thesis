package hardware_inspection

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getCpuInfo(ipAddress string) int {
	var (
		err   error
		cores int = 0
		resp  *http.Response
		body  []byte
		url   string = "http://" + ipAddress + "/cpuinfo.txt"
	)
	resp, err = httpClient.Get(url)
	if err != nil { // If request didn't work
		if debug {
			log.Println("INF Couldn't access '" + url + "'.")
		}
		return cores // return empty size, since request failed
	}
	// Read full content of response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ERR Couldn't read response body;", err)
	}
	resp.Body.Close()

	// body (multiple blocks, one for each cpu/core. Each block contains a line with the amount of cpu cores):
	// ...
	// cpu cores	: 1
	// ...
	// cpu cores	: 1
	// ...
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "") // spaces
		line = strings.ReplaceAll(line, "	", "") // tabs
		if strings.HasPrefix(line, "cpucores:") {
			cores, err = strconv.Atoi(strings.Split(line, ":")[1])
			if err != nil {
				log.Fatalln("ERR Couldn't parse cpu cores;", err)
			}

			break
		}
	}

	if debug {
		log.Println("INF Found core amount for ip", ipAddress)
	}

	return cores
}
