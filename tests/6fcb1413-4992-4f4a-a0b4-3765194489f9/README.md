# Cobalt Strike

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:AR21-148a](https://www.cisa.gov/news-events/analysis-reports/ar21-148a)</kbd>

This test drops a defanged Cobalt Strike DLL. Cobalt Strike is a commercial penetration testing and threat emulation software. It provides a powerful platform for security professionals to conduct simulated cyber-attacks and evaluate the effectiveness of an organization's defenses. It offers a wide range of features, such as reconnaissance, social engineering, network exploitation, and post-exploitation tools. Cobalt Strike has also been abused by cybercriminals and advanced persistent threat (APT) groups for malicious purposes. In the hands of attackers, it can be used to gain unauthorized access to systems, exfiltrate sensitive data, and create backdoors for ongoing access. Its versatile and sophisticated capabilities have made it a popular tool among both legitimate security testers and malicious threat actors. This test will monitor if any endpoint defense quarantines the malware.

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
