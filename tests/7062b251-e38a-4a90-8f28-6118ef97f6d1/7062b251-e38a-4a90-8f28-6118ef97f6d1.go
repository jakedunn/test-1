/*
ID: 7062b251-e38a-4a90-8f28-6118ef97f6d1
NAME: Is an endpoint control running?
CREATED: 2023-04-11
*/
package main

import (
	"fmt"
	"runtime"
	"strings"

	Endpoint "github.com/preludeorg/test/endpoint"
)

var supported = map[string]string{
	"windows": "tasklist /FI 'IMAGENAME eq %s'",
	"darwin":  "ps aux | grep %s | grep -v grep",
	"linux":   "ps aux | grep %s | grep -v grep",
}

var exec = map[string][]string{
	"windows": {"powershell.exe"},
	"darwin":  {"bash", "-c"},
	"linux":   {"bash", "-c"},
}

func test() {
	edrList := []string{"com.crowdstrike.falcon.Agent", "CSFalconService.exe", "falcon-sensor", "MsSense.exe", "mdatp"}

	for _, process := range edrList {
		command := fmt.Sprintf(supported[runtime.GOOS], process)
		finalCommand := append(exec[runtime.GOOS], command)
		result, err := Endpoint.Shell(finalCommand)
		if err != nil {
			continue
		}
		if strings.Contains(result, process) {
			println("[+]", process, "is running")
			Endpoint.Stop(100)
		}
	}
	println("[-] No endpoint control identified")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
