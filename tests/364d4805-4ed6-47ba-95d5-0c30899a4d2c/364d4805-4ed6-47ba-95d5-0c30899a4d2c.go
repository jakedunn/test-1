/*
ID: 364d4805-4ed6-47ba-95d5-0c30899a4d2c
NAME: Will your computer quarantine Yanluowang ransomware?
CREATED: 2023-04-25
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed yanluowang.exe
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 3 seconds to gauge defensive reaction")
	if Endpoint.Quarantined("yanluowang.exe", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
