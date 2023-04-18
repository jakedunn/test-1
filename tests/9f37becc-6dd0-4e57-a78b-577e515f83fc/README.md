# Will your computer quarantine a malicious MSI file?

This test drops a defanged MSI file which contains a defanged IcedID dropper used for command and control and other post-exploitation activities. MSI files are a type of installer file used to distribute software applications. Malicious .msi files can be used to install malware, such as trojans or ransomware, on the victim's computer. Once installed, the malware may allow an attacker to steal sensitive information, monitor user activity, or take control of the compromised system. It is important to note that legitimate MSI files can also be used for software distribution, so it is crucial to verify the authenticity and safety of any downloaded MSI files. This test will monitor if any endpoint defense quarantines the malware.

Example Output:
```
[+] Extracting file for quarantine test
[+] Pausing for 1 second to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the MSI malware has been defanged, so even if run, it will immediately exit.

An MSI file is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
