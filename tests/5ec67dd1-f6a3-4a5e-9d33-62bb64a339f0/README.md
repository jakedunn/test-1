# LockBit

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:AA23-075a](https://www.cisa.gov/news-events/cybersecurity-advisories/aa23-075a)</kbd>

This test drops a defanged LockBit executable, which threat actors use to encrypt and exfiltrate data off of compromised systems. LockBit has been known to specifically target non-Russian systems and performs a series of post-exploitation tasks, including gathering system information, stopping a process, creating and deleting a file, and scanning the network for other online machines. Since January 2020, LockBit has functioned as an affiliate-based ransomware variant; affiliates deploying the LockBit RaaS use many varying TTPs and attack a wide range of businesses and critical infrastructure organizations making prevention challenging. This test will monitor if any endpoint defense quarantines the malware.

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
