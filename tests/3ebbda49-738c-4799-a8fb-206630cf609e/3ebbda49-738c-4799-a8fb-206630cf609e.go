/*
ID: 3ebbda49-738c-4799-a8fb-206630cf609e
NAME: Safety Check
CREATED: 2023-03-13
*/
package main

import (
	"runtime"
	"time"

	Endpoint "github.com/preludeorg/test/endpoint"
)

var supported = map[string][]string{
	"windows": {"powershell.exe", "-c", "sleep 15"},
	"darwin":  {"bash", "-c", "sleep 15"},
	"linux":   {"bash", "-c", "sleep 15"},
}

func test() {
	command := supported[runtime.GOOS]
	println("[+] Sleeping for 15 seconds")
	Endpoint.Shell(command)
	println("[-] VST did not stop properly, returned exit code 101")
	Endpoint.Stop(101)
}

func main() {
	println("[+] Starting test")
	// Adding raw Endpoint.Start code so that we can safely update the Endpoint.Stop to exit 100 instead of 102.
	go func() {
		test()
	}()

	select {
	case <-time.After(10 * time.Second):
		println("[+] Timeout reached as intended.")
		Endpoint.Stop(100)
	}
}
