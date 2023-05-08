# AsyncRAT Rootkit

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

AsyncRAT is one of the most popular open-source Remote Access Tools (RATs). Some capabilities of AsyncRAT include command and control (C2), keylogging, browser credential dumping, remote file transfer, and remote desktop control. In addition, AsyncRAT is delivered to victims in multiple ways, including spear-phishing attachments, .bat files, and obfuscated PowerShell scripts. First discovered in 2018 by PaloAlto's Unit 42, AsyncRAT has been used in multiple campaigns by serveral APT groups and has been most recently observed targetting the Travel industry. This test will monitor if any endpoint defense quarantines the malware.

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
