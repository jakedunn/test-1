# Will your computer quarantine Mimikatz?

This test drops a defanged Mimikatz executable. Mimikatz is a post-exploitation tool commonly used by attackers to steal sensitive authentication credentials from a compromised system. With these credentials, attackers can gain access to other systems and networks within the victim's organization. Mimikatz exploits weaknesses in the Windows security architecture to extract the credentials, making it a high-risk to any organization. This test will monitor if any endpoint defense quarantines the malware.

Example Output:
```
[+] Extracting file for quarantine test
[+] Pausing for 1 second to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the Mimikatz binary has been defanged, so even if run, it will immediately exit.

An executable file of Mimikatz is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
