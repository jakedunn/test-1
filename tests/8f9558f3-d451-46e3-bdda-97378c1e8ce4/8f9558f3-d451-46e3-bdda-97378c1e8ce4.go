/*
ID: 8f9558f3-d451-46e3-bdda-97378c1e8ce4
NAME: Will your computer quarantine a malicious HTA file?
CREATED: 2023-04-11
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed WORDOFJAH.hta
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 3 seconds to gauge defensive reaction")
	if Endpoint.Quarantined("WORDOFJAH.hta", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
