/*
ID: f33c82f4-5381-4842-97f9-061a1d19a7cf
NAME: BugSplat Trojan
CREATED: 2023-03-20
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed Instructions.iso
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 3 seconds to gauge defensive reaction")
	if Endpoint.Quarantined("Instructions.iso", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
