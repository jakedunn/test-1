# Mimikatz

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:AA18-284a](https://www.cisa.gov/news-events/cybersecurity-advisories/aa18-284a)</kbd>

This test drops a defanged Mimikatz executable. Mimikatz is a post-exploitation tool commonly used by attackers to steal sensitive authentication credentials from a compromised system. With these credentials, attackers can gain access to other systems and networks within the victim's organization. Mimikatz exploits weaknesses in the Windows security architecture to extract the credentials, making it a high-risk to any organization. This test will monitor if any endpoint defense quarantines the malware.

## How

> Safety: the malware used has been defanged, so even if run, it will immediately exit.

The malware file is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.

Example Output:
```bash
[+] Extracting file for quarantine test
[+] Pausing for 3 seconds to gauge defensive reaction
[-] Malicious file was not caught
```

## Resolution

If this test fails:

* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
