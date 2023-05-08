/*
ID: 994c0cb2-a84c-41b2-ab46-7d4d104c09d8
NAME: QuakBot Trojan
CREATED: 2023-03-27
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed quakbot.one
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 3 seconds to gauge defensive reaction")
	if Endpoint.Quarantined("quakbot.one", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
