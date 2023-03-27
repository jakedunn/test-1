# Will your computer quarantine QuakBot malicious OneNote file?

This test drops a defanged QuakBot OneNote file. OneNote allows attackers to embed malicious files such as .hta and executables. Malicious OneNote files require users to click on the embedded file for execution. While there is a pop-up warning message when clicking on the object, most users will click through and allow the object to be executed, in turn infecting their system. QuakBot is a banking trojan primarily used to steal financial data, including browser information, keystrokes, and credentials. Once QakBot has successfully infected an environment, the malware installs a backdoor allowing the threat actor to drop additional malware and has been commonly used by larger ransomware groups such as BlackBasta. This test will monitor if any endpoint defense quarantines the malware.

Example Output:

```
[+] Extracting file for quarantine test
[+] Pausing for 1 second to gauge defensive reaction
[-] Malicious file was not caught
```

Defenses should quarantine files with known signatures immediately.

## How

> Safety: the QuakBot OneNote file has been defanged, so even if run, it won't execute.

A OneNote file of the QuakBot malware is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.
