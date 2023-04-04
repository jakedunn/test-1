# Will your computer quarantine AsyncRAT malware?

AsyncRAT is one of the most popular open-source Remote Access Tools (RATs). Some capabilities of AsyncRAT include command and control (C2), keylogging, browser credential dumping, remote file transfer, and remote desktop control. In addition, AsyncRAT is delivered to victims in multiple ways, including spear-phishing attachments, .bat files, and obfuscated PowerShell scripts. This test will monitor if any endpoint defense quarantines the malware.

Example Output:

```
[+] Extracting file for quarantine test
[+] Pausing for 1 second to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the AsyncRAT file has been defanged, so even if run, it won't execute.

A PowerShell file of the AsyncRAT malware is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
