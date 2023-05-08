/*
ID: b696a087-0605-4d24-8f2e-eb838cc31498
NAME: AsyncRAT Rootkit
CREATED: 2023-04-03
*/
package main

import (
	_ "embed"

	Endpoint "github.com/preludeorg/test/endpoint"
)

//go:embed asyncrat.ps1
var malicious []byte

func test() {
	println("[+] Extracting file for quarantine test")
	println("[+] Pausing for 3 seconds to gauge defensive reaction")
	if Endpoint.Quarantined("asyncrat.ps1", malicious) {
		println("[+] Malicious file was caught!")
		Endpoint.Stop(105)
	}
	println("[-] Malicious file was not caught")
	Endpoint.Stop(101)
}

func main() {
	Endpoint.Start(test)
}
