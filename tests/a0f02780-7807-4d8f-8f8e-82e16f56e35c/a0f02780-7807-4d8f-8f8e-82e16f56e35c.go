/*
ID: a0f02780-7807-4d8f-8f8e-82e16f56e35c
NAME: Will your computer quarantine Mimikatz?
CREATED: 2023-04-17
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed mimikatz.exe
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 1 second to gauge defensive reaction")
	if Endpoint.Quarantined("mimikatz.exe", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
