# Yanluowang Ransomware

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test drops a defanged Yanluowang executable, which threat actors use to encrypt and exfiltrate data off of compromised systems. The ransomware group was first discovered by Symantec back in October 2021. Yanluowang has been used in highly targeted attacks against organizations such as Cisco, Walmart and others. Before deploying the ransomware, the threat actors would utilize AdFind, a legitimate command-line Active Directory query tool, on the victim organization's network. This allows the threat actor to do reconnaissance tasks, including gaining access to information needed for lateral movement through their victim's network. Once deployed, the ransomware would drop a ransom note which warns the victim not to reach out for help or risk DDoS and repeated ransomware attacks. This test will monitor if any endpoint defense quarantines the malware.

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
