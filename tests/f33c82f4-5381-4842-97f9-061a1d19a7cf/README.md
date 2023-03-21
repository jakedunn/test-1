# Will your computer quarantine BugSplat malware?

This test drops a defanged NOBELIUM iso file containing two files, `BugSplatRc64.dll` aimed at collecting and exfiltrating information about the target system to Notion.com and `Instruction.lnk`, which is used to execute the BugSplat .dll file. The BugSplat malware was first observed by BlackBerry researchers in March 2023, being used by NOBELIUM, aka APT29, a sophisticated Russian state-sponsored threat actor in recent campaigns against the European Union. This test will monitor if any endpoint defense quarantines the malware.

Example Output:
```
[+] Extracting file for quarantine test
[+] Pausing for 1 second to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the Bugsplat malware and .lnk file has been defanged, so even if run, it will immediately exit.

An iso file containing the Blugsplat malware is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
