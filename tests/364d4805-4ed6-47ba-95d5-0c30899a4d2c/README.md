# Will your computer quarantine Yanluowang ransomware?

This test drops a defanged Yanluowang executable, which threat actors use to encrypt and exfiltrate data off of compromised systems. Yanluowang ransomware was first used against a high-profile organization in a targeted attack, and is a underdeveloped ransomware family. Before deploying the ransomware, the threat actors would utilize AdFind, a legitimate command-line Active Directory query tool, on the victim organization's network. This allows the threat actor to do reconnaissance tasks, including gaining access to information needed for lateral movement through their victim's network. Once deployed, the ransomware would drop a ransom note which warns the victim not to reach out for help or risk DDoS and repeated ransomware attacks. This test will monitor if any endpoint defense quarantines the malware.

Example Output:

```
[+] Extracting file for quarantine test
[+] Pausing for 3 seconds to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the exe file has been defanged, so even if run, it won't execute.

A malicious exe file is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
